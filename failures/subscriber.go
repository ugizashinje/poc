package failures

func SUBSCRIBER_DOES_NOT_EXSIS() SuperError {
	instance := create()
	instance.Text = "No such user. "
	instance.status = 404
	return instance
}

func SUBSCRIPTION_ABOUT_TO_EXPIRE() SuperError {

	instance := create()
	instance.status = 404
	instance.Text = "subscriber about to expire"
	return instance
}
