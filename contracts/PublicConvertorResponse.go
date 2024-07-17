package contracts

type PublicUtilConvertorResponse struct {
	ID       string `json:"id"`
	Status   string `json:"status"`
	From     string `json:"from"`
	To       string `json:"to"`
	Payload  string `json:"payload"`
	Response string `json:"response"`
}
