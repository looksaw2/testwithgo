package server

import (
	"context"
	"log/slog"
	"net"
	"reflect"
	"strings"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func RunGRPCServer(servicename string, registerService func(service *grpc.Server)) {
	addr := viper.Sub(servicename).GetString("grpc-addr")
	if addr == "" {
		slog.Error("service Run grpc server",
			slog.String("service name", servicename),
			slog.String("on Port ", addr),
		)
		return
	}
	slog.Info("service Run grpc server",
		slog.String("service name", servicename),
		slog.String("on Port ", addr),
	)
	RunGRPCServerOnAddr(addr,registerService)
}

func RunGRPCServerOnAddr(addr string, registerService func(service *grpc.Server)) {
	grpcServer := grpc.NewServer(
		//单次的Interceptor
		grpc.ChainUnaryInterceptor(
			//提取对应的字段
			grpc.UnaryServerInterceptor(RequestLoggingInterceptor(slog.Default())),
		),
		//流式的Interceptor
		grpc.ChainStreamInterceptor(),
	)
	registerService(grpcServer)
	lis ,err := net.Listen("tcp",addr)
	if err != nil {
		slog.Error("GRPC Serverice Error",
			slog.String("service port",addr),
			slog.String("error ", err.Error()),
		)
		return 
	}
	if err = grpcServer.Serve(lis); err != nil {
		slog.Error("GRPC Serverice Error",
			slog.String("service port",addr),
			slog.String("error ", err.Error()),
		)
		return 
	}

}


//提取发送的grpc里面的所有的请求(Unary)
func RequestLoggingInterceptor(logger *slog.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// 1. 提取请求元数据（如header）
		mdFields := extractMetadataFields(ctx)

		// 2. 提取请求结构体的所有字段
		reqFields := extractRequestFields(req)

		// 3. 合并所有日志字段
		logFields := append(mdFields, reqFields...)
		logFields = append(logFields, "method", info.FullMethod)

		// 4. 记录请求日志
		logger.Info("gRPC request received", logFields...)

		// 5. 执行实际处理逻辑
		resp, err := handler(ctx, req)

		// 6. 记录响应日志（含错误）
		if err != nil {
			logger.Error("gRPC request failed", append(logFields, "error", err.Error())...)
		} else {
			// 可选：提取响应字段
			respFields := extractRequestFields(resp)
			logger.Info("gRPC request succeeded", append(logFields, respFields...)...)
		}

		return resp, err
	}
}

//提取grpc发送数据里面的元数据
func extractMetadataFields(ctx context.Context) []any {
	var fileds []any
	if md  ,ok := metadata.FromIncomingContext(ctx); ok {
		for k ,v := range md {
			fileds = append(fileds, "metadata_"+k,strings.Join(v,","))
		}
	}
	return fileds
}


//利用反射提取信息
func extractRequestFields(req any) []any {
	var fields []any
	if req == nil {
		return fields
	}
	val := reflect.ValueOf(req)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		fields = append(fields, req)
	}
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		//得到Tag
		fieldName := getFieldName(fieldType)
		if field.Kind() == reflect.Struct && ! (field.Type().PkgPath() == "time") { // 排除time.Time等基础类型
			nestedFields := extractRequestFields(field.Interface())
			// 为嵌套字段添加前缀
			for j := 0; j < len(nestedFields); j += 2 {
				nestedKey := nestedFields[j].(string)
				nestedVal := nestedFields[j+1]
				fields = append(fields, fieldName+"."+nestedKey, nestedVal)
			}
		} else if field.Kind() == reflect.Slice || field.Kind() == reflect.Array {
			// 处理切片/数组（简单转换为字符串）
			fields = append(fields, fieldName, field)
		} else {
			// 基础类型直接记录
			fields = append(fields, fieldName, field.Interface())
		}
	}
	return fields
}


//得到Tag
func getFieldName(fieldType reflect.StructField) string {
	protoTag := fieldType.Tag.Get("json")
	if protoTag != ""{
		return strings.Split(protoTag,",")[0]
	}
	return fieldType.Name
}
