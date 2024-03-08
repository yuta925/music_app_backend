package port

type ULID interface {
	GenerateID() string
}
