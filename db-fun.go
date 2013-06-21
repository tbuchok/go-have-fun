package main

import (
	"fmt"
	"github.com/fzzy/radix/redis"
	"os"
	"strconv"
	"sync"
	"time"
)

func handleErr(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func dbFun(wg *sync.WaitGroup, wait int) {
	client, err := redis.Dial("tcp", "127.0.0.1:6379")
	handleErr(err)
	defer client.Close()
	defer wg.Done()

	response := client.Cmd("select", 5)
	handleErr(response.Err)

	key := strconv.Itoa(wait)
	value := "bar"
	response = client.Cmd("set", key, value)
	handleErr(response.Err)
	fmt.Printf("Waiting %v seconds ... \n", wait)
	time.Sleep(time.Duration(wait) * time.Second)
	result, err := client.Cmd("get", key).Str()
	handleErr(err)

	_ = client.Cmd("del", key)

	fmt.Printf("%s:%s\n", key, result)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go dbFun(&wg, 5)
	wg.Add(1)
	go dbFun(&wg, 2)
	wg.Wait()
}
