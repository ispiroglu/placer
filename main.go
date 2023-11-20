package main

import (
	"os"

	"github.com/ispiroglu/placer/config"
	"github.com/ispiroglu/placer/file"
	"github.com/ispiroglu/placer/genetic"
)

func main() {
	cfg := config.ParseConfig(os.Args)
	env := file.ReadEnv()
	

	genetic.Calculate(cfg, env)
}
