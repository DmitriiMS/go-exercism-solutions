//package matrix provides functions to get matrix rows and colums, and also to set individual values of elements
package matrix

import (
	"errors"
	"strconv"
	"strings"
)

//matrix consists of matrix body and two variables, representing its dimensions
type Matrix struct {
	body         [][]int //TODO: rewrite in 1D format, this will simplify things
	numberOfRows int
	rowLength    int
}

//New() constructs a new Matrix object from given string representation
func New(s string) (*Matrix, error) {
	stringRows := strings.Split(s, "\n")                    //get individual strings that represent rows
	splitRows, errValidation := validateStrings(stringRows) // validate strings and get the text representation of the matrix
	if errValidation != nil {                               //if something went wrong during the validation, return error
		return nil, errValidation
	}
	numberOfRows, rowLength := len(splitRows), len(splitRows[0]) // get matrix dimensions
	tempMatrix := make([][]int, numberOfRows)                    //temp int matrix
	//walk over the text matrix and convert each element to int
	//if element can't be converted to int, return error
	for i := 0; i < numberOfRows; i++ {
		tempMatrix[i] = make([]int, rowLength)
		for j := 0; j < rowLength; j++ {
			number, errConversion := strconv.Atoi(splitRows[i][j]) //if element can't be converted to int, return error
			if errConversion != nil {
				return nil, errConversion
			}
			tempMatrix[i][j] = number
		}
	}
	return &Matrix{tempMatrix, numberOfRows, rowLength}, nil // construct and return matrix
}

//Rows() constructs a copy of matrix body, which is stored in rows-columns format, and returns it.
func (m *Matrix) Rows() [][]int {
	temp := make([][]int, m.numberOfRows)
	for i := 0; i < m.numberOfRows; i++ {
		temp[i] = make([]int, m.rowLength)
		for j := 0; j < m.rowLength; j++ {
			temp[i][j] = m.body[i][j]
		}
	}
	return temp
}

//Cols() returns the body of the transposed matrix, this way we get columns of the original matrix.
func (m *Matrix) Cols() [][]int {
	return m.constructTransposed().body
}

//Set() sets a new value of a given element in the matrix.
func (m *Matrix) Set(row, column, value int) bool {
	//check if given indexes are not out of bounds for the given matrix
	if (row < 0 || row > m.numberOfRows-1) ||
		(column < 0 || column > m.rowLength-1) {
		return false
	}
	m.body[row][column] = value // if everything is ok, set a new value.
	return true
}

//constructTransposed() transposes matrix, creating a new one where rows are columns and columns are rows
func (m *Matrix) constructTransposed() *Matrix {
	tempMatrix := make([][]int, m.rowLength)
	for i := 0; i < m.rowLength; i++ {
		tempMatrix[i] = make([]int, m.numberOfRows)
		for j := 0; j < m.numberOfRows; j++ {
			tempMatrix[i][j] = m.body[j][i]
		}
	}
	return &Matrix{tempMatrix, m.rowLength, m.numberOfRows}
}

//ValidateStrings() validates strings that represent matrix rows, splits them into elements and returns
//the text representation of the matrix.
func validateStrings(Rows []string) ([][]string, error) {
	splitRows := [][]string{}
	//trim the first string and split it into elements, get the lentgh of the resulting slice
	//this is the length against which lengths of other strings will be compared
	prevLen := len(strings.Split(strings.TrimSpace(Rows[0]), " "))
	for i, _ := range Rows {
		//trim each string, split it into elements and compare length of the resulting slice with the first measurement
		splitRow := strings.Split(strings.TrimSpace(Rows[i]), " ")
		currLen := len(splitRow)
		if currLen != prevLen { // if this row is of different length than the first one, return error
			return nil, errors.New("rows are of uneven length")
		}
		splitRows = append(splitRows, splitRow) // add sliced text matrix row to the text matrix
	}
	return splitRows, nil
}
