package algorithm1

import (
	"fmt"
	"math/rand"
	"sync"
)

//3X4的方格 从左上角A走到右下角B 只能向右向下走 一共有多少种走法？
//思路 从最后的节点往前走
var a = 3
var b = 4

func a1(i, j int) int {
	if j == 1 {
		return 1
	}
	if i == 1 {
		return 1
	}
	m := a1(i-1, j)
	n := a1(i, j-1)
	return m + n
}

//sort

func quicksort(a []int, b, e int) {
	m := a[b]
	i := b
	j := e
	sw := true
	for i < j {
		if sw == true {
			if a[j] >= m {
				j--
				continue
			} else {
				a[i] = a[j]
				i++
			}
		} else {
			if a[i] > m {
				i++
				continue
			} else {
				a[j] = a[i]
				j--
			}
		}

	}
	a[i] = m
}

//producer - consumer

func pc() {
	wg := sync.WaitGroup{}
	out := make(chan int)
	wg.Add(2)
	go func() {
		for i := 0; i < 5; i++ {
			out <- rand.Intn(10)
		}
		close(out)
		wg.Done()
	}()

	go func() {

		// for v := range out {
		// 	fmt.Println(v)
		// }

		for {
			if i, ok := <-out; ok {
				fmt.Println(i)
			} else {
				break
			}

		}

		wg.Done()
	}()
	wg.Wait()
}

//给定一个字符串s ，请你找出其中不含有重复字符的最长连续子字符串
func findlongcontinusstr(str string) int {
	strl := []rune(str)
	n := len(strl)
	//start := 0
	left := 0
	right := 0

	// while(right < n){
	// 	//windows.add(s[right])
	// 	right++

	// 	while(while needs shrik){
	// 		windwos.remove(s[left])
	// 		left++
	// 	}
	// }

	needm := make(map[rune]int)
	max := 0
	for right < n {
		c := strl[right]
		right++
		if _, ok := needm[c]; !ok {
			needm[c] = 1
		} else {
			needm[c] = needm[c] + 1
		}

		for needm[c] > 1 {
			d := strl[left]
			left++
			needm[d]--
		}

		if max < (right - left) {
			max = right - left
		}

	}
	return max

}

//reverse
//给定m个不重复的字符[a,b,c,d]，以及一个长度为n的字符串tbcacbdata，问能否在这个字符串中找到一个长度为m的连续子串

func find111(n string, set []rune) bool {
	fn := []rune(n)
	l := len(fn)
	mapw := make(map[rune]int)
	for _, v := range set {
		mapw[v] = 0
	}
	left := 0
	right := 0
	vaild := len(set)

	for right < l {
		d := fn[right]
		right++
		if _, ok := mapw[d]; ok {
			mapw[d]++
		}
		fmt.Println(d)
		if mapw[d] == 1 {
			vaild--
		}
		fmt.Println(vaild)
		if vaild == 0 {
			return true
		}

		for mapw[d] > 1 {
			t := fn[left]
			left++
			mapw[t]--
			vaild++

		}

	}
	return false

}

//t :=abc string=adbcadef

// func findlongcontinussubstr(s, t string) {
// 	str := []rune(s)
// 	l := len(str)
// 	left := 0
// 	right := 0
// 	valid:=0
// 	need := make(map[rune]int)
// 	for i :=range t{
// 		need[i]++
// 	}
// 	win := make(map[rune]int)
// 	LL:=l+1
// 	for right < n {
// 		d := str[right]
// 		right++

// 		if _,ok:=need[d],ok{
// 			win[d]++
// 			if win[d]==need[d]{
// 				vaild++
// 			}

// 		}

// 		for valid==len(need){
// 			if right-left<LL{
// 				start=left
// 				LL=right-left
// 			}

// 			d:=s[left]
// 			left++
// 			if _,ok:=need[d],ok{
// 				if win[d]==need[d]{
// 					vaild--
// 				}else{
// 					win[d]--
// 				}
// 			}
// 		}

// 		return str[start,LL]
// 	}
// }

//levelorder

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

var result [][]int

// func levelorder(root *TreeNode) []int {
// 	dfs(root, 0)
// }

// func dfs(root *TreeNode, level int) {
// 	if root == nil {
// 		return
// 	}
// 	if len(result) == level {
// 		result = append(result, []int{})
// 	}
// 	result[level] = append(result[level], root.Value)
// 	dfs(root.Left, level+1)
// 	dfs(root.Right, level+1)
// }

// build tree
// preorder
// inorder

func buildTree(preorder []int, inorder []int) {
	build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func build(preorder []int, ps, pe int, inorder []int, is, ie int) *TreeNode {

	if ps < pe {
		return nil
	}
	h := preorder[ps]
	iis := is
	for iis < ie {
		if h == inorder[iis] {
			break
		}
		iis++
	}

	r := &TreeNode{
		Value: h,
	}
	r.Left = build(preorder, ps+1, ps+iis-is, inorder, is, iis-1)
	r.Right = build(preorder, ps+iis-is+1, iis-1, inorder, iis+1, ie)

	return r
}

//func lru

// type Lru struct {
// 	cap int
// 	S   []int
// 	M   map[int]int
// }

// func (l *Lru) Add(k int) {
// 	if _, ok := l.M[k]; ok {
// 		l.M[k] = time.Now()
// 	}
// }

type LruNode1 struct {
	Key  int
	Pre  *LruNode1
	Next *LruNode1
}

type Lru struct {
	cap int
	M   map[int]*LruNode1
	H   *LruNode1
	T   *LruNode1
}

func (l *Lru) add(k int) {
	if v, ok := l.M[k]; ok {
		//更改链表
		v.Pre.Next = v.Next
		v.Next.Pre = v.Pre
		v.Next = l.H
		v.Next.Pre = v
		return
	}

	// cap
	if len(l.M) < l.cap {
		t := &LruNode1{
			Key:  k,
			Next: nil,
			Pre:  l.T,
		}
		l.T.Next = t
		l.T = t
		l.M[k] = t
	}

	// 需要替换

}
