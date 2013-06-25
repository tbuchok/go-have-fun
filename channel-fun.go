package main

import (
  "fmt"
  "time"
)

func a(c chan string) {
  for {
    c <- "channel!"
    time.Sleep(time.Second * 1) // wait 1 second and keep going!
  }
}

func main() {
  var c chan string = make(chan string)
  
  go a(c)

  for {
    select {
      default:
        fmt.Println(<-c)
    }
  }

}