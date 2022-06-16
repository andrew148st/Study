package main

import "fmt"

type rectangle struct {
	height int
	width  int
}

// 									Периметр или площадь
// если прямоугольник является квадратом - вернуть его периметр, иначе - вернуть площадь
func (r rectangle) perimeterOrArea() int {
	if r.width == r.height {
		return 2*r.height + 2*r.width
	} else {
		return r.height * r.width
	}
}
func (r rectangle) aspectRatio() rectangle {
	return rectangle{height: r.height, width: r.height * 16 / 9}
}

func main() {
	r1 := rectangle{height: 5, width: 7}
	r2 := rectangle{height: 8, width: 8}
	r3 := rectangle{height: 1080, width: 212}

	area := r1.perimeterOrArea()
	perimeter := r2.perimeterOrArea()
	aspectRatio := r3.aspectRatio()
	fmt.Printf("периметр равен %d\n", perimeter)
	fmt.Printf("площадь равна %d\n", area)
	fmt.Printf("высота равна %d, ширина равна %d\n", aspectRatio.height, aspectRatio.width)

}
