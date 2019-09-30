package users

// UserRepository abstracts accesses to users-related data.
type UserRepository interface {
	ExistsByUsername(string) (bool, error)
	ExistsByEmail(string) (bool, error)
}
