package types

type (
	Cruder interface {
		Create() // INSERT
		Read()   // SELECT
		Update() // UPDATE
		Delete() // DELETE
	}
)
