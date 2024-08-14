package model

import (
	"time"
)

type AccountStatus int

const (
	Draft AccountStatus = iota
	Published
	Archived
)

type Tour struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`          // Primary key with auto-increment
	Name        string        `gorm:"column:name;type:varchar(255)" json:"name"`        // Column name and type
	Description string        `gorm:"column:description;type:text" json:"description"` // Column name and type
	// Status      AccountStatus `gorm:"column:status;type:int" json:"status"`      // Column name and type
	// Difficulty  int           `gorm:"column:difficulty;type:int" json:"difficulty"`  // Column name and type
	// Price       float64       `gorm:"column:price;type:float" json:"price"`       // Column name and type
	// Tags        []string      `gorm:"column:tags;type:text" json:"tags"`        // Column name and type
	// IsDeleted   bool          `gorm:"column:is_deleted;type:boolean" json:"is_deleted"`  // Column name and type
	// AuthorID    int       `gorm:"column:author_id;type:int" json:"author_id"`    // Column name and type
	// PublishTime time.Time `gorm:"column:publish_time;type:timestamp" json:"publish_time"` // Column name and type
	// Equipments  []int64   `gorm:"-" json:"equipments"`  // Ignore for now, adjust as needed
	// Checkpoints []int64   `gorm:"-" json:"checkpoints"` // Ignore for now, adjust as needed
	// Objects     []int64   `gorm:"-" json:"objects"`     // Ignore for now, adjust as needed
	// FootTime    float64   `gorm:"column:foot_time;type:float" json:"foot_time"`   // Column name and type
	// BicycleTime float64   `gorm:"column:bicycle_time;type:float" json:"bicycle_time"`// Column name and type
	// CarTime     float64   `gorm:"column:car_time;type:float" json:"car_time"`    // Column name and type
	// TotalLength float64   `gorm:"column:total_length;type:float" json:"total_length"`// Column name and type
	// Images      []string `gorm:"-" json:"images"`      // Ignore for now, adjust as needed

}
