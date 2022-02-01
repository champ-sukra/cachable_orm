package main

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CachedOrm struct {
	db *gorm.DB
	cm *CachedMap
}

func NewCachedOrm(db *gorm.DB, cm *CachedMap) *CachedOrm {
	orm := &CachedOrm{
		db: db,
		cm: cm,
	}
	return orm
}

func (co *CachedOrm) Find(m interface{}, pk string, val string, d time.Duration) error {
	isFound := co.cm.Get(m, val)
	if isFound {
		log.Println("found record")
		co.cm.Set(val, m, d)
		return nil
	}
	log.Println("not found record")

	result := co.db.First(&m, fmt.Sprintf("%s = '%s'", pk, val))

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		log.Println("no record")
		return errors.New("no_record")
	}

	co.cm.Set(val, m, d)
	return nil
}

// Insert object, which contain non-autoincrement primary key
func (co *CachedOrm) Insert(m interface{}, key string, d time.Duration) error {
	isExist, err := co.IsExists(m)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New("invalid_statement")
	}

	co.cm.Set(key, m, d)
	go co.SaveToDb(m)
	return nil
}

func (co *CachedOrm) Update(m interface{}, key string, d time.Duration) error {
	if co.cm.IsExists(key) {
		log.Println("found record")
		co.cm.Set(key, m, d)
		go co.db.Updates(m)
		return nil
	}

	log.Println("not found record")

	isExist, err := co.IsExists(m)
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New("invalid_statement")
	}

	co.cm.Set(key, m, d)
	go co.db.Updates(m)
	return nil
}

func (co *CachedOrm) IsExists(m interface{}) (bool, error) {
	var exists int
	result := co.db.Model(m).
		Select("EXISTS(SELECT 1)").
		Limit(1).
		Scan(&exists)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 1 {
		return true, nil
	}
	return false, nil
}

func (co *CachedOrm) SaveToDb(m interface{}) {
	result := co.db.Save(m)
	if result.Error != nil {
		log.Println(result.Error)
	}
}

//func BulkFind(ms interface{}, items []string) {
//	db.Where("token IN ?", items).Find(&ms)
//}
//
//func SaveUsingAutoInc(m interface{}, key string, isCache bool) {
//	if myLastPri == 0 {
//		m_ := map[string]interface{}{}
//		_ = db.Model(&m).Last(&m_)
//		myLastPri = m_[key].(int64)
//	}
//
//	if isCache {
//		nextPri := myLastPri + 1
//		v := reflect.ValueOf(m).Elem()
//		if f := v.FieldByName(strings.Title(key)); f.IsValid() {
//			f.SetInt(nextPri)
//		}
//		//myMap[nextPri] = m
//		myLastPri = nextPri
//		go SaveToDb(m)
//	} else {
//		SaveToDb(m)
//	}
//}
//
