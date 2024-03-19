package query

const (
	CreatePair = `
		INSERT INTO gophkepper.pairs (user_id, login, password_hash, metadata)
		VALUES($1, $2, $3, $4)
		RETURNING id, login, password_hash
	`

	CreatePairByUsername = `
		INSERT INTO gophkepper.pairs (user_id, login, password_hash, metadata)
		(SELECT user_id, $2, $3, $4
		FROM gophkepper.users
		WHERE gophkepper.users.username = $1)
		RETURNING id, login, password_hash
	`

	UpdatePair = `
		UPDATE gophkepper.pairs
		SET login  = $1, password_hash = $2, metadata = $3
		WHERE id = $4
		RETURNING id, login, password_hash
	`

	GetPairsByID = `
		SELECT id, login, password_hash FROM gophkepper.pairs
		WHERE pairs.id = $1
	`

	GetPairsByUserID = `
		SELECT p.id, login, p.password_hash, metadata FROM gophkepper.pairs p
		INNER JOIN gophkepper.users u on u.id = p.user_id
		WHERE p.user_id = $1
	`

	GetPairsByUsername = `
		SELECT p.id, p.login, p.password_hash, p.metadata FROM gophkepper.pairs p
		    INNER JOIN gophkepper.users u on u.id = p.user_id
		WHERE u.username = $1
	`

	GetPairs = `
		SELECT id, login, password_hash FROM gophkepper.pairs
	`

	DeletePair = ` 
        DELETE FROM gophkepper.pairs
		WHERE id = $1 
    `
	DeletePairByUsername = ` 
        DELETE
		FROM gophkepper.pairs as p
		USING gophkepper.users as u
		WHERE  u.id  = p.user_id AND u.username = $1
    `

	DeletePairByUserID = ` 
        DELETE
		FROM gophkepper.pairs as p
		USING gophkepper.users as u
		WHERE  u.id  = $1 AND u.id  = p.user_id
    `
)
