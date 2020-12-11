package concurrent

import (
	"fmt"
	"sync"
)

func PrintLetterAndNumber() {
	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}
	wait.Add(1)
	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
				break
			default:
				break
			}
		}
	}()
	go func(wait *sync.WaitGroup) {
		i := 0
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		for {
			select {
			case <-letter:
				if i >= len(str)-1 {
					//i = 0
					wait.Done()
					break
				}
				fmt.Print(str[i : i+1])
				i++
				if i >= len(str)-1 {
					//i = 0
					number <- true
					break
				}
				fmt.Print(str[i : i+1])
				i++
				if i >= len(str)-1 {
					//i = 0
					number <- true
					break
				}

				number <- true
			default:
				break
			}
		}
	}(&wait)
	number <- true
	wait.Wait()
}
