package fastmvc

type HttpError struct {
	error
	statusCode int
}