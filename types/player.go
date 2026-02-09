package types

import "github.com/moznion/go-optional"

type Player struct {
	Status     string `json:"status"`
	Order      int    `json:"order"`
	PersonID   int    `json:"personId"`
	JerseyNum  string `json:"jerseyNum"`
	Starter    string `json:"starter"`
	OnCourt    string `json:"onCourt"`
	Played     string `json:"played"`
	Name       string `json:"name"`
	NameI      string `json:"nameI"`
	FirstName  string `json:"firstName"`
	FamilyName string `json:"familyName"`

	// if requested box score
	Statistics optional.Option[PlayerBoxScoreStatistic] `json:"statistics"`
}

func (p Player) IsStarter() bool {
	return p.Starter == "1"
}

func (p Player) IsOnCourt() bool {
	return p.OnCourt == "1"
}

func (p Player) IsPlayed() bool {
	return p.Played == "1"
}
