package stock

import (
	"context"
	"fmt"
	"strings"

	"github.com/looksaw2/gorder3/internal/common/genproto/orderpb"
)



type Repository interface {
	GetItems(ctx context.Context , ids []string)([]*orderpb.Item,error)
}


type NotFoundError struct {
	MissingIDs []string
}


func (e NotFoundError)Error() string {
	return fmt.Sprintf("not found in stock %v",strings.Join(e.MissingIDs,","))
}