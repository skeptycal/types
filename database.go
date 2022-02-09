package types

type (

	// CRUD   SQL
	// Create INSERT
	// Read   SELECT
	// Update UPDATE
	// Delete DELETE
	Cruder interface {
		Create() // INSERT
		Read()   // SELECT
		Update() // UPDATE
		Delete() // DELETE
	}
)
