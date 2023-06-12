package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
)

func main() {
	//runtime.GOMAXPROCS(1)

	ctx := context.Background()
	wg, wgCtx := errgroup.WithContext(ctx)

	for i := 0; i < 3; i++ {
		i := i
		wg.Go(func() error {
			for j := 0; j < 10; j++ {

				if wgCtx.Err() != nil {
					return wgCtx.Err()
				}
				if i == 1 && j == 3 {
					return errors.New("error from go func")
				}
				fmt.Println("i:", i, " j:", j)
				//runtime.Gosched()
			}
			return nil
		})
	}
	err := wg.Wait()

	if err != nil {
		log.Println("error in goroutine.Error: ", err)
	}
}
