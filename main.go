package main

import (
	"errors"
	"sync"
)

func main() {
	//mat := [][]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 4}, {1, 2, 2, 3, 4}, {2, 3, 4, 5, 6}, {5, 1, 2, 3, 4}}
	// mat1 := [][]int{{1, 2, 4}, {2, 2, 4}, {5, 2, 6}}
	// fmt.Println(mat1)

	// fmt.Println(laplas(mat1))
	// fmt.Println(power(-1, 2))
	// fmt.Println(trace(mat1))
	// fmt.Println(transposition(mat1))
	//fmt.Println(power(-2, 3))
	// var int_ch chan int = make(chan int)
	// go func() {
	// 	defer close(int_ch)
	// 	//int_ch <- 52
	// }()
	// fmt.Println(<-int_ch)

}

func power(a, b int) int {
	var wg sync.WaitGroup

	c := a
	for i := 0; i < b-1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c *= a
		}()
	}
	wg.Wait()
	return c
}

func minor(mat [][]int, i int, j int) ([][]int, error) {
	var alg_add_mat [][]int
	n := 0
	k := 0
	for it := 0; it < len(mat)-1; it++ {
		if len(mat) != len(mat[it]) {
			return [][]int{}, errors.New("Matrix isn't square")
		}
		if it == i {
			n = 1
		}
		k = 0
		var tmp []int
		for jt := 0; jt < len(mat[it]); jt++ {
			if jt == j {
				k = 1
			}
			if it+n < len(mat) && jt+k < len(mat[jt]) {
				tmp = append(tmp, mat[it+n][jt+k])
			}
		}
		alg_add_mat = append(alg_add_mat, tmp)
	}

	return alg_add_mat, nil
}

func laplas(mat [][]int) (int, error) {
	var wg sync.WaitGroup

	if len(mat) == 1 {
		return mat[0][0], nil
	}
	if len(mat) == 2 {
		return (mat[0][0] * mat[1][1]) - (mat[0][1] * mat[1][0]), nil
	}
	deter := 0
	for index, _ := range mat[0] {
		ad, err := minor(mat, 0, index)
		if err != nil {
			return 0, err
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			var lp int
			lp, err = laplas(ad)
			deter += mat[0][index] * power(-1, 2+index) * lp
		}()
	}

	wg.Wait()
	return deter, nil
}

func trace(mat [][]int) (int, error) {
	var wg sync.WaitGroup
	tmp := 0
	for i := 0; i < len(mat); i++ {
		if len(mat) != len(mat[i]) {
			return 0, errors.New("Matrix isn't square")
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < len(mat[i]); j++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					if i == j {
						tmp += mat[i][j]
					}
				}()
			}
		}()
	}

	wg.Wait()
	return tmp, nil
}

func transposition(mat [][]int) ([][]int, error) {
	var transpode [][]int
	for i := 0; i < len(mat); i++ {
		if len(mat) != len(mat[i]) {
			return [][]int{}, errors.New("Matrix isn't square")
		}
		var tmp []int
		for j := 0; j < len(mat[i]); j++ {
			tmp = append(tmp, mat[j][i])
		}
		transpode = append(transpode, tmp)
	}
	return transpode, nil
}
