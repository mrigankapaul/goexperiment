package models

//User ...
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID int32 = 1
)

//GetUsers ...
func GetUsers() []*User {
	return GetUsers()
}

//AddUser ...
func AddUser(u User) (User, error) {
	u.ID = int(nextID)
	nextID++
	users = append(users, &u)
	return u, nil
}
