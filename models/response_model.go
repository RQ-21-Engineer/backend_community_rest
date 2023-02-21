package models



type JoinResponse struct {

	HttpResponse  int           `json:"http_response"`
	Status        string        `json:"status"`
	Data          interface{}   `json:"data"`

}
