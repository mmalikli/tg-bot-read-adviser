package telegram

type UpdatesResponce struct {
	OK bool `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID int `json:"id"`
	Message string `json:"message"`
}
