package numberrange

import (
	"reflect"
	"testing"
)

func TestNewRangePasses1(t *testing.T){
	want := Range{"[3,8)", 3, 7}
	input := "[3,8)"
	have, err := NewRange(input)

	if err != nil {
		t.Errorf("NewRange() with args %v : FAILED, expected a value but got error '%v'", 
		input, err)
		t.Fatal()
	} 

	if reflect.DeepEqual(want, have) {
		t.Logf("NewRange() with args %v : PASSED, expected %v and got '%v'", input, 
		want, have)
	} else {
		t.Errorf("NewRange() with args %v : FAILED, expected %v but got '%v'", input, 
		want, have)
		t.Fatal()
	}
}

func TestNewRangePasses2(t *testing.T){
	want := Range{"(-10,7]", -9, 7}
	input := "(-10,7]"
	have, err := NewRange(input)

	if err != nil {
		t.Errorf("NewRange() with args %v : FAILED, expected a value but got error '%v'", 
		input, err)
		t.Fatal()
	} 

	if reflect.DeepEqual(want, have) {
		t.Logf("NewRange() with args %v : PASSED, expected %v and got '%v'", input, 
		want, have)
	} else {
		t.Errorf("NewRange() with args %v : FAILED, expected %v but got '%v'", input, 
		want, have)
		t.Fatal()
	}
}

func TestNewRangeFailsIncorrectString(t *testing.T){
	want := Range{}
	input := "2(3]"
	have, err := NewRange(input)

	if reflect.DeepEqual(want, have) && err != nil {
		t.Logf("NewRange() with args %v : PASSED, expected an error and got '%v'", 
		input, err)
	} else {
		t.Errorf("NewRange() with args %v : FAILED, expected an error but got '%v'", 
		input, have)
		t.Fatal()
	}
}

func TestNewRangeFailsInvertedRange(t *testing.T){
	want := Range{}
	input := "[5,-6)"
	have, err := NewRange(input)

	if reflect.DeepEqual(want, have) && err != nil {
		t.Logf("NewRange() with args %v : PASSED, expected an error and got '%v'", 
		input, err)
	} else {
		t.Errorf("NewRange() with args %v : FAILED, expected an error but got '%v'", 
		input, have)
		t.Fatal()
	}
}

func TestEndPointsPasses1(t *testing.T){
	want := []int{10,17}
	input := Range{"(9,18)", 10, 17}
	have := input.EndPoints()

	if reflect.DeepEqual(want, have) {
		t.Logf("EndPoints() with : PASSED, expected value %v and got '%v'", 
		want, have)
	} else {
		t.Errorf("EndPoints() with : FAILED, expected value %v but got '%v'", 
		want, have)
		t.Fatal()
	}
}

func TestEndPointsPasses2(t *testing.T){
	want := []int{56,101}
	input := Range{"[56,102)", 56, 101}
	have := input.EndPoints()

	if reflect.DeepEqual(want, have) {
		t.Logf("EndPoints() with : PASSED, expected value %v and got '%v'", 
		want, have)
	} else {
		t.Errorf("EndPoints() with : FAILED, expected value %v but got '%v'", 
		want, have)
		t.Fatal()
	}
}

func TestEndPointsPasses3(t *testing.T){
	want := []int{-23,-7}
	input := Range{"[-23,-7]",-23,-7}
	have := input.EndPoints()

	if reflect.DeepEqual(want, have) {
		t.Logf("EndPoints() with : PASSED, expected value %v and got '%v'", 
		want, have)
	} else {
		t.Errorf("EndPoints() with : FAILED, expected value %v but got '%v'", 
		want, have)
		t.Fatal()
	}
}

func TestEndPointsPasses4(t *testing.T){
	want := []int{-236,-7}
	input := Range{"(-235,-7]",-236,-7}
	have := input.EndPoints()

	if reflect.DeepEqual(want, have) {
		t.Logf("EndPoints() with : PASSED, expected value %v and got '%v'", 
		want, have)
	} else {
		t.Errorf("EndPoints() with : FAILED, expected value %v but got '%v'", 
		want, have)
		t.Fatal()
	}
}

func TestGetAllPointsPasses1(t *testing.T){
	input := Range{"(1,10]", 2, 10}
	want := []int{2,3,4,5,6,7,8, 9,10}
	have := input.GetAllPoints()

	if reflect.DeepEqual(want, have) {
		t.Logf("GetAllPoints() : PASSED, expected %v and got value '%v'", want, have)
	} else {
		t.Errorf("GetAllPoints() : PASSED, expected %v and but value '%v'", want, have)
		t.Fatal()
	}
}

func TestGetAllPointsPasses2(t *testing.T){
	input := Range{"[0,21]", 0, 21}
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
		 15, 16, 17, 18, 19, 20, 21}
	have := input.GetAllPoints()

	if reflect.DeepEqual(want, have) {
		t.Logf("GetAllPoints() : PASSED, expected %v and got value '%v'", want, have)
	} else {
		t.Errorf("GetAllPoints() : PASSED, expected %v and but value '%v'", want, have)
		t.Fatal()
	}
}

func TestGetAllPointsPasses3(t *testing.T){
	input := Range{"(-33,-22)", -32, -23}
	want := []int{-32, -31, -30, -29, -28, -27, -26, -25, -24, -23}
	have := input.GetAllPoints()

	if reflect.DeepEqual(want, have) {
		t.Logf("GetAllPoints() : PASSED, expected %v and got value '%v'", want, have)
	} else {
		t.Errorf("GetAllPoints() : PASSED, expected %v and but value '%v'", want, have)
		t.Fatal()
	}
}

func TestGetAllPointsPasses4(t *testing.T){
	input := Range{"[-12,-4]", -12, -4}
	want := []int{-12, -11, -10, -9, -8,-7, -6, -5, -4}
	have := input.GetAllPoints()

	if reflect.DeepEqual(want, have) {
		t.Logf("GetAllPoints() : PASSED, expected %v and got value '%v'", want, have)
	} else {
		t.Errorf("GetAllPoints() : PASSED, expected %v and but value '%v'", want, have)
		t.Fatal()
	}
}

func TestContainsTruePasses1(t *testing.T){
	inputrange := Range{"(2,10)", 3, 9}
	inputint := 3
	want := true
	have := inputrange.Contains(inputint)

	if have == want{
		t.Logf("Contains() with args %v : PASSED, expected %v and got value '%v'",
		inputint, want, have)
	} else {
		t.Logf("Contains() with args %v : FAILED, expected %v but got value '%v'",
		inputint, want, have)
		t.Fatal()
	}
}

func TestContainsTruePasses2(t *testing.T){
	inputrange := Range{"(-42,32]", -41, 32}
	inputint := -32
	want := true
	have := inputrange.Contains(inputint)

	if have == want{
		t.Logf("Contains() with args %v : PASSED, expected %v and got value '%v'",
		inputint, want, have)
	} else {
		t.Logf("Contains() with args %v : FAILED, expected %v but got value '%v'",
		inputint, want, have)
		t.Fatal()
	}
}

func TestContainsFalsePasses1(t *testing.T){
	inputrange := Range{"[3,18]", 3, 18}
	inputint := 21
	want := false
	have := inputrange.Contains(inputint)

	if have == want{
		t.Logf("Contains() with args %v : PASSED, expected %v and got value '%v'",
		inputint, want, have)
	} else {
		t.Logf("Contains() with args %v : FAILED, expected %v but got value '%v'",
		inputint, want, have)
		t.Fatal()
	}
}

func TestContainsFalsePasses2(t *testing.T){
	inputrange := Range{"[-12,-2)", -12, -3}
	inputint := -2
	want := false
	have := inputrange.Contains(inputint)

	if have == want{
		t.Logf("Contains() with args %v : PASSED, expected %v and got value '%v'",
		inputint, want, have)
	} else {
		t.Logf("Contains() with args %v : FAILED, expected %v but got value '%v'",
		inputint, want, have)
		t.Fatal()
	}
}

func TestDoesNotContainsFalsePasses1(t *testing.T){
	inputrange := Range{"(2,10)", 3, 9}
	inputint := 3
	want := false
	have := inputrange.DoesNotContains(inputint)

	if have == want{
		t.Logf("DoesNotContains() with args %v : PASSED, expected %v and got value '%v'",
		inputint, want, have)
	} else {
		t.Logf("DoesNotContains() with args %v : FAILED, expected %v but got value '%v'",
		inputint, want, have)
		t.Fatal()
	}
}

func TestDoesNotContainsFalsePasses2(t *testing.T){
	inputrange := Range{"(-42,32]", -41, 32}
	inputint := -32
	want := false
	have := inputrange.DoesNotContains(inputint)

	if have == want{
		t.Logf("DoesNotContains() with args %v : PASSED, expected %v and got value '%v'",
		inputint, want, have)
	} else {
		t.Logf("DoesNotContains() with args %v : FAILED, expected %v but got value '%v'",
		inputint, want, have)
		t.Fatal()
	}
}

func TestDoesNotContainsTruePasses1(t *testing.T){
	inputrange := Range{"[3,18]", 3, 18}
	inputint := 21
	want := true
	have := inputrange.DoesNotContains(inputint)

	if have == want{
		t.Logf("DoesNotContains() with args %v : PASSED, expected %v and got value '%v'",
		inputint, want, have)
	} else {
		t.Logf("DoesNotContains() with args %v : FAILED, expected %v but got value '%v'",
		inputint, want, have)
		t.Fatal()
	}
}

func TestDoesNotContainsTruePasses2(t *testing.T){
	inputrange := Range{"[-12,-2)", -12, -3}
	inputint := -2
	want := true
	have := inputrange.DoesNotContains(inputint)

	if have == want{
		t.Logf("DoesNotContains() with args %v : PASSED, expected %v and got value '%v'",
		inputint, want, have)
	} else {
		t.Logf("DoesNotContains() with args %v : FAILED, expected %v but got value '%v'",
		inputint, want, have)
		t.Fatal()
	}
}


func TestContainsArrayTruePasses1(t *testing.T){
	inputrange := Range{"[2,53]", 2, 53}
	inputarray := []int{3,5,8,50}
	want := true
	have := inputrange.ContainsArray(inputarray)

	if have == want {
		t.Logf("ContainsArray() with args %v : PASSED, expected %v and got value '%v'", inputarray, want, have)
	} else {
		t.Errorf("ContainsArray() with args %v : FAILED, expected %v but got value '%v'", inputarray, want, have)
		t.Fatal()
	}
}

func TestContainsArrayTruePasses2(t *testing.T){
	inputrange := Range{"(-33,22]", -32, 22}
	inputarray := []int{-31, -10, 0, 10, 11 , 7, -12}
	want := true
	have := inputrange.ContainsArray(inputarray)

	if have == want {
		t.Logf("ContainsArray() with args %v : PASSED, expected %v and got value '%v'", inputarray, want, have)
	} else {
		t.Errorf("ContainsArray() with args %v : FAILED, expected %v but got value '%v'", inputarray, want, have)
		t.Fatal()
	}
}

func TestContainsArrayFalsePasses1(t *testing.T){
	inputrange := Range{"(-123,-52)", -123, -52}
	inputarray := []int{-110,-23,-100,-124,-35}
	want := false
	have := inputrange.ContainsArray(inputarray)

	if have == want {
		t.Logf("ContainsArray() with args %v : PASSED, expected %v and got value '%v'", inputarray, want, have)
	} else {
		t.Errorf("ContainsArray() with args %v : FAILED, expected %v but got value '%v'", inputarray, want, have)
		t.Fatal()
	}
}

func TestContainsArrayFalsePasses2(t *testing.T){
	inputrange := Range{"(10,51]", 11, 51}
	inputarray := []int{19, 23, 17, 7, 52}
	want := false
	have := inputrange.ContainsArray(inputarray)

	if have == want {
		t.Logf("ContainsArray() with args %v : PASSED, expected %v and got value '%v'", inputarray, want, have)
	} else {
		t.Errorf("ContainsArray() with args %v : FAILED, expected %v but got value '%v'", inputarray, want, have)
		t.Fatal()
	}
}

func TestDoesNotContainsArrayFalsePasses1(t *testing.T){
	inputrange := Range{"[2,53]", 2, 53}
	inputarray := []int{3,5,8,50}
	want := false
	have := inputrange.DoesNotContainsArray(inputarray)

	if have == want {
		t.Logf("DoesNotContainsArray() with args %v : PASSED, expected %v and got value '%v'", inputarray, want, have)
	} else {
		t.Errorf("DoesNotContainsArray() with args %v : FAILED, expected %v but got value '%v'", inputarray, want, have)
		t.Fatal()
	}
}

func TestDoesNotContainsArrayFalsePasses2(t *testing.T){
	inputrange := Range{"(-33,22]", -32, 22}
	inputarray := []int{-31, -10, 0, 10, 11 , 7, -12}
	want := false
	have := inputrange.DoesNotContainsArray(inputarray)

	if have == want {
		t.Logf("DoesNotContainsArray() with args %v : PASSED, expected %v and got value '%v'", inputarray, want, have)
	} else {
		t.Errorf("DoesNotContainsArray() with args %v : FAILED, expected %v but got value '%v'", inputarray, want, have)
		t.Fatal()
	}
}

func TestDoesNotContainsArrayTruePasses1(t *testing.T){
	inputrange := Range{"(-123,-52)", -123, -52}
	inputarray := []int{-110,-23,-100,-124,-35}
	want := true
	have := inputrange.DoesNotContainsArray(inputarray)

	if have == want {
		t.Logf("DoesNotContainsArray() with args %v : PASSED, expected %v and got value '%v'", inputarray, want, have)
	} else {
		t.Errorf("DoesNotContainsArray() with args %v : FAILED, expected %v but got value '%v'", inputarray, want, have)
		t.Fatal()
	}
}

func TestDoesNotContainsArrayTruePasses2(t *testing.T){
	inputrange := Range{"(10,51]", 11, 51}
	inputarray := []int{19, 23, 17, 7, 52}
	want := true
	have := inputrange.DoesNotContainsArray(inputarray)

	if have == want {
		t.Logf("DoesNotContainsArray() with args %v : PASSED, expected %v and got value '%v'", inputarray, want, have)
	} else {
		t.Errorf("DoesNotContainsArray() with args %v : FAILED, expected %v but got value '%v'", inputarray, want, have)
		t.Fatal()
	}
}

func TestContainsRangeTruePasses1(t *testing.T){
	inputrange := Range{"(2,10]", 3, 10}
	inputrange2 := Range{"(4,6)", 5, 5}
	want := true
	have := inputrange.ContainsRange(inputrange2)

	if want == have {
		t.Logf("ContainsRange() with args %v : PASSED, expected %v and got value '%v'",
		inputrange2, want, have)
	} else {
		t.Errorf("ContainsRange() with args %v : FAILED, expected %v but got value '%v'",
		inputrange2, want, have)
		t.Fatal()
	}
}

func TestContainsRangeTruePasses2(t *testing.T){
	inputrange := Range{"[-10,3)", -10, 2}
	inputrange2 := Range{"[-10,2]", -10, 2}
	want := true
	have := inputrange.ContainsRange(inputrange2)

	if want == have {
		t.Logf("ContainsRange() with args %v : PASSED, expected %v and got value '%v'",
		inputrange2, want, have)
	} else {
		t.Errorf("ContainsRange() with args %v : FAILED, expected %v but got value '%v'",
		inputrange2, want, have)
		t.Fatal()
	}
}

func TestContainsRangeFalsePasses1(t *testing.T){
	inputrange := Range{"(50,160)", 51, 159}
	inputrange2 := Range{"[150,200]", 150, 200}
	want := false
	have := inputrange.ContainsRange(inputrange2)

	if want == have {
		t.Logf("ContainsRange() with args %v : PASSED, expected %v and got value '%v'",
		inputrange2, want, have)
	} else {
		t.Errorf("ContainsRange() with args %v : FAILED, expected %v but got value '%v'",
		inputrange2, want, have)
		t.Fatal()
	}
}

func TestContainsRangeFalsePasses2(t *testing.T){
	inputrange := Range{"[-45,32)", -45, 31}
	inputrange2 := Range{"(-55,0]", -56, 0}
	want := false
	have := inputrange.ContainsRange(inputrange2)

	if want == have {
		t.Logf("ContainsRange() with args %v : PASSED, expected %v and got value '%v'",
		inputrange2, want, have)
	} else {
		t.Errorf("ContainsRange() with args %v : FAILED, expected %v but got value '%v'",
		inputrange2, want, have)
		t.Fatal()
	}
}

func TestDoesNotContainsRangeFalsePasses1(t *testing.T){
	inputrange := Range{"(2,10]", 3, 10}
	inputrange2 := Range{"(4,6)", 5, 5}
	want := false
	have := inputrange.DoesNotContainsRange(inputrange2)

	if want == have {
		t.Logf("DoesNotContainsRange() with args %v : PASSED, expected %v and got value '%v'",
		inputrange2, want, have)
	} else {
		t.Errorf("DoesNotContainsRange() with args %v : FAILED, expected %v but got value '%v'",
		inputrange2, want, have)
		t.Fatal()
	}
}

func TestDoesNotContainsRangeFalsePasses2(t *testing.T){
	inputrange := Range{"[-10,3)", -10, 2}
	inputrange2 := Range{"[-10,2]", -10, 2}
	want := false
	have := inputrange.DoesNotContainsRange(inputrange2)

	if want == have {
		t.Logf("DoesNotContainsRange() with args %v : PASSED, expected %v and got value '%v'",
		inputrange2, want, have)
	} else {
		t.Errorf("DoesNotContainsRange() with args %v : FAILED, expected %v but got value '%v'",
		inputrange2, want, have)
		t.Fatal()
	}
}

func TestDoesNotContainsRangeTruePasses1(t *testing.T){
	inputrange := Range{"(50,160)", 51, 159}
	inputrange2 := Range{"[150,200]", 150, 200}
	want := true
	have := inputrange.DoesNotContainsRange(inputrange2)

	if want == have {
		t.Logf("DoesNotContainsRange() with args %v : PASSED, expected %v and got value '%v'",
		inputrange2, want, have)
	} else {
		t.Errorf("DoesNotContainsRange() with args %v : FAILED, expected %v but got value '%v'",
		inputrange2, want, have)
		t.Fatal()
	}
}

func TestDoesNotContainsRangeTruePasses2(t *testing.T){
	inputrange := Range{"[-45,32)", -45, 31}
	inputrange2 := Range{"(-55,0]", -56, 0}
	want := true
	have := inputrange.DoesNotContainsRange(inputrange2)

	if want == have {
		t.Logf("DoesNotContainsRange() with args %v : PASSED, expected %v and got value '%v'",
		inputrange2, want, have)
	} else {
		t.Errorf("DoesNotContainsRange() with args %v : FAILED, expected %v but got value '%v'",
		inputrange2, want, have)
		t.Fatal()
	}
}

func TestEqualsTruePasses1(t *testing.T){
	inputrange := Range{"(2,9)", 3, 8}
	inputrange2 := Range{"[3,8]", 3, 8}
	want := true
	have := inputrange.Equals(inputrange2)

	if want == have{
		t.Logf("Equals() with args %v : PASSED, expected %v and got value '%v'", inputrange2,
		want, have)
	} else {
		t.Errorf("Equals() with args %v : FAILED, expected %v and got value '%v'", inputrange2,
		want, have)
		t.Fatal()
	}
}

func TestEqualsTruePasses2(t *testing.T){
	inputrange := Range{"[-13,2)", -13, 1}
	inputrange2 := Range{"[-13,1]", -13, 1}
	want := true
	have := inputrange.Equals(inputrange2)

	if want == have{
		t.Logf("Equals() with args %v : PASSED, expected %v and got value '%v'", inputrange2,
		want, have)
	} else {
		t.Errorf("Equals() with args %v : FAILED, expected %v and got value '%v'", inputrange2,
		want, have)
		t.Fatal()
	}
}

func TestEqualsFalsePasses1(t *testing.T){
	inputrange := Range{"(2,9]", 3, 9}
	inputrange2 := Range{"[1,11)", 1, 11}
	want := false
	have := inputrange.Equals(inputrange2)

	if want == have{
		t.Logf("Equals() with args %v : PASSED, expected %v and got value '%v'", inputrange2,
		want, have)
	} else {
		t.Errorf("Equals() with args %v : FAILED, expected %v and got value '%v'", inputrange2,
		want, have)
		t.Fatal()
	}
}

func TestEqualsFalsePasses2(t *testing.T){
	inputrange := Range{"[-10,7]", -10, 7}
	inputrange2 := Range{"[1,11)", 1, 10}
	want := false
	have := inputrange.Equals(inputrange2)

	if want == have{
		t.Logf("Equals() with args %v : PASSED, expected %v and got value '%v'", inputrange2,
		want, have)
	} else {
		t.Errorf("Equals() with args %v : FAILED, expected %v and got value '%v'", inputrange2,
		want, have)
		t.Fatal()
	}
}

func TestNotEqualsFalsePasses1(t *testing.T){
	inputrange := Range{"(2,9)", 3, 8}
	inputrange2 := Range{"[3,8]", 3, 8}
	want := false
	have := inputrange.NotEquals(inputrange2)

	if want == have{
		t.Logf("NotEquals() with args %v : PASSED, expected %v and got value '%v'", inputrange2,
		want, have)
	} else {
		t.Errorf("NotEquals() with args %v : FAILED, expected %v and got value '%v'", inputrange2,
		want, have)
		t.Fatal()
	}
}

func TestNotEqualsFalsePasses2(t *testing.T){
	inputrange := Range{"[-13,2)", -13, 1}
	inputrange2 := Range{"[-13,1]", -13, 1}
	want := false
	have := inputrange.NotEquals(inputrange2)

	if want == have{
		t.Logf("NotEquals() with args %v : PASSED, expected %v and got value '%v'", inputrange2,
		want, have)
	} else {
		t.Errorf("NotEquals() with args %v : FAILED, expected %v and got value '%v'", inputrange2,
		want, have)
		t.Fatal()
	}
}

func TestNotEqualsTruePasses1(t *testing.T){
	inputrange := Range{"(2,9]", 3, 9}
	inputrange2 := Range{"[1,11)", 1, 11}
	want := true
	have := inputrange.NotEquals(inputrange2)

	if want == have{
		t.Logf("NotEquals() with args %v : PASSED, expected %v and got value '%v'", inputrange2,
		want, have)
	} else {
		t.Errorf("NotEquals() with args %v : FAILED, expected %v and got value '%v'", inputrange2,
		want, have)
		t.Fatal()
	}
}

func TestNotEqualsTruePasses2(t *testing.T){
	inputrange := Range{"[-10,7]", -10, 7}
	inputrange2 := Range{"[1,11)", 1, 10}
	want := true
	have := inputrange.NotEquals(inputrange2)

	if want == have{
		t.Logf("NotEquals() with args %v : PASSED, expected %v and got value '%v'", inputrange2,
		want, have)
	} else {
		t.Errorf("NotEquals() with args %v : FAILED, expected %v and got value '%v'", inputrange2,
		want, have)
		t.Fatal()
	}
}

func TestOverlapsTruePasses1(t *testing.T){
	inputrange := Range{"(-15,23]", -14, 23}
	inputrange2 := Range{"(10,55]", 11, 55}
	want := true
	have := inputrange.Overlaps(inputrange2)

	if want == have {
		t.Logf("Overlaps() with args %v : PASSED, expected %v and got value '%v'",
		inputrange2, want, have)
	} else {
		t.Errorf("Overlaps() with args %v : FAILED, expected %v and got value '%v'",
		inputrange2, want, have)
		t.Fatal()
	}
}

func TestOverlapsTruePasses2(t *testing.T){
	inputrange := Range{"[20,66]", 20, 66}
	inputrange2 := Range{"(65,100]", 66, 100}
	want := true
	have := inputrange.Overlaps(inputrange2)

	if want == have {
		t.Logf("Overlaps() with args %v : PASSED, expected %v and got value '%v'",
		inputrange2, want, have)
	} else {
		t.Errorf("Overlaps() with args %v : FAILED, expected %v and got value '%v'",
		inputrange2, want, have)
		t.Fatal()
	}
}

func TestOverlapsFalsePasses1(t *testing.T){
	inputrange := Range{"(-11,22)", -10, 21}
	inputrange2 := Range{"(22,33]", 23, 33}
	want := false
	have := inputrange.Overlaps(inputrange2)

	if want == have {
		t.Logf("Overlaps() with args %v : PASSED, expected %v and got value '%v'",
		inputrange2, want, have)
	} else {
		t.Errorf("Overlaps() with args %v : FAILED, expected %v and got value '%v'",
		inputrange2, want, have)
		t.Fatal()
	}
}

func TestOverlapsFalsePasses2(t *testing.T){
	inputrange := Range{"[2,9)", 2, 8}
	inputrange2 := Range{"(9,13)", 10, 12}
	want := false
	have := inputrange.Overlaps(inputrange2)

	if want == have {
		t.Logf("Overlaps() with args %v : PASSED, expected %v and got value '%v'",
		inputrange2, want, have)
	} else {
		t.Errorf("Overlaps() with args %v : FAILED, expected %v and got value '%v'",
		inputrange2, want, have)
		t.Fatal()
	}
}

