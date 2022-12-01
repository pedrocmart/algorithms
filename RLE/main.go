package main

import "strconv"

func solution(inputString string) string {
	n := len(inputString)
	ret := ""
	for i := 0; i < n; i++ {
		count := 1
		for i < n-1 && string(inputString[i]) == string(inputString[i+1]) {
			count++
			i++
		}
		// Print character and its count
		ret += strconv.Itoa(count)
		ret += string(inputString[i])
	}
	return ret
}
