package responses

//Общий ответ от API, независимо от результата.
type Response struct {
	Result string      `json:"result"`
	Meta   GeneralMeta `json:"meta"`
	Data   interface{} `json:"data"`
}

func NewResponse(result string, meta GeneralMeta, data interface{}) *Response {
	return &Response{
		Result: result,
		Meta:   meta,
		Data:   data,
	}
}

//ПОЛЕ META
//Общий интерфейс
type GeneralMeta interface {
	GetMessage() string
}

//Meta в случае успеха
//Если ошибки нет - возвращаем пустое поле Meta, либо с нужным описанием.
type MetaSuccess struct {
	Message string `json:"message"`
}

func NewMetaSuccess(message string) MetaSuccess {
	return MetaSuccess{
		Message: message,
	}
}

func (ms MetaSuccess) GetMessage() string {
	return ms.Message
}

//Meta в случае ошибки
//В случае ошибки - поле Data возвращаем пустым, а в поле Meta отправляем эту структуру.
type MetaError struct {
	Code    int16  `json:"code"`
	Message string `json:"message"`
}

func NewMetaError(code int16, message string) MetaError {
	return MetaError{
		Code:    code,
		Message: message,
	}
}

func (me MetaError) GetMessage() string {
	return me.Message
}
