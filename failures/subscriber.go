package failures

func SUBSCRIBER_DOES_NOT_EXSIS() SuperError {
	instance := create()
	instance.Text = "No such user. "
	instance.status = 404
	return instance
}
func GENERAL() SuperError {
	instance := create()
	instance.Text = "Unkown error"
	instance.status = 404
	return instance
}

func INVALID_TYPE() SuperError {

	instance := create()
	instance.status = 404
	instance.Text = "subscriber about to expire"
	return instance
}
