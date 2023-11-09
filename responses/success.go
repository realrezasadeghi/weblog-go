package responses

type Success struct {
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
}
