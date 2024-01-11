package cronfunc

import (
	"github.com/mileusna/crontab"
)

type cronFunc struct {
	*crontab.Crontab
	registered map[string]string
}
