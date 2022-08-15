package server

type httpError struct {
	Status  int    `json:"status"`
	Message string `json:"error"`
}

func newHttpError(status int, message string) httpError {
	return httpError{
		Status:  status,
		Message: message,
	}
}

func (err httpError) Error() string {
	return err.Message
}
