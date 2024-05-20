package responses

import (
	"strconv"
)

type (
	Basic struct {
		Code int    `json:"code,string"`
		Msg  string `json:"msg,omitempty"`
	}
)

func (b *Basic) Error() string {
	return strconv.Itoa(b.Code) + " " + b.Msg
}
