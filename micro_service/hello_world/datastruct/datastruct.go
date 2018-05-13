package datastruct

// HelloWorldRequest get request with name string
type HelloWorldRequest struct {
	NAME string `json:"name"`
}

// HelloWorldResponse respond with code and message
type HelloWorldResponse struct {
	STATUS  int    `json:"code"`
	MESSAGE string `json:"message"`
}
