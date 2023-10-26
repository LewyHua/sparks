package model

type Relation struct {
	// 用户的关注信息
	ID       int64 `gorm:"primaryKey"`
	UserId   int64 `gorm:"index;not null"` // 用户id
	FollowId int64 `gorm:"index;not null"` // 关注用户id
}

func (*Relation) TableName() string {
	return "relation"
}
