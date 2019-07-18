package github

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var count int = 0

func main() {
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