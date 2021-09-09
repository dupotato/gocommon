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
