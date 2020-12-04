package main

import (
	"fmt"
	"strconv"
	"strings"
	"io/ioutil"
)

func readFile(fname string) (nums []int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func main() {
	fmt.Println("Advent of Code 2020 - Puzzle Day 1")

	nums, err := readFile("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(nums)
	fmt.Println(len(nums))
}
