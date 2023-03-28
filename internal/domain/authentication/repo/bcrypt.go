package repo

type BcryptRepo interface {
	CompareHashAndPassword(hashedPassword, password string) error
	GenerateFromPassword(password string) (string, error)
}
