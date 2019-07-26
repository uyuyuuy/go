package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

var count int = 0
type answer []string

func main() {

	//var mutex sync.Mutex{}
	//var answers answer

	testQuestions := []string{}
	testQuestions = append(testQuestions, "1.")
	testQuestions = append(testQuestions, "2.")
	testQuestions = append(testQuestions, "3.")
	testQuestions = append(testQuestions, "4.")
	testQuestions = append(testQuestions, "5.")
	testQuestions = append(testQuestions, "6.")
	testQuestions = append(testQuestions, "7.")
	testQuestions = append(testQuestions, "8.")
	testQuestions = append(testQuestions, "9.")
	testQuestions = append(testQuestions, "10.")

	//fmt.Print(testQuestions)
	var answers answer
	//p_answer := *answer
	var mutex sync.Mutex
	var wait_group sync.WaitGroup
	for k,question := range testQuestions {
		wait_group.Add(1)	//调用协程之前执行
		//time.Sleep(1*time.Second)
		//fmt.Print(question, "\n")
		//wait_group.Add(1)
		go do_question1(k, question, &answers, &mutex, &wait_group)
	}
	wait_group.Wait()

	//time.Sleep(5*time.Second)
	fmt.Print("done:", answers, "\n")
	os.Exit(0)



	lock := &sync.Mutex{}
	for i := 0; i < 10; i++  {
		go Count(lock, i)
	}
	j := 0
	for  {
		lock.Lock()
		j++
		fmt.Println("ii=", j, ";count=", count)
		lock.Unlock()
		runtime.Gosched() //出让CPU时间片，降低处理优先级
		if count >= 10{
			break
		}

	}





}

func Count(lock *sync.Mutex, i int) {
	lock.Lock()
	time.Sleep(1*time.Second)
	//time.Sleep(1)
	count++
	fmt.Println("i=", i, ";count=", count)
	lock.Unlock()
}

func do_question1(k int, question string, mu *answer, mutex2 *sync.Mutex,wait_group *sync.WaitGroup) {
	defer wait_group.Done()
	mutex2.Lock()
	//time.Sleep(1*time.Second)
	*mu = append(*mu, question)
	mutex2.Unlock()
	//fmt.Print(*mu, "\n")
}

/*
在协程中单独使用Mutex锁，程序的运行没有阻塞，会立马执行下面的程序，不会管协程中的运行
在协程外面再加一个锁也没有用，最先加的锁最先释放，所以程序也会立马执行下面的程序，不会管协程中的运行

go do_question(k, question, &answers, &mutex, &wait_group)
& 符号很重要，类似引用，获取变量的指针，方法里面的操作都是操作同一个变量，否则每次传进来的都是新的变量

WaitGroup用于等待一组线程的结束。父线程调用Add方法来设定应等待的线程的数量。每个被等待的线程在结束时应调用Done方法。同时，主线程里可以调用Wait方法阻塞至所有线程结束

*/