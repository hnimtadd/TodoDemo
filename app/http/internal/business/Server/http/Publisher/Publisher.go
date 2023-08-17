package publisher

type PublishRequest struct {
	ContentType string `json:"Content-Type"`
	Action      string `json:"action"`
	Body        []byte `json:"body"`
}

type PublishResponse struct {
	Code int    `json:"status-code"`
	Body []byte `json:"body"`
}
type Publisher interface {
	Publish(PublishRequest) (*PublishResponse, error)
}
