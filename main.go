package main

import (
	"context"
	"fmt"

	"github.com/ferranbt/example-pkl/config"
)

func main() {
	cfg, err := config.LoadFromPath(context.Background(), "./other.pkl")
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg)
}
