package apisvc

type NotifyReq struct {
	UserUUID string `json:"user_uuid"`
	Message  string `json:"message"`
}
