package models

// import "gorm.io/gorm"

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (User) TableName() string {
	return "Users"
}

type User struct {
	// gorm.Model        // adds ID, created_at etc.
	Userid string `gorm:"column:Userid; type:varchar(50);primaryKey"`
	Name   string `gorm:"column:Name; type:varchar(50)"`
}
