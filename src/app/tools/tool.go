package tools

import (
	"gambituser/src/app/constants"
	"time"
)

func DateMySql() string {
	return time.Now().Format(constants.CreatedFormat)
}
