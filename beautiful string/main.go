package main

func solution(inputString string) bool {

	var alphabet [26]int
	for i := 0; i < len(inputString); i++ {
		alphabet[inputString[i]-97]++
	}
	for j := 1; j < len(alphabet); j++ {
		if alphabet[j] > alphabet[j-1] {
			return false
		}
	}
	return true
}
