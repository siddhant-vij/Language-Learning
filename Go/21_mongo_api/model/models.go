package model

type Netflix struct {
	Movie   string `json:"movie,omitempty" 		bson:"movie,omitempty"`
	Watched bool   `json:"watched,omitempty" 	bson:"watched,omitempty"`
}
