package model

import (
	"time"
)

type Tour struct {
	Id          int       `json:"id"`          // Fixed JSON tag syntax
	Info        *TourInfo `json:"info"`        // Fixed JSON tag syntax
	AuthorId    int       `json:"author_id"`    // Fixed JSON tag syntax
	PublishTime time.Time `json:"publish_time"` // Fixed JSON tag syntax
	Equipments  []int64   `json:"equipments"`  // Fixed JSON tag syntax
	Checkpoints []int64   `json:"checkpoints"` // Fixed JSON tag syntax
	Objects     []int64   `json:"objects"`     // Fixed JSON tag syntax
	FootTime    float64   `json:"foot_time"`   // Fixed JSON tag syntax
	BicycleTime float64   `json:"bicycle_time"`// Fixed JSON tag syntax
	CarTime     float64   `json:"car_time"`    // Fixed JSON tag syntax
	TotalLength float64   `json:"total_length"`// Fixed JSON tag syntax
	Images      []string `json:"images"`      // Fixed JSON tag syntax
}
