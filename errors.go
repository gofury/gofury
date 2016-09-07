package fastmvc

type HttpError struct {
	error
	statusCode int
}

func NewHttpError(err error, status int) *HttpError {
	return &HttpError{
		error: err,
		statusCode: status,
	}
}