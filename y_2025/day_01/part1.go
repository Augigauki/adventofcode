package main

func turnDial(direction string, distance int, current int) int {
	if direction == "L" {
		for distance > 99 {
			distance = distance - 100
		}
		current -= distance
		if current < 0 {
			current = 100 + current
		}
	}

	if direction == "R" {
		for distance > 99 {
			distance = distance - 100
		}
		current += distance
		if current > 99 {
			current = current - 100
		}
	}
	return current
}
