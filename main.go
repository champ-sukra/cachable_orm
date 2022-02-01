package main

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron"
	"io/ioutil"
	"log"
	"math/big"
	"sync"
	"time"
)

func setupCron() {
	c := cron.New()
	c.AddFunc("@every 5s", func() {
		//for f := range fs {
		//	delete(fs, f)
		//}

	})
	c.Start()
}

var wg = &sync.WaitGroup{}

func routine() {
	defer wg.Done()
	fmt.Println("routine finished")
}

func main() {

	//wg.Add(100)
	//go routine()

	//setupCron()

	start := time.Now()
	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	items := make([]string, 0)
	for i := 0; i < 10000; i++ {

		qs := "a=5555&b=chaithat"
		//*
		//b.WriteString(qs)
		items = append(items, qs)
		//_, err := fmt.Fprintln(osF, qs)
		/*/
		size := len(log)
		_, err := fmt.Fprint(osF, strconv.Itoa(size)+":"+log)
		//*/
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}
	}

	//var osF *os.File
	//if osF == nil {
	//	fmt.Println("not found..")
	//	osF, _ = os.OpenFile("buffer_to_string", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//}
	//osF.Close()

	conf, _ := ioutil.ReadFile("qs_to_index.json")
	var m map[string]int
	_ = json.Unmarshal(conf, &m)
	//
	//bytes, err := ioutil.ReadFile("buffer_to_string")
	//if err != nil {
	//	fmt.Println("Err")
	//}
	//
	//for _, s := range strings.Split(string(bytes), "\n") {
	//	log.Println(s)
	//}

	//log.Printf("len: %d\n", len(strings.Split(string(bytes), "\n")))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)

	//wg.Wait()
	//fmt.Println("main finished")
}
