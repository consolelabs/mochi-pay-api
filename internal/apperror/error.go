package apperror

type AppErrorer interface {
	Error() string
	StatusCode() int
	Message() string
}
