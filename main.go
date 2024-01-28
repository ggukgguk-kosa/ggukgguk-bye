package main

import (
	"log"

	"github.com/ggukgguk-kosa/ggukgguk-bye/common"
	"github.com/ggukgguk-kosa/ggukgguk-bye/step1"
	"github.com/ggukgguk-kosa/ggukgguk-bye/step2"
)

func main() {
	config := common.InitConfig()

	if config.NextStep == 1 {
		step1.Start(config)
	} else if config.NextStep == 2 {
		step2.Start(config)
	} else {
		log.Fatal("NextStep config is not valid")
	}
}
