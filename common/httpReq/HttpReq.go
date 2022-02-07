package httpReq

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// AuthResp 授权码返回结构
type AuthResp struct {
	Token string `json:"token"`
}

// GetAuthHeader 获取auth header
func GetAuthHeader() (header map[string]string, err error) {
	// 1、获取授权码
	auth, err := getAuthorization()
	if err != nil {
		return
	}
	// 2、生成header
	header = map[string]string{
		"Authorization": auth,
	}
	return
}

// getAuthorization 获取授权码
func getAuthorization() (auth string, err error) {
	// 1、构建需要的参数
	data := map[string]interface{}{
		"username": "username",
		"password": "password",
	}

	// 2、请求获取授权码
	resp, err := SendRequest("https://ip:port/auth", data, nil, "POST")
	if err != nil {
		return
	}

	// 3、从结果中获取token
	authResp := AuthResp{}
	err = json.Unmarshal(resp, &authResp)
	if err != nil {
		return
	}
	auth = authResp.Token
	return
}

// sendRequest 发送request
func SendRequest(url string, data interface{}, addHeaders map[string]string, method string) (resp []byte, err error) {
	jsonData, err := json.Marshal(data)
	body := bytes.NewReader(jsonData)
	// 1、创建req
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")

	// 2、设置headers
	if len(addHeaders) > 0 {
		for k, v := range addHeaders {
			req.Header.Add(k, v)
		}
	}

	// 3、发送http请求
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		err = errors.New("http status err")
		return
	}

	// 4、结果读取
	resp, err = ioutil.ReadAll(response.Body)
	return
}

// POST 发送request
func Post(url string, data map[string]interface{}, addHeaders map[string]string) (resp string, err error) {
	res, err := SendRequest(url, data, addHeaders, "POST")
	resp = string(res)
	return
}

// GET 发送request
func Get(url string, addHeaders map[string]string) (resp string, err error) {
	data := map[string]interface{}{}
	res, err := SendRequest(url, data, addHeaders, "GET")
	resp = string(res)
	return
}
