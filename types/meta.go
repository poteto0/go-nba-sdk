package types

type Meta struct {
	Version int    `json:"version"`
	Request string `json:"request"`
	Time    string `json:"time"`
	Code    int    `json:"code"`
}
