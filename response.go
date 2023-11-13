package meituan

type CallbackResponse struct {
	Data string `json:"data"`
}

type Response interface {
	GetResultCode() int
	GetErrorList() []ErrorData
	GetError() Error
	GetErrorString() string
}

type BaseResponse struct {
	ResultCode *int `json:"result_code,omitempty"`
}

func (r BaseResponse) GetResultCode() int {
	return *r.ResultCode
}

func (r BaseResponse) GetErrorList() []ErrorData {
	return []ErrorData{}
}

func (r BaseResponse) GetError() Error {
	return Error{}
}

func (r BaseResponse) GetErrorString() string {
	return r.GetError().Msg
}

type ErrorResponse struct {
	ResultCode *int         `json:"result_code"`
	ErrorList  *[]ErrorData `json:"error_list,omitempty"`
	Error      *Error
	Msg        *string
	Data       string `json:"data"`
}

func (r ErrorResponse) GetResultCode() int {
	if r.ResultCode != nil {
		return *r.ResultCode
	} else if r.Error != nil {
		return r.Error.Code
	} else if r.ErrorList != nil {
		return r.GetError().Code
	}

	return 0
}

func (r ErrorResponse) GetErrorList() []ErrorData {
	return *r.ErrorList
}

func (r ErrorResponse) GetError() Error {
	if r.Error != nil && r.Error.Code > 1 {
		return *r.Error
	} else {
		t := *r.ErrorList
		return Error{
			Msg:  t[0].Msg,
			Code: t[0].Code,
		}
	}
}

func (r ErrorResponse) GetErrorString() string {
	if r.Error != nil && r.Error.Code > 1 {
		return r.Error.Msg
	} else {
		var msg string
		for k, v := range *r.ErrorList {
			if k == 0 {
				msg = v.Msg
			} else {
				msg = msg + "\n" + v.Msg
			}
		}
		return msg
	}
}

type ErrorData struct {
	Msg        string `json:"msg"`
	BlockFlag  int    `json:"blockflag"`
	Code       int    `json:"code"`
	AppSpuCode string `json:"app_spu_code"`
}

type Error struct {
	Msg  string
	Code int
}

type StringResponse struct {
	BaseResponse
	Data string
}

type SliceStringResponse struct {
	BaseResponse
	Data []string
}

type Logistics struct {
	ExpressCompany string `json:"express_company"`
	ExpressNumber  string `json:"express_number"`
}
