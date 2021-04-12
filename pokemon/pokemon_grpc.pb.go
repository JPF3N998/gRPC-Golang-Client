// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pokemon

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SearchPokedexClient is the client API for SearchPokedex service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchPokedexClient interface {
	GetPokemon(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type searchPokedexClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchPokedexClient(cc grpc.ClientConnInterface) SearchPokedexClient {
	return &searchPokedexClient{cc}
}

func (c *searchPokedexClient) GetPokemon(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/pokemon.SearchPokedex/GetPokemon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchPokedexServer is the server API for SearchPokedex service.
// All implementations must embed UnimplementedSearchPokedexServer
// for forward compatibility
type SearchPokedexServer interface {
	GetPokemon(context.Context, *SearchRequest) (*SearchResponse, error)
	mustEmbedUnimplementedSearchPokedexServer()
}

// UnimplementedSearchPokedexServer must be embedded to have forward compatible implementations.
type UnimplementedSearchPokedexServer struct {
}

func (UnimplementedSearchPokedexServer) GetPokemon(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPokemon not implemented")
}
func (UnimplementedSearchPokedexServer) mustEmbedUnimplementedSearchPokedexServer() {}

// UnsafeSearchPokedexServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchPokedexServer will
// result in compilation errors.
type UnsafeSearchPokedexServer interface {
	mustEmbedUnimplementedSearchPokedexServer()
}

func RegisterSearchPokedexServer(s grpc.ServiceRegistrar, srv SearchPokedexServer) {
	s.RegisterService(&SearchPokedex_ServiceDesc, srv)
}

func _SearchPokedex_GetPokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchPokedexServer).GetPokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.SearchPokedex/GetPokemon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchPokedexServer).GetPokemon(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchPokedex_ServiceDesc is the grpc.ServiceDesc for SearchPokedex service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchPokedex_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pokemon.SearchPokedex",
	HandlerType: (*SearchPokedexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPokemon",
			Handler:    _SearchPokedex_GetPokemon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pokemon/pokemon.proto",
}
