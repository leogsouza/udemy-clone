package common

type Error struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewError(key string, err error) Error {
	res := Error{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}
