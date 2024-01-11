package cronfunc

import (
	"github.com/mileusna/crontab"
)

// https://crontab.guru/
// mon:1,th:2, we:3,th:4,fr:5,sat:6,sun:0

var cron *cronFunc

// schedule (cronograma) ej:
// "0 7 15 1 *"  // A las 07:00 del día del mes 15 de enero
// "0 7 * * 1,4" // 2 veces a la semana lunes y jueves
// "*/1 * * * *" // cada 1 min
func AddFuncToSchedule(schedule, description string, fun any, args ...any) (err string) {
	const e = "AddFuncToSchedule "

	if cron == nil {
		cron = &cronFunc{
			Crontab:    crontab.New(),
			registered: map[string]string{},
		}
	}

	if descr, exist := cron.registered[schedule]; exist {
		return e + "ya existe cronograma: " + schedule + " = " + descr
	}

	er := cron.AddJob(schedule, fun, args...)
	if er != nil {
		return e + er.Error()
	}

	return ""
}
