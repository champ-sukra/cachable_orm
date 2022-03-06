package main

import (
	"fmt"
	"github.com/chaithat-sukra/cachable_orm/test_orm_handler/orm_handler"
	"log"
	"math/big"
	"time"
)

func main() {

	start := time.Now()
	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	db, g := orm_handler.SetupDatabase(orm_handler.DBConfig{
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		Password: "password",
		DBName:   "customer",
	}, []string{"customers"})

	for i := 0; i < 1000; i++ {
		now := time.Now()
		id := g.GetUniqueIdForKey("customers")
		cust := Customer{
			Id:          id,
			FirstName:   "",
			LastName:    "",
			Email:       fmt.Sprintf("%d@xxx.com", id),
			CreatedDate: now,
			UpdatedDate: now,
		}

		db.Create(&cust)
	}

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
	fmt.Println("main finished")
}

//func main() {
//
//	start := time.Now()
//	r := new(big.Int)
//	fmt.Println(r.Binomial(1000, 10))
//
//	db, _ := orm_handler.SetupDatabase(orm_handler.DBConfig{
//		Host:     "localhost",
//		Port:     "3306",
//		User:     "root",
//		Password: "password",
//		DBName:   "customer",
//	}, []string{"customers"})
//
//	for i := 0; i < 1000; i++ {
//		now := time.Now()
//		//id := g.GetUniqueIdForKey("customers")
//		cust := Customer{
//			//Id:          id,
//			FirstName:   "",
//			LastName:    "",
//			Email:       fmt.Sprintf("AB%d@xxx.com", i),
//			CreatedDate: now,
//			UpdatedDate: now,
//		}
//
//		db.Create(&cust)
//	}
//
//	elapsed := time.Since(start)
//	log.Printf("Binomial took %s", elapsed)
//
//	//wg.Wait()
//	//fmt.Println("main finished")
//}
