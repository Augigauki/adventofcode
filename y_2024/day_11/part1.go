package main

import (
	"fmt"
	"math"
)

func CountIntStones(stones []int, blinkAmount int) int {
	fmt.Println("Initial arrangement:")
	fmt.Println(stones)
	cache := map[int]int{}

	for _, v := range stones {
		cache[v] = 1
	}

	for i := 0; i < blinkAmount; i++ {
		cache = blinkCache(cache)
	}

	sum := 0

	for _, v := range cache {
		sum += v
	}

	return sum
}

func blinkCache(stones map[int]int) map[int]int {
	newStones := map[int]int{}

	add := func(key, incr int) {
		newStones[key] += incr
	}

	for stone, count := range stones {
		if stone == 0 {
			add(1, count)
		} else if digits := numDigits(stone); digits%2 == 0 {
			filter := powTen(digits / 2)
			left, right := stone/filter, stone%filter
			add(left, count)
			add(right, count)
		} else {
			add(stone*2024, count)
		}
	}
	return newStones
}

func numDigits(i int) int {
	if i == 0 {
		return 1
	}
	return int(math.Log10(float64(i))) + 1
}

func powTen(pow int) int {
	return int(math.Pow10(pow))
}
