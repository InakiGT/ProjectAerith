package domain

type PasswordHasher interface {
	Hash(password string) (string, error)
	CompareHashAndPassword(hash, password string) error
}
