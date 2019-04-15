//What will be printed when the code below is executed?
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	timeout := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout) //上层需要指定超时的情况:
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("waited for 1 sec")
	case <-time.After(2 * time.Second):
		fmt.Println("waited for 2 sec")
	case <-time.After(3 * time.Second):
		fmt.Println("waited for 3 sec")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}