package repo

// User contains the properties stored in the repo
type User struct {
	ID   string `json:"Id,omitempty"`
	Name string `json:"name,omitempty"`
}
