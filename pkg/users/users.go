package users

type Interface interface {
	CreateUserProfile(username string) error
	DeleteUserProfile(username string) error
}
