package state

type UserStore interface {
	GetUsername() string
	SetUsername(string)
}

type userStore struct {
	username string
}

func NewUserStore(username string) UserStore {
	return userStore{username: username}
}

func (u userStore) GetUsername() string {
	panic("Implement Me")
}

func (u userStore) SetUsername(s string) {
	panic("Implement Me")
}



