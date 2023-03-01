package account

import "time"

type Entity struct {
	ID        int
	Name      string
	CPF       string
	Secret    string
	Balance   int
	CreatedAt time.Time
}
