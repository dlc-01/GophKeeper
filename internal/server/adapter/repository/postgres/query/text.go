package query

const (
	CreateTextInf = `
		INSERT INTO gophkepper.text (user_id, note, metadata)
		VALUES($1, $2, $3)
		RETURNING id, note, metadata
	`

	CreateTextInfByUsername = `
		INSERT INTO gophkepper.text (user_id, note, metadata)
		(SELECT user_id, $2, $3, $4
		FROM gophkepper.users
		WHERE gophkepper.users.username = $1)
		RETURNING id, note, metadata
	`

	UpdateTextInf = `
		UPDATE gophkepper.text
		SET note = $2, metadata = $3
		WHERE id = $1
		RETURNING id, note, metadata
	`

	GetBTextInfByID = `
		SELECT id, note, metadata FROM gophkepper.text 
		WHERE gophkepper.text.id = $1
	`

	GetTextInfsByUserID = `
		SELECT t.id, t.note, t.metadata FROM gophkepper.text t
		    INNER JOIN gophkepper.users u on u.id = t.user_id
		WHERE u.id = $1
	`

	GetTextInfsByUsername = `
		SELECT t.id, t.note, t.metadata FROM gophkepper.text t
		    INNER JOIN gophkepper.users u on u.id = t.user_id
		WHERE u.username = $1
	`

	GetTextInfs = `
		SELECT id, user_id, note, metadata FROM gophkepper.text
	`

	DeleteTextInf = ` 
        DELETE FROM gophkepper.text
		WHERE id = $1 
    `
	DeleteTextInfByUsername = ` 
        DELETE
		FROM gophkepper.text as t
		USING gophkepper.users as u
		WHERE  u.id  = t.user_id AND u.username = $1
    `
	DeleteTextInfByUserID = ` 
        DELETE
		FROM gophkepper.text as t
		USING gophkepper.users as u
		WHERE u.id = $1 AND  u.id  = t.user_id
    `
)
