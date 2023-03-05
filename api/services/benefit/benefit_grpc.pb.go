package msbenefit

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion7

const (
	BenefitService_GetBenefit_FullMethodName = "/proto.BenefitService/GetBenefit"
)

type BenefitServiceClient interface {
	GetBenefit(ctx context.Context, in *NewBenefit, opts ...grpc.CallOption) (*User, error)
}

type benefitServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBenefitServiceClient(cc grpc.ClientConnInterface) BenefitServiceClient {
	return &benefitServiceClient{cc}
}

func (c *benefitServiceClient) GetBenefit(ctx context.Context, in *NewBenefit, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, BenefitService_GetBenefit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type BenefitServiceServer interface {
	GetBenefit(context.Context, *NewBenefit) (*User, error)
	mustEmbedUnimplementedBenefitServiceServer()
}

type UnimplementedBenefitServiceServer struct {
}

func (UnimplementedBenefitServiceServer) GetBenefit(context.Context, *NewBenefit) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBenefit not implemented")
}
func (UnimplementedBenefitServiceServer) mustEmbedUnimplementedBenefitServiceServer() {}

type UnsafeBenefitServiceServer interface {
	mustEmbedUnimplementedBenefitServiceServer()
}

func RegisterBenefitServiceServer(s grpc.ServiceRegistrar, srv BenefitServiceServer) {
	s.RegisterService(&BenefitService_ServiceDesc, srv)
}

func _BenefitService_GetBenefit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewBenefit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BenefitServiceServer).GetBenefit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BenefitService_GetBenefit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BenefitServiceServer).GetBenefit(ctx, req.(*NewBenefit))
	}
	return interceptor(ctx, in, info, handler)
}

var BenefitService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.BenefitService",
	HandlerType: (*BenefitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBenefit",
			Handler:    _BenefitService_GetBenefit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/benefit.proto",
}
