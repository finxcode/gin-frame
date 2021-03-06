package models

import "strconv"

type User struct {
	ID
	Name     string `json:"name" gorm:"not null;comment:user name"`
	Mobile   string `json:"mobile" gorm:"not null;index;comment:user phone"`
	Password string `json:"password" gorm:"not null;default:'';comment:user password"`
	Timestamps
	SoftDeletes
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}
