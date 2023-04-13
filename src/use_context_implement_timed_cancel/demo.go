package main

import (
	"context"
	"fmt"
	"time"
)

func FetchData(ctx context.Context) (res []int64) {
	select {
	case <-time.After(3 * time.Second):
		return []int64{100, 200, 300}
	case <-ctx.Done():
		return []int64{1, 2, 3}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ids := FetchData(ctx)
	fmt.Println(ids)
}
