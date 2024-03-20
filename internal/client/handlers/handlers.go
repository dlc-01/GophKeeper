package handlers

import "google.golang.org/grpc"

type Handlers struct {
	Auth      *AuthClient
	User      *UserClient
	Pairs     *PairClient
	Cards     *BankClient
	Notes     *TextClient
	Token     string
	SecretKey string
}

func New(conn *grpc.ClientConn) *Handlers {
	return &Handlers{
		Auth:  NewAuthClient(conn),
		User:  NewUserClient(conn),
		Pairs: NewPairClient(conn),
		Cards: NewBankClient(conn),
		Notes: NewTextClient(conn),
	}
}
