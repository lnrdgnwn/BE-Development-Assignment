package main

import (
	"fmt"
	"strings"
)

func main (){
	var t int
	fmt.Scanln(&t)

	results := make([]string, t)

	for i := 0; i<t; i++ {
		var s string
		fmt.Scanln(&s)
		s=strings.ToUpper(s)

		countG := 0
		countC := 0
		for _, word := range s {
			if (word == 'G'){
				countG++
			} else if word == 'C'{
				countC++
			}
		}

		if (countG == countC && !strings.Contains(s, "DGD")){
			results[i] = "VALID"
			
		} else {
			results[i] = "TIDAK VALID"
		}

	}

	for i := 0; i<t; i++ {
		fmt.Println(results[i])
	}
}

