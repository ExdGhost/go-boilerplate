package rating

// GetRequest request for getting a rating
type GetRequest struct {
	ID string
}

// GetResponse response after getting a rating
type GetResponse struct {
	ID    string `json:"id,omitempty"`
	Stars string `json:"stars,omitempty"`
}
