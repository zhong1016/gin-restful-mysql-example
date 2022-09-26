package restful

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(data any) *Result {
	R := Result{}

	R.Code = 200
	R.Message = "Success"
	R.Data = data

	return &R
}
func Fail(err string) *Result {
	R := Result{}

	R.Code = 200
	R.Message = err

	return &R
}
