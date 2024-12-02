package main

import (
	"fmt"
)

func findSafe(levels [][]int){
	var safeLevels [][]int
	for _, level := range levels {
		var safe bool
		var increasing bool = level[0] < level[1]
		for _, num := range level {
			if(increasing){
				if(level[num] < level[num+1] && level[num+1] - level[num] <= 3){
					safe = true
				} else {
					safe = false
					break
				}
			}
			if(!increasing){
				if(level[num] > level[num+1] && level[num] - level[num+1] <= 3){
					safe = true
				} else {
					safe = false
					break
				}
			}
		}
		if(safe){
			safeLevels = append(safeLevels, level)
		}
	}
	fmt.Println(safeLevels)
}