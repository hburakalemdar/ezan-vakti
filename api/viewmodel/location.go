package viewmodel

import (
	"github.com/hayrat/ezan-vakti/common/model"
)

type LocationResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	ParentID int64  `json:"parent_id,omitempty"`
	Type     int    `json:"type,omitempty"`
}

func (r *LocationResponse) ToDBModel() any {
	var parentID *int64
	if r.ParentID > 0 {
		parentID = &r.ParentID
	}

	return &model.Location{
		ApiID:    r.ID,
		Name:     r.Name,
		Code:     r.Code,
		Type:     model.LocationType(r.Type),
		ParentID: parentID,
	}
}
