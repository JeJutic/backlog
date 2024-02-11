package response

type Response struct {
	Error string `json:"error,omitempty"`
}

func Error(error string) Response {
	return Response{
		Error: error,
	}
}
