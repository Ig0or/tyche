package responses

type BaseApiResponse struct {
	Payload    interface{} `json:"payload,omitempty"`
	Message    string      `json:"message,omitempty"`
	Success    bool        `json:"success"`
	StatusCode int         `json:"status_code"`
}
