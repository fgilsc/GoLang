package main

import "fmt"

func suma(matrixA [3][3]int, matrixB [3][3]int) [3][3]int {
	var res [3][3]int

	for i := 0; i < len(res); i++ {
		sub := res[i]
		for j := 0; j < len(sub); j++ {
			n := matrixA[i][j] + matrixB[i][j]
			res[i][j] = n
		}
	}
	return res

}

func resta(matrixA [3][3]int, matrixB [3][3]int) [3][3]int {
	var res [3][3]int

	for i := 0; i < len(res); i++ {
		sub := res[i]
		for j := 0; j < len(sub); j++ {
			n := matrixA[i][j] - matrixB[i][j]
			res[i][j] = n
		}
	}
	return res

}

func obtieneColumna(n int, matrix [3][3]int) [3]int {
	var res [3]int
	for i := 0; i < len(matrix); i++ { //recorreFila
		a := matrix[i][n]
		res[i] = a
	}
	return res
}

func multiplica(matrixA [3][3]int, matrixB [3][3]int) [3][3]int {
	var res [3][3]int
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res); j++ {
			fila := matrixA[i]
			column := obtieneColumna(j, matrixB)
			sum := 0
			for k := 0; k < len(column); k++ {
				t := fila[k] * column[k]
				sum = sum + t
				res[i][j] = sum
			}
		}
	}
	return res

}

//Testing Area
func main() {
	A := [3][3]int{{1, 2, 6}, {9, 5, 1}, {4, 2, 8}}
	B := [3][3]int{{3, 2, 1}, {2, 3, 4}, {5, 6, 1}}
	C := multiplica(A, B)

	for fila := 0; fila < len(C); fila++ {
		fmt.Println(C[fila])
	}
}
