package model

type User struct {
	userState map[int64]string
}

func NewUser() *User {
	return &User{
		userState: make(map[int64]string),
	}
}

func (u *User) SetUserState(ID int64, state string) {
	if u.userState == nil {
		u.userState = make(map[int64]string)
	}

	u.userState[ID] = state
}

func (u *User) GetUserState(ID int64) (string, bool) {
	if u.userState == nil {
		return "", false
	}

	return u.userState[ID], true
}

func (u *User) ClearUserState(ID int64) {
	if u.userState != nil {
		//delete
	}
}
