package contract

type BaseRequest struct {
}

type BaseResponse struct {
	StatusCode int32       `json:"status_code"`
	StatusInfo *StatusInfo `json:"status_info"`
}

type StatusInfo struct {
	Time    int64  `json:"time"`
	Message string `json:"message"`
}
