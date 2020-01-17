package main

import (
    "fmt"
    "time"
)

func sum(a int, b int, c chan int) {
    c <- a + b // send sum to c
}

func sum1(a int, b int, c chan int) {
    time.Sleep(100 * time.Millisecond)
    c <- a + b // send sum to c
}

func main() {
    c := make(chan int)
    go sum1(10, 10, c)
    go sum(20, 22, c)
    x, y := <-c, <-c // receive from c

    fmt.Println(x, y)
}
