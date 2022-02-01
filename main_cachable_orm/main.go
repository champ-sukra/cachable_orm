package main

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var dsn string
	dbHost := "ichaithat.local"
	dbPort := "3306"
	dbUser := "root"
	dbPass := "password"
	dbName := "auth"
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}

var cm *CachedMap
var cOrm *CachedOrm

func main() {

	start := time.Now()
	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	cm = NewCachedMap()
	cOrm = NewCachedOrm(db, cm)

	//var ts []Token
	//BulkFind(&ts, []string{"YWRtaW46YWRtaW4xMjM=", "ZS1wZW5zaW9uOmFkbWluMTIz"})
	//log.Println(ts[1].ProfileId)

	n := Token{Token: "xxxxx", ProfileId: 5}
	cOrm.Insert(&n, "xxxxx", time.Second*5)

	//t := Token{Token: "YWRtaW46YWRtaW4xMjM="}
	//Find(&t, "YWRtaW46YWRtaW4xMjM=", time.Second)
	//log.Println(t.ProfileId)
	//
	for i := 0; i < 10; i++ {
		x := Token{Token: "xxxxx"}
		cOrm.Find(&x, "token", n.Token, time.Second*5)
		log.Println(n.ProfileId)
		time.Sleep(time.Second)
	}

	//n := Token{Token: "YWRtaW46YWRtaW4xMjM=", ProfileId: 6}
	//Update(&n, "YWRtaW46YWRtaW4xMjM=")

	//time.AfterFunc(time.Second*5, func() {
	//	Find(&t, "YWRtaW46YWRtaW4xMjM=")
	//})

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)

	http.ListenAndServe(":8080", nil)
}
