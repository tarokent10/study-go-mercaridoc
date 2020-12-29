package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	// 1つのサブタスクでエラーが発生したときに他の全てのサブタスクをキャンセルできる
	// withcontext使わない場合はGroup()を使う
	eg, ctx := errgroup.WithContext(context.Background())
	ps := []string{"tom", "jhon", "yam"}
	for _, p := range ps {
		eg.Go(func() error {
			return workerWithContext(ctx, p)
		})
		time.Sleep(2 * time.Second)
	}
	if err := eg.Wait(); err != nil {
		log.Fatal(err.Error())

	}
	fmt.Println("done!")
}
func workerWithContext(ctx context.Context, s string) error {
	time.Sleep(2 * time.Second)
	fmt.Printf("my name is %s\n", s)
	return nil
}
