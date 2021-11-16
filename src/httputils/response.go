package httputils

type response struct {
	Data   interface{}   `json:"data"`
	Errors []interface{} `json:"errors"`
}

func NewResponse() response {
	return response{Data: map[string]string{}, Errors: []interface{}{}}
}
