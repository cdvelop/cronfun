package cronfunc

func (c *cronFunc) AddFuncToSchedule(schedule, description string, fun any, args ...any) (err string) {
	const e = "AddFuncToSchedule "

	if descr, exist := c.registered[schedule]; exist {
		return e + "al agregar cronograma:" + description + ", ya existe cronograma: " + schedule + " = " + descr
	}

	er := c.AddJob(schedule, fun, args...)
	if er != nil {
		return e + er.Error()
	}

	c.registered[schedule] = description

	return ""
}
