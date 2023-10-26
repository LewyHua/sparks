package models

type Relation struct {
	ID       int64 `gorm:"primarykey"`
	userID   int64
	followID int64
}

func (r *Relation) TableName() string {
	return "relation"
}
