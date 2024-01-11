package cronfunc

// mon:1,th:2, we:3,th:4,fr:5,sat:6,sun:0

// schedule (cronograma) ej:
// "0 7 15 1 *"  // A las 07:00 del día del mes 15 de enero
// "0 7 * * 1,4" // 2 veces a la semana lunes y jueves
// "*/1 * * * *" // cada 1 min
func (c *cronFunc) AddFuncToSchedule(schedule, description string, fun any, args ...any) (err string) {
	const e = "AddFuncToSchedule "

	if descr, exist := c.registered[schedule]; exist {
		return e + "ya existe cronograma: " + schedule + " = " + descr
	}

	er := c.AddJob(schedule, fun, args...)
	if er != nil {
		return e + er.Error()
	}

	c.registered[schedule] = description

	return ""
}
