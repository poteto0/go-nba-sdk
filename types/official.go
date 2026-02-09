package types

type Official struct {
	PersonID   int    `json:"personId"`
	Name       string `json:"name"`
	NameI      string `json:"nameI"`
	FirstName  string `json:"firstName"`
	FamilyName string `json:"familyName"`
	JerseyNum  string `json:"jerseyNum"`
	Assignment string `json:"assignment"`
}
