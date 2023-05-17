package fastHttp

import (
	"github.com/valyala/fasthttp"
	"log"
)

func CreateGetRequest(url string) *fasthttp.Request {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	req.SetRequestURI(url)
	return req
}

func CreatePutRequest(url string, body []byte) *fasthttp.Request {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("PUT")
	req.SetBody(body)
	req.SetRequestURI(url)
	return req
}

func SendRequest(url string, req *fasthttp.Request) (*fasthttp.Response, error) {
	res := fasthttp.AcquireResponse()
	err := fasthttp.Do(req, res)
	if err != nil {
		log.Println("error while sending request. Error: ", err.Error())
	}
	return res, err
}
