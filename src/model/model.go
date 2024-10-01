package model

import (
	"time"

	"gorm.io/datatypes"
)

type Posts struct {
	Id        int `json:"id" gorm:"primary_key auto_increment"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Tags      datatypes.JSON `json:"tags"`
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"type:datetime"`
}
