//package matrix provides functions to get matrix rows and colums, and also to set individual values of elements
package matrix

import (
	"errors"
	"strconv"
	"strings"
)

//matrix consists of matrix body and length of its rows
type Matrix struct {
	body      []int
	rowLength int
}

//New() constructs a new Matrix object from given string representation
func New(s string) (*Matrix, error) {
	stringRows := strings.Split(s, "\n")                               //get individual strings that represent rows
	tempBody, rowLength, errValidation := validateAndBuild(stringRows) // validate strings and construct the body of a matrix
	if errValidation != nil {                                          //if something went wrong during the validation, return error
		return nil, errValidation
	}
	return &Matrix{tempBody, rowLength}, nil // construct and return matrix
}

//Rows() constructs a rows-first representation of a matrix. It should be just like in the original input.
func (m *Matrix) Rows() [][]int {
	temp := make([][]int, 0)
	for start, end := 0, m.rowLength; end <= len(m.body); start, end = start+m.rowLength, end+m.rowLength {
		temp = append(temp, m.body[start:end])
	}
	//create a copy, so our representation won't be referencing matrix structure directly
	cpy := make([][]int, len(m.body)/m.rowLength)
	for i := range cpy {
		cpy[i] = make([]int, m.rowLength)
		copy(cpy[i], temp[i])
	}
	return cpy
}

//Cols() constructs a columns-first representation of a matrix. It should look like original input, but transposed.
func (m *Matrix) Cols() [][]int {
	temp := make([][]int, 0)
	for i := 0; i < m.rowLength; i++ {
		col := make([]int, 0)
		for j := i; j < len(m.body); j += m.rowLength {
			col = append(col, m.body[j])
		}
		temp = append(temp, col)
	}
	return temp
}

//Set() sets a new value of a given element in the matrix.
func (m *Matrix) Set(row, column, value int) bool {
	//check if given indexes are not out of bounds for the given matrix
	if (row < 0 || row >= len(m.body)/m.rowLength) ||
		(column < 0 || column >= m.rowLength) {
		return false
	}
	m.body[row*m.rowLength+column] = value // if everything is ok, set a new value.
	return true
}

//ValidateStrings() validates strings that represent matrix rows, splits them into elements, converts elements to int
//and returns the body of a matrix represented by a 1D slice.
func validateAndBuild(Rows []string) ([]int, int, error) {
	longForm := []int{}
	//trim the first string and split it into elements, get the lentgh of the resulting slice
	//this is the length against which lengths of other strings will be compared
	prevLen := len(strings.Split(strings.TrimSpace(Rows[0]), " "))
	for i := range Rows {
		//trim each string, split it into elements and compare length of the resulting slice with the first measurement
		splitRow := strings.Split(strings.TrimSpace(Rows[i]), " ")
		currLen := len(splitRow)
		if currLen != prevLen { // if this row is of different length than the first one, return error
			return nil, 0, errors.New("rows are of uneven length")
		}
		for _, number := range splitRow { //walk over a slie of elements, convert them to int and add each of them to the matrix slice
			numberInt, errConversion := strconv.Atoi(number) //if element can't be converted to int, return error
			if errConversion != nil {
				return nil, 0, errConversion
			}
			longForm = append(longForm, numberInt)
		}
	}
	return longForm, prevLen, nil
}
