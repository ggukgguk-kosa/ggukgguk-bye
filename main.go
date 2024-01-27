package main

import (
	"github.com/ggukgguk-kosa/ggukgguk-bye/common"
	"github.com/ggukgguk-kosa/ggukgguk-bye/step1"
)

func main() {
	config := common.InitConfig()

	if config.NextStep == 1 {
		step1.Start(config)
	}
}
