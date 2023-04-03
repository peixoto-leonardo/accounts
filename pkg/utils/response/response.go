package response

type Response struct {
	StatusCode int
	Data       any
}

func New(code int, data any) Response {
	return Response{code, data}
}
