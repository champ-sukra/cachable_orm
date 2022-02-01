package main

import (
	"fmt"
	cmap "github.com/streamrail/concurrent-map"
	"io/ioutil"
	"log"
	"math/big"
	"net/url"
	"sync"
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

func main() {

	var wg sync.WaitGroup

	start := time.Now()
	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	f, _ := ioutil.ReadFile("qs_to_index.json")
	//var confs []map[string]string
	//_ = json.Unmarshal(f, &confs)

	/*
		var myMap sync.Map
		//myMap := make(map[string][]byte, 0)
		wg.Add(1)
		go addToNativeMap(&myMap, "1", f, &wg)
		wg.Add(1)
		go addToNativeMap(&myMap, "2", f, &wg)
		/*/
	myMap := cmap.New()
	myMap.Count()
	wg.Add(1)
	go addToConcurrentMap(myMap, "1", f, &wg)
	wg.Add(1)
	go addToConcurrentMap(myMap, "2", f, &wg)
	//*/

	wg.Wait()

	//val, _ := myMap.Load("99992")
	log.Printf("count: %d", len(myMap))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

type x struct {
	key string
}

var mutex sync.Mutex

func addToNativeMap(myMap *sync.Map, t string, bs []byte, wg *sync.WaitGroup) {
	for i := 0; i < 10000; i++ {
		myMap.Store(fmt.Sprintf("%d%s", i, t), bs)
	}
	defer wg.Done()
}

func addToConcurrentMap(myMap cmap.ConcurrentMap, t string, bs []byte, wg *sync.WaitGroup) {
	for i := 0; i < 10000; i++ {
		myMap.Set(fmt.Sprintf("%d", i), string(bs))
	}
	defer wg.Done()
}
