package main

import (
	"github.com/ispiroglu/placer/config"
	"github.com/ispiroglu/placer/genetic"
	"os"
)

func main() {
	cfg := config.ParseConfig(os.Args)

	genetic.Calculate(cfg)
}
