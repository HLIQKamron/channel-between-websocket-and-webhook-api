package handler

type handlerV1 struct{}

var myChannel = make(chan map[string]interface{})

type HandlerV1Config struct{}

func NewHandlerV1(h *HandlerV1Config) *handlerV1 {
	return &handlerV1{}
}
