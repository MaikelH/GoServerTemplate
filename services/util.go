package services

func NotFoundError(err error) bool {
	if err != nil && err.Error() == "sql: no rows in result set" {
		return true
	}

	return false
}
