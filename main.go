package main

import "fmt"

func main() {
	i := 2
	var (
		answer bool
		num    int
	)
	var _, err = fmt.Scan(&num)
	if err != nil {
		return
	}
	for ; i < num; i++ {
		if num%i == 0 {
			answer = false
			break
		} else {
			answer = true
		}
	}
	if answer == false {
		fmt.Println("否")
	} else {
		fmt.Println("是")
	}

}
