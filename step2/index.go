package step2

import (
	"github.com/ggukgguk-kosa/ggukgguk-bye/common"
)

func Start(config common.Config) {
	MailInit()
	sendAllPdf()
}
