package main

import (
    "fmt"
    "time"
)    

func read(ch1 chan int,ch2 chan bool){
    for{       
        // v,ok := <-ch1
        // if ok{
        //     fmt.Printf("read a int is %d\n",v)
        // } else{
        //     ch2 <-true
        // }
        select{
            case v,ok :=<-ch1:
                if ok{
                    fmt.Printf("read a int is %d\n",v)
                } else{
                    ch2 <- true
                }
            default:
                ch2<-false
        }
     }
    
}

func write(ch chan int){
    for i:=0;i<10;i++{
         ch <- i
    }
    close(ch)
}

func main() {
    ch1 := make(chan int)
    ch2 := make(chan bool)

    go write(ch1)
    go read(ch1,ch2)

    //<-ch2
    time.Sleep(time.Second*2)
}