package failures

func SUBSCRIBER_DOES_NOT_EXSIS() SuperError {
	instance := create()
	instance.Text = "No such user. "
	instance.status = 404
	return instance
}

func SUBSCRIBER_DO_NOT_NOTIFY() SuperError {

	instance := create()
	instance.status = 404
	instance.Text = "No such user. "
	return instance
}
