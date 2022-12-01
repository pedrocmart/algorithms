package main

import (
	"reflect"
)

func solution(shoes [][]int) bool {

	return pairOfShoes(shoes)
}

func pairOfShoes(shoes [][]int) bool {
	if reflect.DeepEqual(shoes, []interface{}{}) || len(shoes)%2 != 0 {
		return false
	}
	for !reflect.DeepEqual(shoes, []interface{}{}) {
		if len(shoes) == 0 {
			return true
		}
		shoe := shoes[0]
		if func() int {
			for i, v := range shoes {
				if reflect.DeepEqual(v, []int{1 - shoe[0], shoe[1]}) {
					return i
				}
			}
			return -1
		}() == -1 {
			return false
		} else {
			shoes = shoes[1:]
			i := 0
			for !reflect.DeepEqual(shoes[i], []int{1 - shoe[0], shoe[1]}) {
				i += 1
			}
			sh1 := shoes[0:i]
			sh2 := shoes[i+1:]

			shoes = nil
			shoes = append(shoes, sh1...)
			shoes = append(shoes, sh2...)
		}
	}
	return true
}
