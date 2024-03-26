package query

const (
	CreateUser = `
		INSERT INTO gophkepper.users (username, password_hash)
		VALUES($1, $2)
		RETURNING id, username, password_hash
	`

	UpdateUser = `
		UPDATE gophkepper.users
		SET username  = $1, password_hash = $2
		WHERE id = $3
		RETURNING id, username, password_hash
	`

	GetUserByID = `
		SELECT id, username, password_hash FROM gophkepper.users
		WHERE users.id = $1
	`

	GetUserByUsername = `
		SELECT id, username, password_hash FROM gophkepper.users
		WHERE users.username = $1
	`

	GetUsers = `
		SELECT id, username FROM gophkepper.users
	`

	DeleteUsers = ` 
        DELETE FROM gophkepper.users
		WHERE id = $1 
    `

	SearchUserTemp = `SELECT id, password_hash FROM gophkepper.users`
)
