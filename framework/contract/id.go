package contract

const IDKey = "fly:id"

type IDService interface {
	NewID() string
}
