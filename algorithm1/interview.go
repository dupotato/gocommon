package algorithm1

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func Alternateprint() {
	// var a = 'a'
	// //生成26个字符
	// for i := 1; i <= 26; i++ {
	// 	fmt.Printf("%c", a)
	// 	a++
	// }
	letter := 'A'
	num := 1
	chanLetter := make(chan bool, 1)
	chanNum := make(chan bool, 1)
	chanNum <- true
	fmt.Println("begin")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for num <= 26 {
			select {
			case <-chanNum:
				fmt.Print(num)
				num++
				chanLetter <- true
			case <-chanLetter:
				fmt.Printf("%c", letter)
				letter++
				chanNum <- true
			}
		}
		wg.Done()
	}()
	wg.Wait()
}

func reverseString(str string) string {
	l := []rune(str)
	fmt.Println(l)
	if len(l) == 1 {
		return str
	}
	n := len(l)

	for i := 0; i < len(l)/2; i++ {
		l[i], l[n-i-1] = l[n-i-1], l[i]
	}
	return string(l)
}

//jz offer Cqueue
type CQueue struct {
	q    [2][]int
	flag int //not use
}

func Constructor() CQueue {
	return CQueue{
		flag: 0,
	}
}

func (this *CQueue) AppendTail(value int) {
	this.q[this.flag] = append(this.q[this.flag], value)
}

func (this *CQueue) DeleteHead() int {
	l := len(this.q[this.flag])
	if l == 0 {
		return -1
	}
	v := this.q[this.flag][0]
	this.q[this.flag] = append(this.q[this.flag][:0], this.q[this.flag][1:]...)
	return v
}

//jz -10
//var gfib map[int]int

func fib(n int) int64 {
	gfib := make([]int64, n+1)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	gfib[0] = 0
	gfib[1] = 1
	for i := 2; i < n+1; i++ {
		gfib[i] = gfib[i-1] + gfib[i-2]
	}
	fmt.Println(gfib)
	return gfib[n]
}

func fib2(n int) int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		switch {
		case i < 2:
			arr[i] = i
		default:
			arr[i] = arr[i-1] + arr[i-2]
		}
		//fmt.Println(arr[i])
	}
	fmt.Println(arr)
	return arr[n-1]
}

//==========
//orderp
var wg = &sync.WaitGroup{}
var wg1 = &sync.WaitGroup{}

func orderp1(i int) {
	fmt.Println("first")
	wg.Done()
}
func orderp2(i int) {
	wg.Wait()
	fmt.Println("second")
	wg1.Done()
}
func orderp3(i int) {
	wg1.Wait()
	fmt.Println("third")
}

func ordermain() {
	p := 5
	wg.Add(1)
	wg1.Add(1)
	go orderp1(p)
	go orderp2(p)
	go orderp3(p)

	time.Sleep(5 * time.Second)
}

//scienceEat
const N = 5

var forks [N]chan int

var state [N]int

func philoWork(i int) {
	for {
		time.Sleep(time.Second * time.Duration(rand.Intn(2)))
		getForks(i)
		time.Sleep(time.Second * time.Duration(rand.Intn(2)))
	}
}

func getForks(i int) {
	left := (i - 1 + N) % N
	right := i % N
	select {
	case <-forks[left]:
		select {
		case <-forks[right]:
			state[i] = 1
			eattime := rand.Intn(10) + 10
			fmt.Printf("%d eattime %d\n", i, eattime)
			time.Sleep(time.Second * time.Duration(eattime))
			forks[left] <- 1
			forks[right] <- 1
			state[i] = 0
		default:
			forks[left] <- 1
			fmt.Printf("%d cannot get right \n ", i)
		}
	default:
		fmt.Printf("%d cannot get left \n", i)
	}
}

func philoWorkMain() {
	for i := 0; i < N; i++ {
		forks[i] = make(chan int, 1)
		forks[i] <- 1
		state[i] = 0
		go philoWork(i)

	}
	time.Sleep(100 * time.Second)
}

///--------------------
var casInt int32 = 10

func casuse() {
	v := casInt
	fmt.Printf("%d \n", v)
	if atomic.CompareAndSwapInt32(&casInt, v, 100) {
		fmt.Printf("cas %d \n", casInt)
	} else {
		fmt.Println("no ok")
	}
}

func casMain() {
	go func() {
		casInt = casInt + 1
		fmt.Printf("xx %d \n", casInt)
	}()

	time.Sleep(10 * time.Second)
	go casuse()
	time.Sleep(2 * time.Second)

}

///----

func orderp11() {
	cChan := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	fmt.Println("begin")
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			cChan <- i
		}
		close(cChan)
	}()

	go func() {
		defer wg.Done()
		for i := range cChan {
			fmt.Println(i)
		}
	}()
	wg.Wait()
}

// map wait

type WaitMap struct {
	m  map[string]string
	rw *sync.RWMutex
	ot int
}

func (w *WaitMap) Put(key, value string) error {
	w.rw.Lock()
	defer w.rw.Unlock()
	if w.m == nil {
		return errors.New("error map")
	}
	w.m[key] = value
	return nil
}

func (w *WaitMap) Get(key string) string {
	w.rw.RLock()
	defer w.rw.RUnlock()
	timeDead := time.NewTimer(time.Duration(w.ot) * time.Second)
	for {
		select {
		case <-timeDead.C:
			fmt.Println("outtime")
			return ""
		}
		if v, ok := w.m[key]; ok {
			return v
		} else {
			time.Sleep(100 * time.Millisecond)
		}
	}
}

/// ipban

type Ipmap struct {
	m  map[string]*IpInfo
	rw *sync.RWMutex
}

type IpInfo struct {
	lastvist  time.Time
	isvist    bool
	vistcount int
	vistband  int
}

func NewIpmap(ctx context.Context) *Ipmap {
	o := &Ipmap{
		m: make(map[string]*IpInfo),
	}
	go func() {
		timer := time.NewTimer(time.Minute * 1)
		for {
			select {
			case <-timer.C:
				o.rw.Lock()
				for k, iv := range o.m {
					if time.Now().Sub(iv.lastvist) >= time.Minute*1 {
						delete(o.m, k)
					}
				}
				o.rw.Unlock()
				timer.Reset(time.Minute * 1)
			case <-ctx.Done():
				return
			}
		}
	}()
	return o
}

func (p *Ipmap) visit(ip string) bool {
	p.rw.Lock()
	defer p.rw.Unlock()
	if v, ok := p.m[ip]; ok {
		if v.isvist == false {
			v.vistband++
			return false
		} else {
			v.isvist = true
			v.lastvist = time.Now()
			v.vistcount++
		}
	} else {
		v := &IpInfo{
			isvist:    true,
			vistcount: 1,
			lastvist:  time.Now(),
		}
		p.m[ip] = v
	}
	return true
}

// recover continue proc

// func continueProc(){

// }

// func proc(){
// 	panic("ok")
// }

// produce consume

var appnum = 10
var apple = make(chan int, appnum)
var closed = make(chan int, 1)
var timer = time.NewTicker(1 * time.Millisecond)

func produce(id int) {
	if len(apple) == appnum {
		time.Sleep(1 * time.Second)
	}
	apple <- id
	fmt.Printf("produce:%d\n", id)
}

func consume() {

	// for v := range apple {

	// 	fmt.Println("consue:", v)

	// }
	// closed <- 1

	for {
		select {
		case i, v := <-apple:

			if !v {
				closed <- 1
				return
			}
			fmt.Println(i, v)
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func pcMain() {
	//timer.Stop()
	go func() {
		for i := 0; i < 100; i++ {
			produce(i)
		}
		close(apple)
	}()
	// //timer.Stop()
	// go func() {
	// 	timer.Stop()
	// 	//timer.Reset(1 * time.Second)
	// }()
	go consume()
	<-closed
}
