package numberrange

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// Range : Definition of Range Struct
type Range struct{
	RangeString string
	LowerEndPoint int
	UpperEndPoint int
}

// NewRange : Constructor for a Range
func NewRange(s string) (Range, error){
	var r Range
	
	s = strings.ReplaceAll(s, " ", "")
	re := regexp.MustCompile(`(\[|\()(-?\d+),(-?\d+)(\]|\))`)
	slice := re.FindAllStringSubmatch(s, -1)

	if slice == nil {
		return Range{}, errors.New("Input string is not valid")
	}

	if slice[0][1] == "(" {
		r.LowerEndPoint, _ = strconv.Atoi(slice[0][2])
		r.LowerEndPoint++
	}
	if slice[0][1] == "[" {
		r.LowerEndPoint, _ = strconv.Atoi(slice[0][2])
	}

	if slice[0][4] == ")" {
		r.UpperEndPoint, _ = strconv.Atoi(slice[0][3])
		r.UpperEndPoint--
	}
	if slice[0][4] == "]" {
		r.UpperEndPoint, _ = strconv.Atoi(slice[0][3])
	}

	if r.UpperEndPoint <= r.LowerEndPoint {
		return Range{}, errors.New("Inverted range")
	}

	r.RangeString = s
	return r, nil
}

// EndPoints : Returns an array that includes lower and upper endpoints
func (range1 Range) EndPoints() []int{
	return []int{range1.LowerEndPoint, range1.UpperEndPoint}
}

// GetAllPoints : Returns an array that includes all point from a range
func (range1 Range) GetAllPoints() []int{
	var a []int

	for i := range1.LowerEndPoint; i <= range1.UpperEndPoint; i++ {
		a = append(a, i)
	}

	return a
}

// Contains : Checks if an int is contained in the range
func (range1 Range) Contains(v int) bool{
	arr := range1.GetAllPoints()

	for _, n := range arr {
		if n == v {
			return true
		}
	}

	return false
}

// DoesNotContains : Checks if an int is not contained in a range
func (range1 Range) DoesNotContains(v int) bool{
	return !range1.Contains(v)
}

// ContainsArray : Checks if int array is contained in a range
func (range1 Range) ContainsArray(arr []int) bool{
	a := range1.GetAllPoints()

	if len(arr) > len(a) || len(arr) == 0 {
		return false
	}

	i := 0
	j := 0

	for i = 0; i < len(arr); i++ {
		for j = 0; j < len(a); j++ {
			if arr[i] == a[j] {
				break;
			}
		}

		if j == len(a) {
			return false
		}
	}
	return true;
}

// DoesNotContainsArray : Checks if int array is not contained in a range
func (range1 Range) DoesNotContainsArray(arr []int) bool{
	return !range1.ContainsArray(arr)
}

// ContainsRange : Checks if a given range is contained in another range
func (range1 Range) ContainsRange(r Range) bool{
	return range1.ContainsArray(r.GetAllPoints())
}

// DoesNotContainsRange : Checks if a given range is not contained in another range
func (range1 Range) DoesNotContainsRange(r Range) bool{
	return !range1.ContainsRange(r)
}

// Equals : Checks if the given range is equal
func (range1 Range) Equals(r Range) bool{
	if range1.LowerEndPoint == r.LowerEndPoint && range1.UpperEndPoint == r.UpperEndPoint {
		return true
	}
	return false
}

// NotEquals : Checks if the given range is not equal
func (range1 Range) NotEquals(r Range) bool{
	return !range1.Equals(r)
}

// Overlaps : Checks if two ranges have some sort of overlap
func (range1 Range) Overlaps(r Range) bool{
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
		condition = range1.Contains(r2Lower)
	}

	return condition
}