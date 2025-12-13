package main

import (
	"aoc/utils-go"
	"fmt"
)

const SUM_TO_BE_EQ int = 2020

func first(input []int) {
	println("First ------")

	for i, vi := range input {
		for _, vj := range input[i+1:] {
			if vi+vj == SUM_TO_BE_EQ {
				fmt.Printf("%d * %d = %d\n", vi, vj, vi*vj)
				return
			}
		}
	}
}

func second(input []int) {
	println("Second ------")

	for i, vi := range input {
		for j, vj := range input[i+1:] {
			for _, vk := range input[j+1:] {
				if vi+vj+vk == SUM_TO_BE_EQ {
					fmt.Printf("%d * %d * %d = %d\n", vi, vj, vk, vi*vj*vk)
					return
				}
			}
		}
	}
}

func main() {
	inputList := utils.ReadLines("1.in")
	inputInts := utils.ConvertToInts(inputList)

	first(inputInts)

	second(inputInts)
}
