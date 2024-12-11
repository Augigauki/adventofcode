package main

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"strings"
	"sync"
)

func CountStones(stones string, blinkAmount int) int {
	fmt.Println("Initial arrangement:")
	fmt.Println(stones + "\n")
	for i := 0; i < blinkAmount; i++ {
		fmt.Printf("Blink %v\n", i+1)
		stones = blinkConcurrent(stones)
		fmt.Println(stones + "\n")
	}
	return len(strings.Split(stones, " "))
}

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

func splitStones(stones string) []string {
	return strings.Split(stones, " ")
}

func blinkConcurrent(input string) string {
	stones := splitStones(input)
	numGoroutines := 4
	var wg sync.WaitGroup
	resultChans := make([]chan string, numGoroutines)

	chunkSize := (len(stones) + numGoroutines - 1) / numGoroutines
	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if start >= len(stones) {
			break
		}
		if end > len(stones) {
			end = len(stones)
		}
		resultChans[i] = make(chan string, chunkSize*2) // Buffer size to avoid blocking
		wg.Add(1)
		go func(stonesChunk []string, resultChan chan string) {
			defer wg.Done()
			for _, stone := range stonesChunk {
				if stone == "0" {
					resultChan <- "1"
				} else {
					if len(stone)%2 == 0 {
						lefStone, rightStone := splitStoneInTwoOptimized(stone)
						resultChan <- lefStone
						resultChan <- rightStone
					} else {
						num := convStringToBigInt(stone)
						if num == nil {
							log.Fatal("Error converting stone to int")
						}
						num.Mul(num, big.NewInt(2024))
						resultChan <- num.String()
					}
				}
			}
			close(resultChan)
		}(stones[start:end], resultChans[i])
	}

	wg.Wait()
	var result []string
	for _, resultChan := range resultChans {
		for res := range resultChan {
			result = append(result, res)
		}
	}

	return strings.Join(result, " ")
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

func convStringToBigInt(stone string) *big.Int {
	bigNum := new(big.Int)
	num, ok := bigNum.SetString(stone, 10)
	if !ok {
		fmt.Printf("Error converting %v to big.Int", stone)
		return nil
	}
	return num
}

func splitStoneInTwoOptimized(stone string) (string, string) {
	half := len(stone) / 2
	firstHalf := stone[:half]
	secondHalf := strings.TrimLeft(stone[half:], "0")
	return firstHalf, secondHalf
}

/* func blink2(input string) string {
	stones := splitStones(input)
	result := make([]string, 0, len(stones)*2) //Preallocate with an estimated size

	for _, stone := range stones {
		if stone == "0" {
			result = append(result, "1")
		} else {
			if len(stone)%2 == 0 {
				lefStone, rightStone := splitStoneInTwoOptimized(stone)
				result = append(result, lefStone, rightStone)
			} else {
				num := convStringToBigInt(stone)
				if num == nil {
					log.Fatal("Error converting stone to int")
				}
				num.Mul(num, big.NewInt(2024))
				result = append(result, num.String())
			}
		}
	}
	return strings.Join(result, " ")
} */

/* func blink(input string) string {
	//newStones := ""
	stones := splitStones(input)

	var result strings.Builder

	for _, stone := range stones {
		if stone == "0" {
			result.WriteString("1 ")
		} else {
			if len(stone)%2 == 0 {
				lefStone, rightStone := splitStoneInTwo(stone)
				result.WriteString(lefStone + " " + rightStone + " ")
			} else {
				num := convStringToBigInt(stone)
				if num == nil {
					log.Fatal("Error converting stone to int")
				}
				num.Mul(num, big.NewInt(2024))
				result.WriteString(num.String() + " ")
			}
		}
	}
	return strings.TrimSpace(result.String())
} */

/* func splitStoneInTwo(stone string) (string, string) {
	var firstHalf, secondHalf string
	firstHalf = stone[:len(stone)/2]
	secondHalf = stone[len(stone)/2:]
	for len(secondHalf) > 1 && secondHalf[0] == '0' {
		//fmt.Println("Second half before trimming unused zeroes: ", secondHalf)
		secondHalf = secondHalf[1:]
		//fmt.Println("Second half after trimming unused zeroes: ", secondHalf)
	}
	return firstHalf, secondHalf
} */

/* func convStringToInt(stone string) int {
	num, err := strconv.Atoi(stone)
	if err != nil {
		fmt.Printf("Error converting %v to int", stone)
		return -1
	}
	return num
} */
