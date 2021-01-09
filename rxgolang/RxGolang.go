package rxgolang

import (
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

func GoRX()  {
	observable := rxgo.Just("Hello, World!")()
	ch := observable.Observe()
	item := <-ch
	fmt.Println(item.V)
}