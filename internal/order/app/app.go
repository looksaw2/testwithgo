package app

import "github.com/looksaw2/gorder3/internal/order/app/query"

// CORS读写分离
type Application struct {
	Command *Command
	Queries *Queries
}

type Command struct {
}

type Queries struct {
	GetCustomerOrder query.GetCustomerOrderHandler
}
