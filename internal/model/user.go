package model

type User struct {
	ID         uint64
	FullName   string
	Email      string
	Gender     bool
	FollowType uint
	Lang       string
}

func (u User) TableName() string {
	return "t_user"
}
