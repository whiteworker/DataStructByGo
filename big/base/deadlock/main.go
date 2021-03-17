package main

import (
    //"fmt"
    "sync"
    "time"

)

var (
    mu1 sync.Mutex
    mu2 sync.Mutex
    wg sync.WaitGroup
)

func deadlock1() {
    wg.Add(2)

    go func() {
        mu1.Lock()
        time.Sleep(1 * time.Second)
        mu2.Lock()
		wg.Done()
		//mu2.Unlock()
		//mu1.Unlock()
    }()

    go func() {		
        mu2.Lock()
		mu1.Lock()
        wg.Done()
		//mu1.Unlock()
		//mu2.Unlock()		
    }()

   

    wg.Wait()
}

func deadlock2(){
    ch:=make(chan int)  //这就是在main程里面发生的死锁情况    
    ch<-6   //  这里会发生一直阻塞的情况，执行不到下面一句
    go func(){
        <-ch
    }()
}
func main(){
    deadlock2()
}