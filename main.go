package main

import (
	"os"

	"github.com/ispiroglu/placer/area"
	"github.com/ispiroglu/placer/config"
	"github.com/ispiroglu/placer/genetic"
)

func main() {
	cfg := config.ParseConfig(os.Args)
	area := area.ReadArea()

	genetic.Calculate(cfg, *area)
}
