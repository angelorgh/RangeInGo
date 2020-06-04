package rang

import (
	"strconv"
	"strings"
)

// Rang : Declarar variables del objeto Range
type rang struct {
	RangeString string
	LowerEndPoint int
	UpperEndPoint int
	AllPoints []int
}

// New : Constructor para un Range
func New(s string) rang {

	var r rang
	s = strings.Trim(s," ")
	
	var split =  strings.Split(s,",")

	if strings.Contains(split[0],"(") {
		f := strings.Trim(split[0], "(")
		r.LowerEndPoint, _ = strconv.Atoi(f)
		r.LowerEndPoint++
	}
	if strings.Contains(split[0],"[") {
		f := strings.Trim(split[0], "[")
		r.LowerEndPoint, _ = strconv.Atoi(f)
	}

	if strings.Contains(split[1],")") {
		f := strings.Trim(split[1], ")")
		r.UpperEndPoint, _ = strconv.Atoi(f)
		r.UpperEndPoint--
	}
	if strings.Contains(split[1],"]") {
		f := strings.Trim(split[1], "]")
		r.UpperEndPoint, _ = strconv.Atoi(f)
	}
	
	var arr []int
	count := 0

	for i := r.LowerEndPoint; i <= r.UpperEndPoint; i++ {
		arr = append(arr, i)
		count++
	}

	r.AllPoints = arr

	return r
}

func (range1 rang) endPoints() []int{
	var s []int =  []int{range1.LowerEndPoint, range1.UpperEndPoint}
	return s
}

