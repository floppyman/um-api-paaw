package roots

import "github.com/floppyman/um-api-paaw/base"

func (re RootEndPoint) Hi() (bool, error, string) {
	req := base.CreateRequest(base.HttpGet, "/", nil, false)
	
	ok1, body, err1 := base.DoRequest(req)
	if !ok1 {
		return false, err1, ""
	}
	
	res := string(body)
	return true, nil, res
}
