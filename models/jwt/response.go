package mjwt

// Go Api server
// @jeffotoni
// 2021-01-04

//Response struct
type Response struct {
	Token   string `json:"token"`
	Expires string `json:"expires"`
	Message string `json:"message"`
}
