package query

const (
	CreateBankAcc = `
		INSERT INTO gophkepper.bank (user_id, number, card_holder,  expiration_date, security_code, nonce_hex, metadata)
		VALUES($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, number, card_holder,  expiration_date, security_code, nonce_hex, metadata
	`

	CreateBankAccByUsername = `
		INSERT INTO gophkepper.bank (user_id, number, card_holder,  expiration_date, security_code, nonce_hex, metadata)
		(SELECT user_id, $2, $3, $4, $5, $6, $7
		FROM gophkepper.users
		WHERE gophkepper.users.username = $1)
		RETURNING id, number, card_holder,  expiration_date, security_code, nonce_hex, metadata
	`

	UpdateBankAcc = `
		UPDATE gophkepper.bank 
		SET number = $2, card_holder = $3, expiration_date = $4, security_code = $5, nonce_hex = $6, metadata = $7 
		WHERE id = $1
		RETURNING id, number, card_holder,  expiration_date, security_code, nonce_hex, metadata
	`

	GetBankAccountByID = `
		SELECT id, number, card_holder,  expiration_date, security_code, nonce_hex, metadata FROM gophkepper.bank 
		WHERE gophkepper.bank.id = $1
	`

	GetBankAccountsByUserID = `
		SELECT id, number, card_holder,  expiration_date, security_code,  nonce_hex, metadata FROM gophkepper.bank  
		WHERE gophkepper.bank.user_id = $1
	`

	GetBankAccountsByUsername = `
		SELECT id, number, card_holder,  expiration_date, security_code,  nonce_hex, metadata FROM gophkepper.bank b
		    INNER JOIN gophkepper.users u on u.id = b.user_id
		WHERE u.username = $1
	`

	GetBankAccounts = `
		SELECT id, number, card_holder,  expiration_date, security_code, nonce_hex, metadata FROM gophkepper.bank
	`

	DeleteBankAcc = ` 
        DELETE FROM gophkepper.bank
		WHERE id = $1 
    `
	DeleteBankAccByUsername = ` 
        DELETE
		FROM gophkepper.bank as b
		USING gophkepper.users as u
		WHERE  u.id  = b.user_id AND u.username = $1
    `

	DeleteBankAccByUserID = ` 
        DELETE
		FROM gophkepper.bank as b
		USING gophkepper.users as u
		WHERE  u.id  = $1 AND u.id  = b.user_id
    `
)
