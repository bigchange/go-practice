/*
 * @Author: Jerry You
 * @CreatedDate: 2023/3/28
 * @Last Modified by: lolaliva
 * @Last Modified time: 2023/3/28 11:04
 */

// Package leetcode
// training from: https://www.codewars.com/kata/52a382ee44408cea2500074c/train/go
package leetcode

import "fmt"

// Determinant
//
// a b c d
// e f g h
// i j k l
// m n o p
func Determinant(matrix [][]int) int {
	// your code here
	printMatrix(matrix)
	var f func([][]int) int
	f = func(matrix [][]int) int {
		sum := 0
		rowLen := len(matrix)
		// colLen := len(matrix)
		if rowLen == 1 {
			return matrix[0][0]
		}
		if rowLen == 2 {
			return cal2Ddet(matrix)
		}
		rows := matrix[0]
		for i, v := range rows {
			newM := getNewMatrix(matrix, 0, i)
			fmt.Printf("v:%v, \nmatrix:\n", v)
			printMatrix(newM)
			if i % 2 == 0 {
				sum += v * f(newM)
			} else {
				sum -= v * f(newM)
			}
		}
		return sum
	}
	return f(matrix)

}


// matrix 需要使用新的低维度
func cal2Ddet(matrix [][]int) int {
	row1 := matrix[0]
	row2 := matrix[1]
	fmt.Printf("2Dmatrix:\n")
	printMatrix(matrix)
	v := row1[0] * row2[1] - row1[1] * row2[0]
	fmt.Println("value:", v)
	return v
}

// 从源matrix构造低维度matrix
func getNewMatrix(matrix [][]int, row, col int) [][]int {
	newMatrix := make([][]int, len(matrix) - 1)
	rowIndex := 0
	for i,rows := range matrix {
		if i == row {
			continue
		}
		cols := make([]int, len(matrix[0]) - 1)
		colIndex := 0
		for j, v := range rows {
			if j == col {
				continue
			}
			cols[colIndex] = v
			colIndex++
		}
		newMatrix[rowIndex] = cols
		rowIndex++
	}
	return newMatrix
}

func printMatrix(matrix [][]int) {
	for _, rows := range matrix {
		for _, col := range rows {
			fmt.Printf("%v\t", col)
		}
		fmt.Println()
	}
	fmt.Println("ending ")
}

