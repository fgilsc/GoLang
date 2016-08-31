package main

import (
	"fmt"
	"strconv"
)

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

func det(matrix [3][3]int) int { //Sarrus 3x3
	d1 := matrix[0][0] * matrix[1][1] * matrix[2][2]
	s11 := matrix[0][1] * matrix[1][2] * matrix[2][0]
	s12 := matrix[0][2] * matrix[1][0] * matrix[2][1]
	d2 := matrix[0][2] * matrix[1][1] * matrix[2][0]
	s21 := matrix[0][0] * matrix[1][2] * matrix[2][1]
	s22 := matrix[1][0] * matrix[0][1] * matrix[2][2]

	res := (d1 + s11 + s12) - (d2 + s21 + s22)
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
	det := det(A)
	detS := strconv.Itoa(det)
	fmt.Println("El determinante es: " + detS)
}
