package decorator

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
)


type queryLoggingDecorator[C any, R any] struct {
	logger *slog.Logger
	base QueryHandler[C,R]
}


func(q *queryLoggingDecorator[C, R])Handle(
	ctx context.Context,
	cmd C,
)(
	result R,
	err error,
){
	slog.Info("Query command",
		slog.Any("query",generateActionName(cmd)),
		slog.Any("query_body",fmt.Sprintf("%#v",cmd)),
	)
	defer func ()  {
		if err == nil {
			slog.Info("Query execute success")
		}else{
			slog.Error("Query execute Failed",
			slog.String("Error",err.Error()),
		)
		}
	}()
	return q.base.Handle(ctx,cmd)
}


func generateActionName(cmd any) string {
	return strings.Split(fmt.Sprintf("%T",cmd),".")[1]
}