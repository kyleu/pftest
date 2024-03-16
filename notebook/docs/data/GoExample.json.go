package main

import (
	"encoding/json"
	"os"
	"time"
)

func main() {
	ret := [][]any{{"Go", time.Now()}}
	if err := json.NewEncoder(os.Stdout).Encode(ret); err != nil {
		panic(err)
	}
}
