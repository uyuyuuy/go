package main

import (
	"fmt"
	"os"
	"time"
)

var count1 int = 1

func main() {


	ChanSlice := make([]chan int, 10)

	for i := 9; i >= 0; i-- {
	//for i := 0; i < 10; i++ {
		ChanSlice[i] = make(chan int)
		go func(i int) {
			//fmt.Println(ChanSlice)
			//fmt.Println(i)
			ChanSlice[i] <- i
		}(i)
	}

	for _, c := range ChanSlice{
		j := <- c
		fmt.Println(j)
	}

	os.Exit(0)
	//var ch1 chan int
	//var ch2 map[string] chan int
	//var ch3 chan [10]int

	/*
	ch1 := make(chan int)

	go func() {
		fmt.Println("go func")
		time.Sleep(1*time.Second)
		fmt.Println("sleep 1")
		time.Sleep(1*time.Second)
		fmt.Println("sleep 2")
		time.Sleep(1*time.Second)
		fmt.Println("sleep 3")

		ch1 <- 10
	}()

	fmt.Println("out go func")
	v := <- ch1
	fmt.Println(v)

	fmt.Println("out channel")

	//os.Exit(0)
	*/



	/*
	//带缓冲的通道
	ch4 := make([]chan int, 100)

	for i:=0; i<100 ;i++  {
		ch4[i] = make(chan int, 2)
		go Count1(ch4[i], i)
	}
	time.Sleep(5*time.Second)

	for _,ch := range ch4{
		//time.Sleep(1*time.Second)
		va1 := <- ch
		va2 := <- ch
		va3 := <- ch
		fmt.Println(count1, va1, va2, va3)
	}
	 */


	//struct
	/*
	type Address struct {
		City	string
		Country	string
	}

	type Person struct {
		Name	string
		Age int
		Address Address
	}

	chan_struct := make(chan Person, 1)	//这里要用缓冲通道，否则一直在这里阻塞了

	Person_1 := Person{Name:"Xd", Age:30, Address:Address{"shenzhen","china"}}
	chan_struct <- Person_1

	Person_1.Address = Address{"guangzhou","china"}

	Person_chan := <- chan_struct

	//进入通道的数据是独立的，无法修改
	fmt.Println(Person_1)
	fmt.Println(Person_chan)

	os.Exit(0)

	 */

	ChannelTest := make(chan int, 1)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("开始写入",i)
			ChannelTest <- i
			fmt.Println("结束写入",i)

		}
	}()

	go func() {
		for true {
			time.Sleep(2*time.Second)
			v, ok := <-ChannelTest
			if !ok {
				break
			} else {
				fmt.Println("读取",v,ok)
			}

		}
	}()

	time.Sleep(10*time.Second)

	os.Exit(0)



	//关闭channel
	ChannelClose := make(chan int,	5)
	ChannelStep := make(chan int, 2)	//用于阻塞2个go，等待2个go都执行完了后才结束进程
	go func() {
		var i int
		for i = 0; i < 5; i++ {
			ChannelClose <- i
			time.Sleep(time.Second)
		}
		close(ChannelClose)
		ChannelStep <- 1
	}()

	time.Sleep(time.Second)

	go func() {

		for true {
			v,ok := <- ChannelClose
			fmt.Println(v,ok)
			if !ok {
				break
			}
		}
		ChannelStep <- 2

	}()

	<- ChannelStep
	<- ChannelStep


	os.Exit(1)


/*
	//select
	select {
		case v := <- ch4[0]:



	}
*/
	//time.Sleep(1*time.Second)
	//os.Exit(1)

}

//Count
func Count1(ch chan int, i int) {
	fmt.Println("func:", i)
	ch <- i	//阻塞操作，被读取之前,如果channel为缓冲类型，在没有装满的情况下不阻塞
	ch <- 2000
	fmt.Println(i, "=======")
	ch <- 3000	//超过缓冲数量就会阻塞，等待写入，该处会阻塞，上面不会阻塞
	//count1++
	fmt.Println(2, i)


	//fmt.Println("Count:", count1, i)

}
