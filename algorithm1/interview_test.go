package algorithm1

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_Alternateprint(t *testing.T) {
	Alternateprint()
}

func Test_reverseString(t *testing.T) {
	str := "678910"
	newstr := reverseString(str)
	fmt.Println(newstr)
}

func Test_fib(t *testing.T) {
	fmt.Println(fib(45))
	fmt.Println(fib2(46))
}

func Test_ordermain(t *testing.T) {
	ordermain()
}

func Test_philoWorkMain(t *testing.T) {
	philoWorkMain()
}

func Test_rand(t *testing.T) {
	fmt.Println(rand.Intn(10))
}

func Test_casMain(t *testing.T) {
	casMain()
}

func Test_orderp11(t *testing.T) {
	orderp11()
}

func Test_pcMain(t *testing.T) {
	pcMain()
}

func Test_reverse(t *testing.T) {
	reverse()
}

func Test_a1(t *testing.T) {
	fmt.Println(a1(3, 3))
}

func Test_pc1(t *testing.T) {
	pc()
}

func Test_findlongcontinusstr(t *testing.T) {
	s := "abacbaa"
	fmt.Println(findlongcontinusstr(s))
}

func Test_find111(t *testing.T) {
	n := "abcdefabc"
	set := []rune{'a', 'e', 'f'}
	fmt.Println(find111(n, set))
}
