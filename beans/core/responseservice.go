package core

// ResponseService : class store response for business logics
type ResponseService struct {
	Data                 interface{}            `json:"data"`
	MapExtraResponseData map[string]interface{} `json:"mapExtraResponseData"`
	ResponseError        error                  `json:"responseError"`
}
