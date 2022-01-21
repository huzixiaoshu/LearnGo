// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var stockcode = 123
// 	var enddate = "2020-12-31"
// 	var url = "Code=%d&endData=%s"
// 	var target_url = fmt.Sprintf(url, stockcode, enddate)
// 	fmt.Println(target_url)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a string = "huwei"
// 	fmt.Println(a)

// 	var b, c int = 1, 2
// 	fmt.Println(b, c)

// 	var d int
// 	fmt.Println(d)

// 	var e bool
// 	fmt.Println(e)

// 	var f = true
// 	fmt.Println(f)

// 	g := "he huan"
// 	fmt.Println(g)

// }

// package main

// import (
// 	"fmt"
// 	"unsafe"
// )

// const (
// 	a = "abc"
// 	b = len(a)
// 	c = unsafe.Sizeof(a)
// )

// const (
// 	a1 = iota
// 	b1
// 	c1
// )

// func main() {
// 	_, numb, strs := numbers()
// 	fmt.Println("hello,world", numb, strs)
// 	fmt.Println(a, b, c)
// 	fmt.Println(a1, b1, c1)
// }

// func numbers() (int, int, string) {
// 	a, b, c := 1, 2, "hhh"
// 	return a, b, c
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	/* 定义局部变量 */
// 	var a int = 100
// 	var b int = 200

// 	fmt.Printf("交换前，a 的值 : %d\n", a)
// 	fmt.Printf("交换前，b 的值 : %d\n", b)

// 	/* 调用 swap() 函数
// 	* &a 指向 a 指针，a 变量的地址
// 	* &b 指向 b 指针，b 变量的地址
// 	 */
// 	swap(&a, &b)

// 	fmt.Printf("交换后，a 的值 : %d\n", a)
// 	fmt.Printf("交换后，b 的值 : %d\n", b)
// }

// func swap(x *int, y *int) {
// 	*x, *y = *y, *x
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	getSquare := func(x float64) float64 {
// 		return math.Sqrt(x)
// 	}

// 	fmt.Println(getSquare(9))

// }

// package main

// import "fmt"

// type cb func(int) int

// func main() {
// 	testCallBack(1, CallBack)
// 	testCallBack(2, func(x int) int {
// 		fmt.Printf("I'am callback,x:%d\n", x)
// 		return x
// 	})
// }

// func testCallBack(x int, f cb) {
// 	f(x)
// }

// func CallBack(x int) int {
// 	fmt.Printf("I'am callback: %d\n", x)
// 	return x
// }

// package main

// import "fmt"

// func main() {
// 	add_func := add(1, 2)
// 	fmt.Println(add_func())
// 	fmt.Println(add_func())
// 	fmt.Println(add_func())

// 	add2_func := add(1, 2)
// 	fmt.Println(add2_func())
// 	fmt.Println(add2_func())
// 	fmt.Println(add_func())
// }

// func add(x1, x2 int) func() (int, int) {
// 	i := 0
// 	return func() (int, int) {
// 		i++
// 		return i, x1 + x2
// 	}
// }
// type Circle struct {
// 	radius float64
// }

// func main() {
// 	var c1 Circle
// 	c1.radius = 10.00
// 	fmt.Println(c1.radius)
// 	fmt.Println("Circle area: ", c1.getArea())
// 	c1.changeRadius(20)
// 	fmt.Println(c1.radius)
// 	change(&c1, 30)
// 	fmt.Println(c1.radius)

// }

// func (c Circle) getArea() float64 {
// 	return 3.14 * c.radius * c.radius
// }

// func (c *Circle) changeRadius(radius float64) {
// 	c.radius = radius
// }

// func change(c *Circle, radius float64) {
// 	c.radius = radius
// }

// package main

// import "fmt"

// func main() {
// 	/* 数组长度为 5 */
// 	var balance = []int{1000, 2, 3, 17, 50}

// 	/* 数组作为参数传递给函数 */
// 	avg := getAverage(balance, 5)

// 	/* 输出返回的平均值 */
// 	fmt.Printf("平均值为: %f ", avg)
// }
// func getAverage(arr []int, size int) float32 {
// 	var i, sum int
// 	var avg float32

// 	for i = 0; i < size; i++ {
// 		sum += arr[i]
// 	}

// 	avg = float32(sum) / float32(size)

// 	return avg
// }
// package main

// import "fmt"

// type Books struct {
// 	title   string
// 	author  string
// 	book_id int
// }

// // func changeBook(book Books) string { //把book对象传进来，返回的值是string类型的，也就是将被修改的值返回出来。
// // 	book.title = "book1_change"
// // 	return book.title
// // }

// func main() {
// 	var book1 Books
// 	book1.title = "book1"
// 	book1.author = "zuozhe"
// 	book1.book_id = 1
// 	//	var res = changeBook(book1) //然后在外面拿到被修改的值
// 	book1.title = "res" // 再重新赋值
// 	fmt.Println(book1)
// }
// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Println(len(os.Args))
// 	for _, arg := range os.Args {
// 		fmt.Println(arg)
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// type Phone interface {
// 	call()
// }

// type NokiaPhone struct {
// }

// func (nokiaPhone NokiaPhone) call() {
// 	fmt.Println("I am Nokia, I can call you!")
// }

// type IPhone struct {
// }

// func (iPhone IPhone) call() {
// 	fmt.Println("I am iPhone, I can call you!")
// }

// func main() {
// 	var phone Phone

// 	phone = new(NokiaPhone)
// 	phone.call()

// 	phone = new(IPhone)
// 	phone.call()

// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// func main() {
// 	go say("hello")
// 	say("world")
// }

// package main

// import "fmt"

// func sum(s []int, b chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	b <- sum
// }
// func main() {
// 	s := []int{7, 2, 8, -9, 4, 0}

// 	c := make(chan int)

// 	go sum(s[:len(s)/2], c)
// 	go sum(s[len(s)/2:], c)

// 	x, y := 0, <-c

// 	fmt.Println(x, y, x+y)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	ch := make(chan int, 2)

// 	ch <- 1
// 	ch <- 2

// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)

// }

// package main

// import (
// 	"fmt"
// )

// func fibonacci(n int, c chan int) {
// 	x, y := 0, 1
// 	for i := 0; i < n; i++ {
// 		c <- x
// 		x, y = y, x+y
// 	}
// 	close(c)
// }

// func main() {
// 	c := make(chan int, 10)
// 	go fibonacci(cap(c), c)

// 	for i := range c {
// 		fmt.Println(i)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {

// 	c := make(chan int, 10)

// 	go fibonacci(cap(c), c)

// 	for v := range c {
// 		fmt.Println("out:", time.Now())
// 		fmt.Println(v)
// 	}
// }

// func fibonacci(n int, c chan int) {
// 	x, y := 0, 1

// 	for i := 0; i < n; i++ {
// 		c <- x
// 		fmt.Println("in:", time.Now())
// 		time.Sleep(100 * time.Millisecond)
// 		x, y = y, x+y
// 	}

// 	close(c)
// }

package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris"))
}
