package model

type TempUser struct {
	Id       int64
	Username string
	Age      int
}

func (TempUser) TableName() string {
	return "temp_user"
}
