package events

type UserCreated struct {
	ID   int
	Name string
}

func (u UserCreated) Topic() string {
	return "user.created"
}
