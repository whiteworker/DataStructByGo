package main

import (
	"fmt"
)


func reverseWords(s string) string {
	end := len(s)-1
	start := end
	res := ""
	for start >=0{
		for start>=0&&s[start]!=' ' {
			start--
		}
		res +=s[start+1:end+1]
		for(start >=0&&s[start]==' '){
			start--
		}
		end=start
		if(start>=0&&res!=""){
			res += " "
		}
	}
	return res
}
func main(){
	word := " i win  game "
	fmt.Printf(	reverseWords(word))
	fmt.Printf("|%v|",word[:len(word)-1])
}