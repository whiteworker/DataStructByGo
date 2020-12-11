package main

import "fmt"

func f1() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func f2() {
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func f3() {
	for i := 0; i < 5; i++ {
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}
}

func f4() int {
	t := 5
	defer func() {
		t++
	}()
	return t
}

func f5() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f6() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f7() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
func main() {
	// f1()
	// f2()
	// f3()
	// f4()
	// f5()
	// f6()
	// f7()
	fmt.Print(f5())
}
