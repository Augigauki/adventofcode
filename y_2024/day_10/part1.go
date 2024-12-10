package main

import "fmt"

type Trailhead struct {
	Pos       Pos
	Score     Score
	Goals     []Goal
	Traversed []Pos
	Value     int
}

type Score struct {
	UniqueScore int
	AllScores   int
}

type Goal struct {
	Pos Pos
}

type Pos struct {
	Line int
	Char int
}

func FindTrailHeadScores(topoMap [][]int) int {
	var trailHeadScore = 0
	FindTrailHeads(topoMap)

	return trailHeadScore
}

func FindTrailHeads(topoMap [][]int) {
	var trailHeads []Trailhead
	for i, line := range topoMap {
		for j, char := range line {
			if char == 0 {
				trailHeads = append(trailHeads, Trailhead{Pos{i, j}, Score{0, 0}, []Goal{}, []Pos{}, 0})
			}
		}
	}
	for _, trailHead := range trailHeads {
		HikeTrail(&trailHead, topoMap)
	}
	fmt.Println("Trail heads: ", trailHeads)
}

func HikeTrail(trailHead *Trailhead, topoMap [][]int) {
	fmt.Println("Hiking trail...")
	startPos := trailHead.Pos
	if trailHead.Value == 9 {
		fmt.Println("Goal reached!")
		trailHead.Score.AllScores++
		if len(trailHead.Goals) == 0 {
			trailHead.Score.UniqueScore++
		} else {
			for _, goal := range trailHead.Goals {
				var alreadyCounted = false
				if goal.Pos.Line == trailHead.Pos.Line && goal.Pos.Char == trailHead.Pos.Char {
					alreadyCounted = true
				}
				if !alreadyCounted {
					trailHead.Score.UniqueScore++
				}
			}
		}
		trailHead.Goals = append(trailHead.Goals, Goal{Pos{trailHead.Pos.Line, trailHead.Pos.Char}})
		fmt.Println("Trailhead: ", trailHead.Score)
		ResetTrailHead(trailHead, startPos)
		return
	}
	if trailHead.Pos.Line > 0 {
		if topoMap[trailHead.Pos.Line-1][trailHead.Pos.Char] == trailHead.Value+1 {
			fmt.Println("Moving north")
			trailHead.Value++
			trailHead.Pos.Line--
			HikeTrail(trailHead, topoMap)
		}
	}
	if trailHead.Pos.Char < len(topoMap[trailHead.Pos.Line])-1 {
		if topoMap[trailHead.Pos.Line][trailHead.Pos.Char+1] == trailHead.Value+1 {
			fmt.Println("Moving east")
			trailHead.Value++
			trailHead.Pos.Char++
			HikeTrail(trailHead, topoMap)
		}
	}
	if trailHead.Pos.Line < len(topoMap)-1 {
		if topoMap[trailHead.Pos.Line+1][trailHead.Pos.Char] == trailHead.Value+1 {
			fmt.Println("Moving south")
			trailHead.Value++
			trailHead.Pos.Line++
			HikeTrail(trailHead, topoMap)
		}
	}
	if trailHead.Pos.Char > 0 {
		if topoMap[trailHead.Pos.Line][trailHead.Pos.Char-1] == trailHead.Value+1 {
			fmt.Println("Moving west")
			trailHead.Value++
			trailHead.Pos.Char--
			HikeTrail(trailHead, topoMap)
		}

	}
}

func ResetTrailHead(trailHead *Trailhead, startPos Pos) {
	trailHead.Value = 0
	trailHead.Pos = startPos
}
