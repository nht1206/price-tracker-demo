package model

type User struct {
	ID         uint64
	FullName   string
	Email      string
	Gender     bool
	FollowType uint
}

func (u User) TableName() string {
	return "t_user"
}
