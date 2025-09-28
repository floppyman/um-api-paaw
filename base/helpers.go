package base

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
	
	"github.com/floppyman/um-common/logging/logr"
)

var Options PaawOptions
var KnownToken string
var KnownTokenExpires time.Time

type authRequestOptions struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
type authResponseOptions struct {
	Success bool                    `json:"success"`
	Data    authResponseOptionsData `json:"data"`
}
type authResponseOptionsData struct {
	Token     string `json:"token"`
	ExpiresIn int    `json:"expires_in"`
}

type HttpMethod string

const (
	HttpGet    HttpMethod = "GET"
	HttpPost   HttpMethod = "POST"
	HttpPatch  HttpMethod = "PATCH"
	HttpDelete HttpMethod = "DELETE"
)

func Init(options PaawOptions) {
	Options = options
	KnownToken = ""
	KnownTokenExpires = time.Now().Add(-(60 * 24 * time.Hour)) // 60 days back
}

func ValidateOrGetToken() (bool, error) {
	if KnownTokenExpires.After(time.Now()) {
		// token is good, continue
		return true, nil
	}
	
	bodyBytes, err := json.Marshal(authRequestOptions{
		ClientId:     Options.ApiClientId,
		ClientSecret: Options.ApiClientSecret,
	})
	if err != nil {
		return false, err
	}
	req := CreateRequest(HttpPost, "/auth", bodyBytes, false)
	if req == nil {
		return false, errors.New("create auth request failed")
	}
	
	suc, resBytes, err1 := DoRequest(req)
	if !suc || err1 != nil {
		return false, err1
	}
	
	var res authResponseOptions
	suc1, err2 := UnpackBody(resBytes, &res)
	if !suc1 || err2 != nil {
		return false, err2
	}
	
	KnownToken = res.Data.Token
	KnownTokenExpires = time.Now().Add(time.Duration(res.Data.ExpiresIn) * time.Second)
	return true, nil
}

func CreateRequest(method HttpMethod, urlPath string, body []byte, requiresToken bool) *http.Request {
	fullUrl := fmt.Sprintf("%s%s", Options.ApiUrl, urlPath)
	logr.Console.Debug().Msgf("PAAW Url: %s", fullUrl)
	
	var req *http.Request
	var err error
	
	if method != HttpGet && body != nil && len(body) > 0 {
		req, err = http.NewRequest(string(method), fullUrl, bytes.NewBuffer(body))
	} else {
		req, err = http.NewRequest(string(method), fullUrl, nil)
	}
	
	if err != nil {
		return nil
	}
	
	if requiresToken {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", KnownToken))
	}
	if method != HttpGet {
		req.Header.Add("Content-Type", "application/json")
	}
	
	return req
}

func DoRequest(req *http.Request) (bool, []byte, error) {
	client := &http.Client{}
	
	resp, err := client.Do(req)
	if err != nil {
		return false, nil, err
	}
	
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, nil, err
	}
	
	return true, body, nil
}

func UnpackBody(body []byte, res any) (bool, error) {
	err := json.Unmarshal(body, &res)
	if err != nil {
		logr.Console.Error().Str("body", string(body)).Msg("Invalid json")
		return false, err
	}
	
	return true, nil
}
