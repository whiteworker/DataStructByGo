package main
     
import (
    "fmt"
)
 
func main() {
 
    s :=[]int{1,2,3,4}
    m :=make(map[int]*int)
 
    for k,v:=range s{
        n := v
        m[k]=&n
    }
    for key, value := range m {
        fmt.Printf("map[%v]=%v\n", key, *value)
    }
 
    fmt.Println(m)
}
