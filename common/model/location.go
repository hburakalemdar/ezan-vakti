package model

import "github.com/uptrace/bun"

type LocationType int

const (
	LocationTypeCountry LocationType = iota + 1
	LocationTypeState
	LocationTypeCity
)

type Location struct {
	bun.BaseModel `bun:"table:locations"`
	ID            int64  `bun:",pk,autoincrement" json:"id"`
	CreatedAt     string `json:"created_at"`

	ParentID *int64    `json:"parent_id,omitempty"`
	Parent   *Location `bun:"rel:belongs-to,join:parent_id=id"`

	ApiID int64        `bun:",unique:loc_api_type" json:"api_id"`
	Type  LocationType `bun:",unique:loc_api_type" json:"type"`
	Code  string       `json:"code"`
	Name  string       `json:"name"`

	Children []*Location `bun:"rel:has-many,join:id=parent_id"`
}
