package message

import "errors"

var (
	Success  = errors.New("request success")
	NotFound = errors.New("not found")
)

var code = map[error]int{
	Success:  200,
	NotFound: 404,
}

func Code(err error) int {
	if code, ok := code[err]; ok {
		return code
	}

	return 404
}

func String(err error) string {
	return err.Error()
}
