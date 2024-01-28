package step1

import (
	"github.com/ggukgguk-kosa/ggukgguk-bye/common"
)

func Start(config common.Config) {
	DBInit(config.DBUri)
	PdfInit(config.FilePath)
	ReadAndWrite(LoadAllMembers())
	DBClose()
}
