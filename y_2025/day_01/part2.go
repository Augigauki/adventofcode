package main

func countZeroes(direction string, distance int, current int) (int, int) {
	zeroClicks := 0
	startingPoint := current
	if direction == "L" {
		for distance > 99 {
			distance = distance - 100
			zeroClicks++
		}
		current -= distance
		if current < 0 {
			current = 100 + current
			if startingPoint != 0 {
				zeroClicks++
			}
		}
		if current == 0 {
			zeroClicks++
		}
	}

	if direction == "R" {
		for distance > 99 {
			distance = distance - 100
			zeroClicks++
		}
		current += distance
		if current > 99 {
			current = current - 100
			zeroClicks++
		}
	}
	/* if current == 0 {
		zeroClicks++
	} */
	return current, zeroClicks
}
