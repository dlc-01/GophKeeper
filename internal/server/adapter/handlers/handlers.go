package handlers

type ctxKey string

const (
	UserIDKey ctxKey = "user_id"
)

type Handlers struct {
	Auth  *AuthServer
	User  *UserServer
	Pairs *PairServer
	Cards *BankServer
	Notes *TextServer
	Token string
}
