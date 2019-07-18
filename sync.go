package github

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type mylock struct {
	sync.Mutex
	sync.WaitGroup
	res []int
}

func main(){
	time_start := time.Now()

	var mylock1 mylock
	for i:=0; i<10; i++  {
		mylock1.Add(1)
		go mylock1.append(i)
	}
	mylock1.Wait()

	time_end := time.Now()
	fmt.Println(mylock1.res, time_start, time_end)
	os.Exit(0)

}

func (m *mylock) append(i int) {
	defer m.Done()
	defer m.Unlock()
	m.Lock()
	m.res = append(m.res, i)
	//time.Sleep(1*time.Second)
}





