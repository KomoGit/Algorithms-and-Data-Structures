package main

import "log"

var numbers = []int{5, 6, 2, 3, 1, 4, 7, 9, 8, 10}

func main() {
	//i := 0
	for j := 0; j < len(numbers); j++ {
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				swap(numbers, i, i+1)
			}
		}
	}

	log.Println(numbers)
}

func swap(arrNum []int, i, j int) {
	temp := numbers[i]
	numbers[i] = numbers[j]
	numbers[j] = temp

}
