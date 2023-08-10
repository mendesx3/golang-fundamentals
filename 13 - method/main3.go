package main

type User struct {
	Username string
	Password string
}

func (u *User) Authenticate(user, password string) bool {
	return u.Username == user && u.Password == password
}

func main() {

	user := User{
		Username: "admin",
		Password: "admin",
	}

	isAuthenticated := user.Authenticate("admin", "admin")

	if isAuthenticated {
		println("User is authenticated")
	} else {
		println("User is not authenticated")
	}
}
