package github

import (
	"fmt"
	//"os"
	"time"
)

var count1 int = 1

func main() {
	//var ch1 chan int
	//var ch2 map[string] chan int
	//var ch3 chan [10]int
	ch4 := make([]chan int, 10)

	for i:=0; i<10 ;i++  {
		ch4[i] = make(chan int)
		go Count1(ch4[i], i)
	}
	for _,ch := range ch4{
		time.Sleep(1*time.Second)
		val := <- ch
		fmt.Println(count1, val)
	}

	//time.Sleep(1*time.Second)
	//os.Exit(1)

}

//Count
func Count1(ch chan int, i int) {
	fmt.Println("func:", i)
	ch <- i	//阻塞操作，被读取之前
	count1++
	fmt.Println("Count:", count1, i)

}
