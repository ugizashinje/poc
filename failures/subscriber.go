package failures

func SUBSCRIBER_DOES_NOT_EXSIS() SuperError {
	instance := create()
	instance.Text = "No such user. "
	return instance
}
