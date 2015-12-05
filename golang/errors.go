package trip_service

import ()

type CollaboratorCallError struct {
	msg string
}

func (err CollaboratorCallError) Error() string {
	return err.msg
}

type UserNotLoggedInError struct {
	msg string
}

func (err UserNotLoggedInError) Error() string {
	return err.msg
}
