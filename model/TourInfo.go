package model

type TourInfo struct {
	Name        string        `json:"name"`        // Fixed JSON tag syntax
	Description string        `json:"description"` // Fixed JSON tag syntax
	Status      AccountStatus `json:"status"`      // Fixed JSON tag syntax
	Difficulty  int           `json:"difficulty"`  // Fixed JSON tag syntax
	Price       float64       `json:"price"`       // Fixed JSON tag syntax
	Tags        []string      `json:"tags"`        // Fixed JSON tag syntax
	IsDeleted   bool          `json:"is_deleted"`  // Fixed JSON tag syntax
}

type AccountStatus int

const (
	Draft AccountStatus = iota
	Published
	Archived
)