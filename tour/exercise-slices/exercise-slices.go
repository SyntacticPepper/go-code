package main

import "golang.org/x/tour/pic"

/* https://tour.golang.org/moretypes/18 */
func Pic(dx, dy int) [][]uint8 {
	var pic = make([][]uint8, dy)
	// allocate the inner slices
	for i := range pic {
		pic[i] = make([]uint8, dx)
	}

	//draw
	for y := range pic {
		for x := range pic[y] {
			//pic[y][x] = uint8((x+y)/2)	//cool
			//pic[y][x] = uint8(x*y)   //nice
			pic[y][x] = uint8(x ^ y) //wow
		}
	}

	return pic
}

func main() {
	pic.Show(Pic)
}
