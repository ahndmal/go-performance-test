package synch

import (
	"fmt"
	"os"
	"sync"
)

func AsyncHttp(times int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < times; i++ {
		fmt.Println(fmt.Sprintf("AsyncHttp Process is %d", os.Getpid()))
		fmt.Println(i)
	}
}
