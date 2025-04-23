// Use `go run foo.go` to run your program

package main

// 3 DONE 4 NOT DONE

import (
    . "fmt"
    "runtime"
    // "time"
)

var i = 0

func incrementing(incre_done chan bool) {
    //TODO: increment i 1000000 times
    for j := 0 ; j < 1000000; j++ {
        i++
    }
    incre_done <- true
}

func decrementing(decre_done chan bool) {
    //TODO: decrement i 1000000 times
    for j := 0 ; j < 1000001; j++ {
        i--
    }
    decre_done <- true
}

func main() {
    // What does GOMAXPROCS do? What happens if you set it to 1?
    runtime.GOMAXPROCS(2) 
    // *GOMAXPROCS runs as many threads as indicated in its 
    // parameter. Now it has 2, meaning two threads are running.
    // Since the concurrancy isn't handled correctly, we run 
    // into race cond. giving us an unpredictable result, i.
    // If we had a 1 instead of a 2, we would do one thread at
    // a time, giving us a correct value of i, 0.*
	
    incre_done := make(chan bool)
    decre_done := make(chan bool)
	
    // TODO: Spawn both functions as goroutines
    go incrementing(incre_done)
    go decrementing(decre_done)

    <- incre_done
    <- decre_done
    
    // We have no direct way to wait for the completion of a goroutine
    // (without additional synchronization of some sort)
    //  We will do it properly with channels soon. For now: Sleep.
    // time.Sleep(500*time.Millisecond)
    Println("The magic number is:", i)
}
