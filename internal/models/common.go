package models

type CommonResp struct {
	Message string      `json:"message"`
	Status  int64       `json:"status"`
	Result  interface{} `json:"result,omitempty"`
}
