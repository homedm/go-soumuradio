package main

import (
	"context"
	"fmt"
	"os"

	"github.com/tomato3713/soumu-radio-api/soumu"
)

func main() {
	cli, err := soumu.NewClient("")
	if err != nil {
		os.Exit(1)
	}

	opts := soumu.NewNumOptions(soumu.License, soumu.Amateur)
	result, err := cli.GetNum(context.Background(), opts)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(result)
}
