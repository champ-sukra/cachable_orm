package main

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"
)

func convertToFile() {
	temp := qss
	qss = nil

	f, _ := ioutil.ReadFile("qs_to_index.json")
	var confs []map[string]string
	_ = json.Unmarshal(f, &confs)

	var b strings.Builder
	for _, s := range temp {
		qps := strings.Split(s, "&")
		size := len(confs)
		for i := 0; i < size; i++ {
			conf := confs[i]
			val := ""
			for _, qpKv := range qps {
				qp := strings.Split(qpKv, "=")
				k := qp[0]
				v := qp[1]
				if conf["key"] == k {
					val = v
					break
				}
			}
			if val != "" {
				b.WriteString(val)
			}
			if i != size-1 {
				b.WriteString("|")
			}
		}
		b.WriteString("\n")
	}

	osF, _ := os.OpenFile("buffer_to_string", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	_, err := osF.WriteString(b.String())
	if err != nil {
		osF.Close()
		fmt.Println(err)
		return
	}
	osF.Close()
}

func setupCron() {
	c := cron.New()
	c.AddFunc("@every 1s", func() {
		log.Println("cron...")
		convertToFile()
	})
	c.Start()
}

var wg = &sync.WaitGroup{}

func routine() {
	defer wg.Done()
	fmt.Println("routine finished")
}

var qss []string

func main() {

	//setupCron()
	qss = make([]string, 0)

	start := time.Now()
	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	for i := 0; i < 10000; i++ {
		//time.Sleep(time.Millisecond)

		qs := "a=5555&b=chaithat&e=test"
		qss = append(qss, qs)
	}
	convertToFile()

	//for i := 0; i < 100; i++ {
	//	//time.Sleep(time.Nanosecond)
	//
	//	qs := "a=5555&b=chaithat"
	//	qss = append(qss, qs)
	//}
	//convertToFile()

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)

	//content, err := ioutil.ReadFile("buffer_to_string")
	//if err != nil {
	//	fmt.Println("Err")
	//}
	//log.Printf("len: %d\n", len(strings.Split(string(content), ":")))
}
