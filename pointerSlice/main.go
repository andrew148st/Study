package main

func main() {
	pointerSlice := []*int{}
	fillSlice(pointerSlice)
}

func fillSlice(slice []*int) {
	for i := 0; i < 9; i++ {
		slice = append(slice, &i)
	}
}
