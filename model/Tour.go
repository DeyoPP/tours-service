package model

import (
	"time"
)

type Tour struct {
	Id          int       `json: "id"`
	Info        *TourInfo `json: "info"`
	AuthorId    int       `json: "author_id"`
	PublishTime time.Time `json: "publish_time"`
	Equipments  []int64   `json: "equipments"`
	Checkpoints []int64   `json: "checkpoints"`
	Objects     []int64   `json: "objects"`
	//TourReviews []*TourReview
	FootTime    float64  `json: "foot_time"`
	BicycleTime float64  `json: "bicycle_time"`
	CarTime     float64  `json: "car_time"`
	TotalLength float64  `json: "total_length"`
	Images      []string `json: "images"`
}
