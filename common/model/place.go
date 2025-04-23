package model

import "github.com/uptrace/bun"

type Country struct {
	bun.BaseModel `bun:"table:countries"`

	ID    int64 `bun:",pk,autoincrement" json:"id"`
	ApiID int64 `json:"api_id"`

	Code string `json:"code"`
	Name string `json:"name"`

	States []*State `bun:"rel:has-many,join:country_id=id"`
}

type State struct {
	bun.BaseModel `bun:"table:states"`

	ID        int64 `bun:",pk,autoincrement" json:"id"`
	ApiID     int64 `json:"api_id"`
	CountryID int64 `json:"country_id"`

	Name string `json:"name"`
	Code string `json:"code"`

	Country *Country `bun:"rel:belongs-to,join:country_id=id"`
	Cities  []*City  `bun:"rel:has-many,join:state_id=id"`
}

type City struct {
	bun.BaseModel `bun:"table:cities"`

	ID      int64 `bun:",pk,autoincrement" json:"id"`
	ApiID   int64 `json:"api_id"`
	StateID int64 `json:"state_id"`

	Name string `json:"name"`
	Code string `json:"code"`

	State *State `bun:"rel:belongs-to,join:state_id=id"`
}
