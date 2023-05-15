package myerror

// BadRequestError HTTP Status Code: 400
type BadRequestError struct {
	Err error
	Msg string
}

func (e *BadRequestError) Error() string {
	return "Bad Request Error"
}

// UnauthorizedError HTTP Status Code: 401
type UnauthorizedError struct {
	Err error
}

func (e *UnauthorizedError) Error() string {
	return "Unauthorized Error"
}

// NotFoundError HTTP Status Code: 404
type NotFoundError struct {
	Err error
}

func (e *NotFoundError) Error() string {
	return "Not Found Error"
}


// InternalServerError HTTP Status Code: 500
type InternalServerError struct {
	Err error
}

func (e *InternalServerError) Error() string {
	return "Internal Server Error"
}
