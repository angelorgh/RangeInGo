package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"regexp"
)

func main() {
	Tests()
}

// Range : Definition of a Range struct
type Range struct {
	RangeString   string
	LowerEndPoint int
	UpperEndPoint int
	AllPoints     []int
}

// NewRange : Constructor para un Range
func NewRange(s string) Range {
	var r Range

	s = strings.ReplaceAll(s," ", "")
	re := regexp.MustCompile(`(\[|\()(-?\d+),(-?\d+)(\]|\))`) 
	slice := re.FindAllStringSubmatch(s, -1)
	var err error

	if slice == nil {
		err := errors.New("Input string is not valid")
		log.Fatal(err)
	}

	if slice[0][1] == "(" {
		r.LowerEndPoint, err = strconv.Atoi(slice[0][2])
		r.LowerEndPoint++
	}
	if slice[0][1] == "[" {
		r.LowerEndPoint, err = strconv.Atoi(slice[0][2])
	}

	if slice[0][4] == ")" {
		r.UpperEndPoint, err = strconv.Atoi(slice[0][3])
		r.UpperEndPoint--
	}
	if slice[0][4] == "]" {
		r.UpperEndPoint, err = strconv.Atoi(slice[0][3])
	}

	if err != nil {
		log.Fatal(err)
	}
	if r.UpperEndPoint < r.LowerEndPoint {
		err := errors.New("UpperEndPoint is lower than the LowerEndPoint")
		log.Fatal(err)
	}

	var arr []int
	count := 0

	for i := r.LowerEndPoint; i <= r.UpperEndPoint; i++ {
		arr = append(arr, i)
		count++
	}

	r.AllPoints = arr
	r.RangeString = s

	return r
}

// EndPoints : Returns lower and upper endpoints
func (range1 Range) EndPoints() []int {
	var s []int = []int{range1.LowerEndPoint, range1.UpperEndPoint}
	return s
}

// GetAllPoints : Returns an array with all point of the range
func (range1 Range) GetAllPoints() []int {
	var s []int = range1.AllPoints
	return s
}

// Contains : Checks if a given int is contained in a range
func (range1 Range) Contains(v int) bool {
	for _, n := range range1.AllPoints {
		if n == v {
			return true
		}
	}
	return false
}

// DoesNotContains : Checks if a given int is not contained in a range
func (range1 Range) DoesNotContains(v int) bool {
	return !range1.Contains(v)
}

// ContainsArray : Checks if a given int array is contained in a range
func (range1 Range) ContainsArray(arr []int) bool {

	if len(arr) > len(range1.AllPoints) || len(arr) == 0 {
		return false
	}

	i := 0
	j := 0

	for i = 0; i < len(arr); i++ {
		for j = 0; j < len(range1.AllPoints); j++ {
			if arr[i] == range1.AllPoints[j] {
				break
			}
		}

		if j == len(range1.AllPoints) {
			return false
		}
	}

	return true
}

// DoesNotContainsArray : Checks if a given int array is not contained in a range
func (range1 Range) DoesNotContainsArray(arr []int) bool {
	return !range1.ContainsArray(arr)
}

// ContainsRange : Checks if a given range is contained in another range
func (range1 Range) ContainsRange(r Range) bool {
	return range1.ContainsArray(r.AllPoints)
}

// DoesNotContainsRange : Checks if a given range is not contained in another range
func (range1 Range) DoesNotContainsRange(r Range) bool {
	return !range1.ContainsArray(r.AllPoints)
}

// Equals : Checks if the given range is the same
func (range1 Range) Equals(r Range) bool {
	if range1.LowerEndPoint == r.LowerEndPoint && range1.UpperEndPoint == r.UpperEndPoint {
		return true
	}
	return false
}

// NotEquals : Checks if the given range is not the same
func (range1 Range) NotEquals(r Range) bool {
	return !range1.Equals(r)
}

// Overlaps : Checks if two ranges have some sort of overlap
func (range1 Range) Overlaps(r Range) bool {
	r1Lower := range1.LowerEndPoint
	r2Lower := r.LowerEndPoint
	condition := false

	if r2Lower == r1Lower {
		condition = true
	}
	if r2Lower < r1Lower {
		condition = r.Contains(r1Lower)
	}
	if r1Lower < r2Lower {
		condition = r.Contains(r2Lower)
	}

	return condition
}

// Tests : Dank
func Tests() {

	// NewRange 1
	//r := NewRange("[3,8)")
	//fmt.Println(r.LowerEndPoint, r.UpperEndPoint)

	// NewRange 2
	//r := NewRange("(-10,7]")
	//fmt.Println(r.LowerEndPoint, r.UpperEndPoint)

	// NewRange 3
	//r := NewRange("2(3]")
	//fmt.Println(r.LowerEndPoint, r.UpperEndPoint)

	// NewRange 4
	//r := NewRange("[5,    -6)")
	//fmt.Println(r.LowerEndPoint, r.UpperEndPoint)

	// EndPoints 1
	//r := NewRange("(9, 18)")
	//fmt.Println(r.EndPoints())

	// EndPoints 2
	//r := NewRange("[56, 102)")
	//fmt.Println(r.EndPoints())

	// EndPoints 3
	//r := NewRange("[-23, -7]")
	//fmt.Println(r.EndPoints())

	// EndPoints 4
	//r := NewRange("(-235, -7]")
	//fmt.Println(r.EndPoints())

	// GetAllPoints 1
	//r := NewRange("(1,10]")
	//fmt.Println(r.GetAllPoints())

	// GetAllPoints 2
	//r := NewRange("[0,21]")
	//fmt.Println(r.GetAllPoints())

	// GetAllPoints 3
	//r := NewRange("(-33,-22)")
	//fmt.Println(r.GetAllPoints())

	// GetAllPoints 4
	//r := NewRange("(0,0)")
	//fmt.Println(r.GetAllPoints())

	// Contains 1
    //r := NewRange("(2,10)")
	//fmt.Println(r.Contains(3))

	// Contains 2
    //r := NewRange("(-42,32]")
	//fmt.Println(r.Contains(-32))

	// Contains 3
    //r := NewRange("[3,18]")
	//fmt.Println(r.Contains(21))

	// Contains 4
    //r := NewRange("[-12,-2)")
	//fmt.Println(r.Contains(-2))

	// DoesNotContains 1
	//r := NewRange("(-12,-3]")
	//fmt.Println(r.DoesNotContains(-31))

	// DoesNotContains 2
	//r := NewRange("[4,12]")
	//fmt.Println(r.DoesNotContains(15))

	// DoesNotContains 3
	//r := NewRange("(-22,-1]")
	//fmt.Println(r.DoesNotContains(-3))

	// DoesNotContains 4
	//r := NewRange("[10,500)")
	//fmt.Println(r.DoesNotContains(250))

	// ContainsArray 1
	//r := NewRange("[2,53]")
	//fmt.Println(r.ContainsArray([]int{3,5,8,50}))

	// ContainsArray 2
	//r := NewRange("(-33,22]")
	//fmt.Println(r.ContainsArray([]int{-31,-10,0,10,11,7,-12}))

	// ContainsArray 3
	//r := NewRange("(-123,-52)")
	//fmt.Println(r.ContainsArray([]int{-110, -23, -100, -124, -35}))

	// ContainsArray 4
	//r := NewRange("(10, 51]")
	//fmt.Println(r.ContainsArray([]int{19,23,17,7,52}))

	//DoesNotContainsArray 4
	r := NewRange("(10, 51]")
	fmt.Println(r.DoesNotContainsArray([]int{19,23,17,7,52}))
}