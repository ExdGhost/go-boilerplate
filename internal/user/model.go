package user

// User contains all the properties of a user
type User struct {
	ID   string `json:"id,omitempty"   example:"110"`
	Name string `json:"name,omitempty" example:"Shourie Ganguly"`
	// Stars come from an HTTP API
	Stars     string    `json:"stars,omitempty" example:"five"`
	Favourite Favourite `json:"favourite,omitempty"`
}

// Favourite contains the list of users favourite stuff
type Favourite struct {
	// Beers come from a GRPC API
	Beers []string `json:"beers,omitempty"`
}
