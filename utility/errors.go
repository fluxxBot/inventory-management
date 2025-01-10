package utility

type NotFoundError struct {
	Message    string
	StatusCode int
}

func (n NotFoundError) Error() string {
	panic("Item Not Found")
}

type ForbiddenError struct {
	Message    string
	StatusCode int
}

func (f ForbiddenError) Error() string {
	panic("Not Allowed")
}
