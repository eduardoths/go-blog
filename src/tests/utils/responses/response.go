package responses

import (
	"encoding/json"
	"fmt"

	"github.com/eduardothsantos/go-blog/src/httputils"
)

func StrResponse(Data interface{}, Errors []interface{}) string {
	resp := httputils.NewResponse()
	if Data != nil {
		resp.Data = Data
	}
	if Errors != nil {
		resp.Errors = Errors
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		panic(fmt.Sprintf("Error parsing %v", resp))
	}
	return string(jsonResp)
}
