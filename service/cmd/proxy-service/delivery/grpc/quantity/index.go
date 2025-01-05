package quantitygrpc

import "app/generated/proto/servicegrpc"

type quantityGrpc struct {
	servicegrpc.UnimplementedQuantityServiceServer
}

func Register() servicegrpc.QuantityServiceServer {
	return &quantityGrpc{}
}
