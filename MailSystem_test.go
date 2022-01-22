package main

import (
	"fmt"
	"testing"
)

// stack tests
func TestStack(t *testing.T) {
	var st Stack
	if res := st.IsEmpty(); !res {
		t.Errorf("Stack Empty, Fail")
	}

	st.Push("Hi")

	if res := st.IsEmpty(); res {
		t.Errorf("Stack is not Empty, Fail")
	}

	if element, ok := st.Top(); !ok || element != "Hi" {
		t.Errorf("Top got:%v err: %v, Fail", element, ok)
	}

	if element, ok := st.Pop(); !ok || element != "Hi" {
		t.Errorf("Pop got:%v err: %v, Fail", element, ok)
	}

	if element, ok := st.Pop(); ok {
		t.Errorf("Pop should return error got: %v, Fail", element)
	}

	if res := st.IsEmpty(); !res {
		t.Errorf("Stack is not Empty, Fail")
	}

}

func TestManyToOne(t *testing.T) {
	var tests = []struct {
		in   string
		want string
	}{
		{OperationsStr, OpStr},
		{CitiesStr, CityStr},
		{BranchesStr, BranchStr},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.in)
		t.Run(testname, func(t *testing.T) {
			ans := manyToOne(tt.in)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

// empty string test
func TestHandleLineEmpty(t *testing.T) {

	countryArr = make([]City, 0)
	countryMap = make(map[string]int)

	var st Stack
	st.Push(CitiesStr)
	// empty
	if res := handleLine("", CitiesStr, &st); res == nil {
		t.Errorf("insert empty string as city number, Fail")
	}

}

// test: insert float as number should get err, then insert 1 and city A.
func TestParseCity(t *testing.T) {

	countryArr = make([]City, 0)
	countryMap = make(map[string]int)

	var st Stack
	st.Push(CitiesStr)

	if res := handleLine("-1", CitiesStr, &st); res == nil {
		t.Errorf("insert -1 as city number should not be negative, Fail")
	}

	if res := handleLine("2.5", CitiesStr, &st); res == nil {
		t.Errorf("insert 2.5 as city number should be int, Fail")
	}

	if res := handleLine("1", CitiesStr, &st); res != nil {
		t.Errorf("insert \"1\" as city number, Fail")
	}

	currCom, _ := st.Top()
	if res := handleLine("A", currCom, &st); res != nil {
		t.Errorf("insert \"A\" as city, Fail")
	}
}

// test: insert float as number should get err, then insert negative values of branch, then vaild values branch.
func TestParseBranch(t *testing.T) {

	countryArr = make([]City, 0)
	countryMap = make(map[string]int)

	var st Stack
	st.Push(CitiesStr)

	if res := handleLine("1", CitiesStr, &st); res != nil {
		t.Errorf("insert \"1\" as city number, Fail")
	}

	currCom, _ := st.Top()
	if res := handleLine("A", currCom, &st); res != nil {
		t.Errorf("insert \"A\" as city, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("1.5", currCom, &st); res == nil { // number of branch
		t.Errorf("insert \"1.5\" as branches number should get int, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("-1", currCom, &st); res == nil { // number of branch
		t.Errorf("insert \"-1\" as branches number, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("1", currCom, &st); res != nil { // number of branch
		t.Errorf("insert \"1\" as branches number, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("0 -1 0", currCom, &st); res == nil { // branch
		t.Errorf("insert \"0 -1 0\" as branch should get positive int, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("0 0 -1", currCom, &st); res == nil { // branch
		t.Errorf("insert \"0 0 -1\" as branch should get positive int, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("-1 0 0", currCom, &st); res == nil { // branch
		t.Errorf("insert \"-1 0 0\" as branch should get positive int, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("0 3 1", currCom, &st); res == nil { // branch
		t.Errorf("insert \"0 3 1\" as branch, min weight should be smaller than max weight, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("0 2 3", currCom, &st); res != nil { // branch
		t.Errorf("insert \"0 2 3\" as branch, Fail")
	}
}

func TestParsePackage(t *testing.T) {

	countryArr = make([]City, 0)
	countryMap = make(map[string]int)

	var st Stack
	st.Push(CitiesStr)

	if res := handleLine("1", CitiesStr, &st); res != nil {
		t.Errorf("insert \"1\" as city number, Fail")
	}

	currCom, _ := st.Top()
	if res := handleLine("A", currCom, &st); res != nil {
		t.Errorf("insert \"A\" as city, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("1", currCom, &st); res != nil { // number of branch
		t.Errorf("insert \"1\" as branches number, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("2 2 3.5", currCom, &st); res != nil { // branch
		t.Errorf("insert \"2 2 3.5\" as branch, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("a -1", currCom, &st); res == nil { // package
		t.Errorf("insert \"a -1\" as package weight is negative, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("a 0", currCom, &st); res == nil { // package
		t.Errorf("insert \"a 0\" as package should be below the min, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("a 4", currCom, &st); res == nil { // package
		t.Errorf("insert \"a 4\" as package should be above the max, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("a 2.5", currCom, &st); res != nil { // package
		t.Errorf("insert \"a 2.5\" as package, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("a 3", currCom, &st); res != nil { // package
		t.Errorf("insert \"a 3\" as package, Fail")
	}
}

/* test:
1. try to get last City and last Branch when everything is empty.
2. insert 2 cities, A with 1 branch which is empty, B with 2 branches first one empty and second one with 1 package
*/
func TestGetLast(t *testing.T) {

	countryArr = make([]City, 0)
	countryMap = make(map[string]int)

	br := getLastBranch()
	if br != nil {
		t.Errorf("should get nil as branch got: %v", br)
	}

	ct := getLastCity()
	if ct != nil {
		t.Errorf("should get nil as city got: %v", br)
	}

	var st Stack
	st.Push(CitiesStr)
	if res := handleLine("2", CitiesStr, &st); res != nil {
		t.Errorf("insert \"2\" as city number, Fail")
	}

	currCom, _ := st.Top()
	if res := handleLine("A", currCom, &st); res != nil {
		t.Errorf("insert \"A\" as city, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("1", currCom, &st); res != nil { // number of branch
		t.Errorf("insert \"1\" as branches number, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("0 2 3", currCom, &st); res != nil { // branch 1
		t.Errorf("insert \"0 1 3\" as branch 1, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("B", currCom, &st); res != nil {
		t.Errorf("insert \"A\" as city, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("2", currCom, &st); res != nil { // number of branch
		t.Errorf("insert \"2\" as branches number, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("0 1 3", currCom, &st); res != nil { // branch 2
		t.Errorf("insert \"0 1 3\" as branch 1, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("1 1 3", currCom, &st); res != nil { // branch 3
		t.Errorf("insert \"1 1 3\" as branch 2, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("a 2", currCom, &st); res != nil { // package
		t.Errorf("insert \"a 2\" as package, Fail")
	}

	currCom, _ = st.Top()
	if ct = getLastCity(); ct == nil || ct.Name != "B" {
		t.Errorf("should get city B, got: %v", br)
	}

	currCom, _ = st.Top()
	if br = getLastBranch(); br == nil || len(br.Packages) != 1 {
		t.Errorf("should get branch 3, got: %v", br)
	}

}

func TestOpOne(t *testing.T) {

	countryArr = make([]City, 0)
	countryMap = make(map[string]int)

	var st Stack
	st.Push(OperationsStr)
	st.Push(CitiesStr)
	if res := handleLine("2", CitiesStr, &st); res != nil {
		t.Errorf("insert \"2\" as city number, Fail")
	}

	currCom, _ := st.Top()
	if res := handleLine("A", currCom, &st); res != nil {
		t.Errorf("insert \"A\" as city, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("1", currCom, &st); res != nil { // number of branch
		t.Errorf("insert \"1\" as branches number, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("0 2 3", currCom, &st); res != nil { // branch 1
		t.Errorf("insert \"0 1 3\" as branch 1, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("B", currCom, &st); res != nil {
		t.Errorf("insert \"A\" as city, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("2", currCom, &st); res != nil { // number of branch
		t.Errorf("insert \"2\" as branches number, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("1 1 3", currCom, &st); res != nil { // branch 2
		t.Errorf("insert \"1 1 3\" as branch 1, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("a 2", currCom, &st); res != nil { // package
		t.Errorf("insert \"a 2\" as package, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("2 1 3", currCom, &st); res != nil { // branch 3
		t.Errorf("insert \"2 1 3\" as branch 2, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("b 2", currCom, &st); res != nil { // package
		t.Errorf("insert \"b 2\" as package, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("c 2", currCom, &st); res != nil { // package
		t.Errorf("insert \"c 2\" as package, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("-1", currCom, &st); res == nil { // operations number
		t.Errorf("insert \"-1\" as operations number should not be negative, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("1", currCom, &st); res != nil { // operations number
		t.Errorf("insert \"1\" as operations number should not be negative, Fail")
	}

	words := []string{"1", "G"}
	if _, ok := parseOp(words); ok == nil {
		t.Errorf("insert \"1 G\" as operation, G is not a vaild city, Fail")
	}

	wordsWrong := []string{"1"}
	if _, ok := parseOp(wordsWrong); ok == nil {
		t.Errorf("insert \"1\" as operation, too few arguments, Fail")
	}

	wordsWrong = []string{"0"}
	if _, ok := parseOp(wordsWrong); ok == nil {
		t.Errorf("insert \"0\" as operation, 0 is not a vaild op, Fail")
	}

	wordsWrong = []string{"4"}
	if _, ok := parseOp(wordsWrong); ok == nil {
		t.Errorf("insert \"4\" as operation, 4 is not a vaild op, Fail")
	}

	printAStr := "A:\n\t0:\n"
	printBStr := "B:\n\t0:\n\t\ta\n\t1:\n\t\tb\n\t\tc\n"

	words = []string{"1", "A"}
	if str, ok := parseOp(words); ok != nil || str != printAStr {
		t.Errorf("insert \"1 A\" as operation, got: \n%s  \nshould get: \n%s\n, Fail", str, printAStr)
	}

	words = []string{"1", "B"}
	if str, ok := parseOp(words); ok != nil || str != printBStr {
		t.Errorf("insert \"1 B\" as operation, got: \n%s  \nshould get: \n%s\n, Fail", str, printBStr)
	}
}

func TestOpTwo(t *testing.T) {

	countryArr = make([]City, 0)
	countryMap = make(map[string]int)

	var st Stack
	st.Push(OperationsStr)
	st.Push(CitiesStr)
	if res := handleLine("2", CitiesStr, &st); res != nil {
		t.Errorf("insert \"2\" as city number, Fail")
	}

	currCom, _ := st.Top()
	if res := handleLine("A", currCom, &st); res != nil {
		t.Errorf("insert \"A\" as city, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("1", currCom, &st); res != nil { // number of branch
		t.Errorf("insert \"1\" as branches number, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("0 2 3", currCom, &st); res != nil { // branch 1
		t.Errorf("insert \"0 1 3\" as branch 1, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("B", currCom, &st); res != nil {
		t.Errorf("insert \"A\" as city, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("2", currCom, &st); res != nil { // number of branch
		t.Errorf("insert \"2\" as branches number, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("1 1 3", currCom, &st); res != nil { // branch 2
		t.Errorf("insert \"1 1 3\" as branch 1, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("a 1", currCom, &st); res != nil { // package
		t.Errorf("insert \"a 1\" as package, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("2 1 3", currCom, &st); res != nil { // branch 3
		t.Errorf("insert \"2 1 3\" as branch 2, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("b 2", currCom, &st); res != nil { // package
		t.Errorf("insert \"b 2\" as package, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("c 2", currCom, &st); res != nil { // package
		t.Errorf("insert \"c 2\" as package, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("1", currCom, &st); res != nil { // operations number
		t.Errorf("insert \"1\" as operations number should not be negative, Fail")
	}

	words := []string{"2", "G", "0", "B", "0"}
	if _, ok := parseOp(words); ok == nil {
		t.Errorf("insert \"2 G 0 B 0\" as operation, G is not a vaild city, Fail")
	}

	words = []string{"2", "A", "1", "B", "0"}
	if _, ok := parseOp(words); ok == nil {
		t.Errorf("insert \"2 A 1 B 0\" as operation, branch 1 is not a vaild branch in A, Fail")
	}

	words = []string{"2", "A", "-1", "B", "0"}
	if _, ok := parseOp(words); ok == nil {
		t.Errorf("insert \"2 A -1 B 0\" as operation, branch -1 is not a vaild branch in A, Fail")
	}

	words = []string{"2", "B", "1", "A", "0"}
	if _, ok := parseOp(words); ok != nil {
		t.Errorf("insert \"2 B 1 A 0\" as operation, Fail")
	}

	printAStr := "A:\n\t0:\n\t\tb\n\t\tc\n"
	printBStr := "B:\n\t0:\n\t\ta\n\t1:\n"

	words = []string{"1", "A"}
	if str, ok := parseOp(words); ok != nil || str != printAStr {
		t.Errorf("insert \"1 A\" as operation, got: \n%s  \nshould get: \n%s\n, Fail", str, printAStr)
	}

	words = []string{"1", "B"}
	if str, ok := parseOp(words); ok != nil || str != printBStr {
		t.Errorf("insert \"1 B\" as operation, got: \n%s  \nshould get: \n%s\n, Fail", str, printBStr)
	}

	words = []string{"2", "B", "0", "A", "0"}
	if _, ok := parseOp(words); ok != nil {
		t.Errorf("insert \"2 B 1 A 0\" as operation, Fail")
	}

	words = []string{"1", "A"}
	if str, ok := parseOp(words); ok != nil || str != printAStr {
		t.Errorf("insert \"1 A\" as operation, got: \n%s  \nshould get: \n%s\n, Fail", str, printAStr)
	}

	words = []string{"1", "B"}
	if str, ok := parseOp(words); ok != nil || str != printBStr {
		t.Errorf("insert \"1 B\" as operation, got: \n%s  \nshould get: \n%s\n, Fail", str, printBStr)
	}
}

func TestOpThree(t *testing.T) {

	countryArr = make([]City, 0)
	countryMap = make(map[string]int)

	noTown := "No Town"

	words := []string{"3"}
	if str, ok := parseOp(words); ok != nil || str != noTown {
		t.Errorf("insert \"3\" as operation got: \n%s\n, should get:\n%s\n Fail", str, noTown)
	}

	var st Stack
	st.Push(OperationsStr)
	st.Push(CitiesStr)
	if res := handleLine("2", CitiesStr, &st); res != nil {
		t.Errorf("insert \"2\" as city number, Fail")
	}

	currCom, _ := st.Top()
	if res := handleLine("A", currCom, &st); res != nil {
		t.Errorf("insert \"A\" as city, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("1", currCom, &st); res != nil { // number of branch
		t.Errorf("insert \"1\" as branches number, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("0 2 3", currCom, &st); res != nil { // branch 1
		t.Errorf("insert \"0 1 3\" as branch 1, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("B", currCom, &st); res != nil {
		t.Errorf("insert \"A\" as city, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("2", currCom, &st); res != nil { // number of branch
		t.Errorf("insert \"2\" as branches number, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("1 1 3", currCom, &st); res != nil { // branch 2
		t.Errorf("insert \"1 1 3\" as branch 1, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("a 1", currCom, &st); res != nil { // package
		t.Errorf("insert \"a 1\" as package, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("2 1 3", currCom, &st); res != nil { // branch 3
		t.Errorf("insert \"2 1 3\" as branch 2, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("b 2", currCom, &st); res != nil { // package
		t.Errorf("insert \"b 2\" as package, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("c 2", currCom, &st); res != nil { // package
		t.Errorf("insert \"c 2\" as package, Fail")
	}

	currCom, _ = st.Top()
	if res := handleLine("1", currCom, &st); res != nil { // operations number
		t.Errorf("insert \"1\" as operations number should not be negative, Fail")
	}

	aStr := "Town with the most number of packages is A"
	bStr := "Town with the most number of packages is B"

	words = []string{"3"}
	if str, ok := parseOp(words); ok != nil || str != bStr {
		t.Errorf("insert \"3\" as operation got: \n%s\n, should get:\n%s\n Fail", str, bStr)
	}

	words = []string{"2", "B", "1", "A", "0"}
	if _, ok := parseOp(words); ok != nil {
		t.Errorf("insert \"2 B 1 A 0\" as operation, Fail")
	}

	words = []string{"3"}
	if str, ok := parseOp(words); ok != nil || str != aStr {
		t.Errorf("insert \"3\" as operation got: \n%s\n, should get:\n%s\n Fail", str, aStr)
	}
}
