package myhttp

type response struct {
	Data   interface{}   `json:"data"`
	Errors []interface{} `json:"errors"`
}

func New() response {
	return response{Data: map[string]string{}, Errors: []interface{}{}}
}
