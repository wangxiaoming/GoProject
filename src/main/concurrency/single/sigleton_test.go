package single

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Sigleton struct {
}

var sigletonInstance *Sigleton

var once sync.Once

func GetSingletonObj() *Sigleton {
	once.Do(func() {
		fmt.Println("Create obj")
		sigletonInstance = new(Sigleton)
	})
	return sigletonInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Println("%S", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
