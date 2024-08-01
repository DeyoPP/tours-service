package model

type AccountStatus int

const (
	Draft AccountStatus = iota
	Published
	Archived
)

type TourInfo struct {
	Name        string        `json: "name"`
	Description string        `json: "description"`
	Status      AccountStatus `json: "status"`
	Difficulty  int           `json: "difficulty"`
	Price       float64       `json: "price"`
	Tags        []string      `json: "tags"`
	IsDeleted   bool          `json: "is_deleted"`
}
