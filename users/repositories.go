package users

type mockUserRepository struct{}

func (r *mockUserRepository) ExistsByUsername(username string) (bool, error) {
	if username == "existed" {
		return true, nil
	}
	return false, nil
}

func (r *mockUserRepository) ExistsByEmail(email string) (bool, error) {
	if email == "existed@gmail.com" {
		return true, nil
	}
	return false, nil
}

// NewMockUserRepository returns a new mock user repository.
func NewMockUserRepository() UserRepository {
	return new(mockUserRepository)
}
