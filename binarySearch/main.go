package main

// 										"Binary search"

func binaryIncludes(slice []int, lookingFor int) bool {
	if len(slice) == 1 {
		if slice[0] == lookingFor {
			return true
		} else {
			return false
		}
	}
	mid := len(slice) / 2
	midValue := slice[mid]
	if midValue == lookingFor {
		return true
	} else if midValue > lookingFor {
		return binaryIncludes(slice[:mid], lookingFor)
	} else {
		return binaryIncludes(slice[mid:], lookingFor)
	}
}

// [1, 3, 9, 15, 25]
// ищем 15
// len = 5
// mid = 2
// midValue = 9

// [15, 25]
// len = 2
// mid = 1
// midValue = 25

// [15]
// len = 1
// mid = 0
// midValue = 15
// нашли!
