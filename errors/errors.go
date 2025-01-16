package errors

type NotFoundError struct {
	Message    string
	StatusCode int
}

func (n NotFoundError) Error() string {
	panic("Item Not Found" + n.Message)
}

type ForbiddenError struct {
	Message    string
	StatusCode int
}

func (f ForbiddenError) Error() string {
	panic("Not Allowed" + f.Message)
}
