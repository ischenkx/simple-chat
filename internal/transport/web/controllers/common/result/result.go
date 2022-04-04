package result

import (
	"encoding/json"
	"io"
)

type Result struct {
	Error *Error      `json:"error"`
	Data  interface{} `json:"data"`
}

func Err(code int, message string) Result {
	return New(nil, NewError(code, message))
}

func Ok(data interface{}) Result {
	return New(data, nil)
}

func New(data interface{}, err *Error) Result {
	return Result{
		Error: err,
		Data:  data,
	}
}

func WriteSilent(w io.Writer, res Result) {
	// ignoring...
	Write(w, res)
}

func Write(w io.Writer, res Result) error {
	return json.NewEncoder(w).Encode(res)
}
