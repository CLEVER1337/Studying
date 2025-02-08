package main

import "fmt"

func alg_ext(a [5][5]int, i int, j int) ([4][4]int, error){
	var ae [4][4]int
	n := 0
	k := 0
	for it := 0; it < 5; it++ {
		if it == i {
			n = 1
		}
		k = 0
		for jt := 0; jt < 5; jt++ {
			if jt == j {
				k = 1
			}
			if it + n < 5 && jt + k < 5 {
				ae[it][jt] = a[it + n][jt + k]
			}
		}
	}

	return ae, nil
}

func main(){
	arr := [5][5]int {{1, 2, 3, 4, 5}, {1, 2, 3, 4, 4}, {1, 2, 2, 3, 4}, {2, 3, 4, 5, 6}, {5, 1, 2, 3, 4}}
	
	fmt.Println(arr)

	fmt.Println('\n')

	fmt.Println(alg_ext(arr, 0, 0))
}
