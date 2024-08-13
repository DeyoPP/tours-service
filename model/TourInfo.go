package model

type TourInfo struct {
	Name        string        `gorm:"column:name;type:varchar(255)" json:"name"`        // Column name and type
	Description string        `gorm:"column:description;type:text" json:"description"` // Column name and type
	Status      AccountStatus `gorm:"column:status;type:int" json:"status"`      // Column name and type
	Difficulty  int           `gorm:"column:difficulty;type:int" json:"difficulty"`  // Column name and type
	Price       float64       `gorm:"column:price;type:float" json:"price"`       // Column name and type
	Tags        []string      `gorm:"column:tags;type:text" json:"tags"`        // Column name and type
	IsDeleted   bool          `gorm:"column:is_deleted;type:boolean" json:"is_deleted"`  // Column name and type
}

type AccountStatus int

const (
	Draft AccountStatus = iota
	Published
	Archived
)