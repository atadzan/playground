package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type requestBody struct {
	Client    string `json:"client"`
	Transport string `json:"transport"`
	Protocol  string `json:"protocol"`
	Encoding  string `json:"encoding"`
	User      string `json:"user"`
	Channel   string `json:"channel"`
	Data      data   `json:"data"`
}

type data struct {
	Process string `json:"process"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type SuccessResponse struct {
	Result struct{} `json:"result"`
}

type errorResponse struct {
	Error err `json:"error"`
}

type err struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

func serverErr(rw http.ResponseWriter, errMsg string) {
	errResponse, errJson := json.Marshal(errorResponse{
		Error: err{
			Code:    1002,
			Message: errMsg,
		},
	})
	if errJson != nil {
		log.Println("can't marshal error response.Error:", errJson.Error())
	}
	rw.Write(errResponse)
	return
}

func permissionDenied(rw http.ResponseWriter) {
	errResponse, errJson := json.Marshal(errorResponse{
		Error: err{
			Code:    1001,
			Message: "permission denied",
		},
	})
	if errJson != nil {
		log.Println("can't marshal error response.Error:", errJson.Error())
	}
	rw.Write(errResponse)
	return
}

func subscribeProxy(rw http.ResponseWriter, r *http.Request) {
	var input requestBody
	body, errBody := io.ReadAll(r.Body)
	if errBody != nil {
		log.Println("can't read request body.Error:", errBody.Error())
		serverErr(rw, errBody.Error())
		return
	}
	if errJson := json.Unmarshal(body, &input); errJson != nil {
		log.Println("can't unmarshal.Error:", errJson.Error())
		serverErr(rw, errJson.Error())
		return
	}

	// TODO Implement validation logic here

	response, errJson := json.Marshal(&SuccessResponse{})
	if errJson != nil {
		log.Println("can't marshal response.Error:", errJson.Error())
		serverErr(rw, errJson.Error())
		return
	}
	_, errResp := rw.Write(response)
	if errResp != nil {
		log.Println("can't write response.Error:", errResp.Error())
		serverErr(rw, errJson.Error())
		return
	}
	return
}

func publishProxy(rw http.ResponseWriter, r *http.Request) {
	var input requestBody
	body, errBody := io.ReadAll(r.Body)
	if errBody != nil {
		log.Println("can't read request body.Error:", errBody.Error())
		serverErr(rw, errBody.Error())
		return
	}
	if errJson := json.Unmarshal(body, &input); errJson != nil {
		log.Println("can't unmarshal.Error:", errJson.Error())
		serverErr(rw, errJson.Error())
		return
	}

	// TODO Implement validation logic here

	if input.User == "user-3" {
		log.Println(input.User, "tried to publish message but attempt was declined. Message:", input.Data)
		permissionDenied(rw)
		return
	}

	response, errJson := json.Marshal(&SuccessResponse{})
	if errJson != nil {
		log.Println("can't marshal response.Error:", errJson.Error())
		serverErr(rw, errJson.Error())
		return
	}
	_, errResp := rw.Write(response)
	if errResp != nil {
		log.Println("can't write response.Error:", errResp.Error())
		serverErr(rw, errJson.Error())
		return
	}
	return
}

func main() {
	http.HandleFunc("/centrifugo/subscribe", subscribeProxy)
	http.HandleFunc("/centrifugo/publish", publishProxy)

	if errRun := http.ListenAndServe(":8001", nil); errRun != nil {
		log.Println("can't start http server.Error:", errRun.Error())
		return
	}
}
