// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.3
// source: spacemesh/v1/smesher.proto

package v1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_spacemesh_v1_smesher_proto protoreflect.FileDescriptor

var file_spacemesh_v1_smesher_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x6d, 0x65, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe1, 0x08, 0x0a, 0x0e, 0x53, 0x6d,
	0x65, 0x73, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0a,
	0x49, 0x73, 0x53, 0x6d, 0x65, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x20, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x73, 0x53, 0x6d, 0x65, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x6d, 0x65,
	0x73, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x23, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x53, 0x6d, 0x65, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x55, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x6d, 0x65, 0x73, 0x68, 0x69,
	0x6e, 0x67, 0x12, 0x21, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x6d, 0x65, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x6d, 0x65, 0x73, 0x68, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x53, 0x6d, 0x65,
	0x73, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6d,
	0x65, 0x73, 0x68, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x42, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x62, 0x61,
	0x73, 0x65, 0x12, 0x20, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x4d, 0x69, 0x6e, 0x47, 0x61,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x69, 0x6e, 0x47, 0x61, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x4d, 0x69,
	0x6e, 0x47, 0x61, 0x73, 0x12, 0x1e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x69, 0x6e, 0x47, 0x61, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x69, 0x6e, 0x47, 0x61, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0a, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a,
	0x14, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2a, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x23, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x1b, 0x53, 0x74, 0x6f, 0x70, 0x50,
	0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x50, 0x6f, 0x73, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a, 0x1e, 0x50,
	0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x34, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x34, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_spacemesh_v1_smesher_proto_goTypes = []interface{}{
	(*empty.Empty)(nil),                            // 0: google.protobuf.Empty
	(*StopSmeshingRequest)(nil),                    // 1: spacemesh.v1.StopSmeshingRequest
	(*SetCoinbaseRequest)(nil),                     // 2: spacemesh.v1.SetCoinbaseRequest
	(*SetMinGasRequest)(nil),                       // 3: spacemesh.v1.SetMinGasRequest
	(*CreatePostDataRequest)(nil),                  // 4: spacemesh.v1.CreatePostDataRequest
	(*StopPostDataCreationSessionRequest)(nil),     // 5: spacemesh.v1.StopPostDataCreationSessionRequest
	(*IsSmeshingResponse)(nil),                     // 6: spacemesh.v1.IsSmeshingResponse
	(*StartSmeshingResponse)(nil),                  // 7: spacemesh.v1.StartSmeshingResponse
	(*StopSmeshingResponse)(nil),                   // 8: spacemesh.v1.StopSmeshingResponse
	(*SmesherIdResponse)(nil),                      // 9: spacemesh.v1.SmesherIdResponse
	(*CoinbaseResponse)(nil),                       // 10: spacemesh.v1.CoinbaseResponse
	(*SetCoinbaseResponse)(nil),                    // 11: spacemesh.v1.SetCoinbaseResponse
	(*MinGasResponse)(nil),                         // 12: spacemesh.v1.MinGasResponse
	(*SetMinGasResponse)(nil),                      // 13: spacemesh.v1.SetMinGasResponse
	(*PostStatusResponse)(nil),                     // 14: spacemesh.v1.PostStatusResponse
	(*PostComputeProvidersResponse)(nil),           // 15: spacemesh.v1.PostComputeProvidersResponse
	(*CreatePostDataResponse)(nil),                 // 16: spacemesh.v1.CreatePostDataResponse
	(*StopPostDataCreationSessionResponse)(nil),    // 17: spacemesh.v1.StopPostDataCreationSessionResponse
	(*PostDataCreationProgressStreamResponse)(nil), // 18: spacemesh.v1.PostDataCreationProgressStreamResponse
}
var file_spacemesh_v1_smesher_proto_depIdxs = []int32{
	0,  // 0: spacemesh.v1.SmesherService.IsSmeshing:input_type -> google.protobuf.Empty
	0,  // 1: spacemesh.v1.SmesherService.StartSmeshing:input_type -> google.protobuf.Empty
	1,  // 2: spacemesh.v1.SmesherService.StopSmeshing:input_type -> spacemesh.v1.StopSmeshingRequest
	0,  // 3: spacemesh.v1.SmesherService.SmesherId:input_type -> google.protobuf.Empty
	0,  // 4: spacemesh.v1.SmesherService.Coinbase:input_type -> google.protobuf.Empty
	2,  // 5: spacemesh.v1.SmesherService.SetCoinbase:input_type -> spacemesh.v1.SetCoinbaseRequest
	0,  // 6: spacemesh.v1.SmesherService.MinGas:input_type -> google.protobuf.Empty
	3,  // 7: spacemesh.v1.SmesherService.SetMinGas:input_type -> spacemesh.v1.SetMinGasRequest
	0,  // 8: spacemesh.v1.SmesherService.PostStatus:input_type -> google.protobuf.Empty
	0,  // 9: spacemesh.v1.SmesherService.PostComputeProviders:input_type -> google.protobuf.Empty
	4,  // 10: spacemesh.v1.SmesherService.CreatePostData:input_type -> spacemesh.v1.CreatePostDataRequest
	5,  // 11: spacemesh.v1.SmesherService.StopPostDataCreationSession:input_type -> spacemesh.v1.StopPostDataCreationSessionRequest
	0,  // 12: spacemesh.v1.SmesherService.PostDataCreationProgressStream:input_type -> google.protobuf.Empty
	6,  // 13: spacemesh.v1.SmesherService.IsSmeshing:output_type -> spacemesh.v1.IsSmeshingResponse
	7,  // 14: spacemesh.v1.SmesherService.StartSmeshing:output_type -> spacemesh.v1.StartSmeshingResponse
	8,  // 15: spacemesh.v1.SmesherService.StopSmeshing:output_type -> spacemesh.v1.StopSmeshingResponse
	9,  // 16: spacemesh.v1.SmesherService.SmesherId:output_type -> spacemesh.v1.SmesherIdResponse
	10, // 17: spacemesh.v1.SmesherService.Coinbase:output_type -> spacemesh.v1.CoinbaseResponse
	11, // 18: spacemesh.v1.SmesherService.SetCoinbase:output_type -> spacemesh.v1.SetCoinbaseResponse
	12, // 19: spacemesh.v1.SmesherService.MinGas:output_type -> spacemesh.v1.MinGasResponse
	13, // 20: spacemesh.v1.SmesherService.SetMinGas:output_type -> spacemesh.v1.SetMinGasResponse
	14, // 21: spacemesh.v1.SmesherService.PostStatus:output_type -> spacemesh.v1.PostStatusResponse
	15, // 22: spacemesh.v1.SmesherService.PostComputeProviders:output_type -> spacemesh.v1.PostComputeProvidersResponse
	16, // 23: spacemesh.v1.SmesherService.CreatePostData:output_type -> spacemesh.v1.CreatePostDataResponse
	17, // 24: spacemesh.v1.SmesherService.StopPostDataCreationSession:output_type -> spacemesh.v1.StopPostDataCreationSessionResponse
	18, // 25: spacemesh.v1.SmesherService.PostDataCreationProgressStream:output_type -> spacemesh.v1.PostDataCreationProgressStreamResponse
	13, // [13:26] is the sub-list for method output_type
	0,  // [0:13] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_spacemesh_v1_smesher_proto_init() }
func file_spacemesh_v1_smesher_proto_init() {
	if File_spacemesh_v1_smesher_proto != nil {
		return
	}
	file_spacemesh_v1_smesher_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spacemesh_v1_smesher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spacemesh_v1_smesher_proto_goTypes,
		DependencyIndexes: file_spacemesh_v1_smesher_proto_depIdxs,
	}.Build()
	File_spacemesh_v1_smesher_proto = out.File
	file_spacemesh_v1_smesher_proto_rawDesc = nil
	file_spacemesh_v1_smesher_proto_goTypes = nil
	file_spacemesh_v1_smesher_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SmesherServiceClient is the client API for SmesherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SmesherServiceClient interface {
	// Returns true iff node is currently smeshing
	IsSmeshing(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*IsSmeshingResponse, error)
	// Start smeshing. Will return false if post data is incomplete or missing
	StartSmeshing(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StartSmeshingResponse, error)
	// Stop smeshing and optionally attempt to delete post init file(s)
	// Returns true if request is accepted by node, false if it fails
	StopSmeshing(ctx context.Context, in *StopSmeshingRequest, opts ...grpc.CallOption) (*StopSmeshingResponse, error)
	// Get the current smesher id generated by the node
	SmesherId(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*SmesherIdResponse, error)
	// Get the current coinbase
	Coinbase(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CoinbaseResponse, error)
	// Set the coinbase.
	// Returns true if request succeeds, false if it fails.
	SetCoinbase(ctx context.Context, in *SetCoinbaseRequest, opts ...grpc.CallOption) (*SetCoinbaseResponse, error)
	// Get the current min gas for including txs in blocks by this smesher
	MinGas(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*MinGasResponse, error)
	// Set a min gas units for including txs in blocks by this smesher
	// Returns true if request succeeds, false if it fails.
	SetMinGas(ctx context.Context, in *SetMinGasRequest, opts ...grpc.CallOption) (*SetMinGasResponse, error)
	// Returns post data status from the node
	PostStatus(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PostStatusResponse, error)
	// Returns a list of available post compute providers for creating post data
	PostComputeProviders(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PostComputeProvidersResponse, error)
	// Starts (or continues) a post init phase. Supports resuming a previously
	// started init session, as well as changing post params (e.g., post data size)
	// after initial setup.
	// Returns true if request is accepted by node, false if it fails
	CreatePostData(ctx context.Context, in *CreatePostDataRequest, opts ...grpc.CallOption) (*CreatePostDataResponse, error)
	// Stop an ongoing post data init phase and optionally attempt to delete
	// the post data file(s)
	// Returns true if request is accepted by node, false if it fails
	StopPostDataCreationSession(ctx context.Context, in *StopPostDataCreationSessionRequest, opts ...grpc.CallOption) (*StopPostDataCreationSessionResponse, error)
	// Returns a stream of updates to post data file(s) during the init phase
	PostDataCreationProgressStream(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (SmesherService_PostDataCreationProgressStreamClient, error)
}

type smesherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmesherServiceClient(cc grpc.ClientConnInterface) SmesherServiceClient {
	return &smesherServiceClient{cc}
}

func (c *smesherServiceClient) IsSmeshing(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*IsSmeshingResponse, error) {
	out := new(IsSmeshingResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.SmesherService/IsSmeshing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smesherServiceClient) StartSmeshing(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StartSmeshingResponse, error) {
	out := new(StartSmeshingResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.SmesherService/StartSmeshing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smesherServiceClient) StopSmeshing(ctx context.Context, in *StopSmeshingRequest, opts ...grpc.CallOption) (*StopSmeshingResponse, error) {
	out := new(StopSmeshingResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.SmesherService/StopSmeshing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smesherServiceClient) SmesherId(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*SmesherIdResponse, error) {
	out := new(SmesherIdResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.SmesherService/SmesherId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smesherServiceClient) Coinbase(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CoinbaseResponse, error) {
	out := new(CoinbaseResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.SmesherService/Coinbase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smesherServiceClient) SetCoinbase(ctx context.Context, in *SetCoinbaseRequest, opts ...grpc.CallOption) (*SetCoinbaseResponse, error) {
	out := new(SetCoinbaseResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.SmesherService/SetCoinbase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smesherServiceClient) MinGas(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*MinGasResponse, error) {
	out := new(MinGasResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.SmesherService/MinGas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smesherServiceClient) SetMinGas(ctx context.Context, in *SetMinGasRequest, opts ...grpc.CallOption) (*SetMinGasResponse, error) {
	out := new(SetMinGasResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.SmesherService/SetMinGas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smesherServiceClient) PostStatus(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PostStatusResponse, error) {
	out := new(PostStatusResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.SmesherService/PostStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smesherServiceClient) PostComputeProviders(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PostComputeProvidersResponse, error) {
	out := new(PostComputeProvidersResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.SmesherService/PostComputeProviders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smesherServiceClient) CreatePostData(ctx context.Context, in *CreatePostDataRequest, opts ...grpc.CallOption) (*CreatePostDataResponse, error) {
	out := new(CreatePostDataResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.SmesherService/CreatePostData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smesherServiceClient) StopPostDataCreationSession(ctx context.Context, in *StopPostDataCreationSessionRequest, opts ...grpc.CallOption) (*StopPostDataCreationSessionResponse, error) {
	out := new(StopPostDataCreationSessionResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.SmesherService/StopPostDataCreationSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smesherServiceClient) PostDataCreationProgressStream(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (SmesherService_PostDataCreationProgressStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SmesherService_serviceDesc.Streams[0], "/spacemesh.v1.SmesherService/PostDataCreationProgressStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &smesherServicePostDataCreationProgressStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SmesherService_PostDataCreationProgressStreamClient interface {
	Recv() (*PostDataCreationProgressStreamResponse, error)
	grpc.ClientStream
}

type smesherServicePostDataCreationProgressStreamClient struct {
	grpc.ClientStream
}

func (x *smesherServicePostDataCreationProgressStreamClient) Recv() (*PostDataCreationProgressStreamResponse, error) {
	m := new(PostDataCreationProgressStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SmesherServiceServer is the server API for SmesherService service.
type SmesherServiceServer interface {
	// Returns true iff node is currently smeshing
	IsSmeshing(context.Context, *empty.Empty) (*IsSmeshingResponse, error)
	// Start smeshing. Will return false if post data is incomplete or missing
	StartSmeshing(context.Context, *empty.Empty) (*StartSmeshingResponse, error)
	// Stop smeshing and optionally attempt to delete post init file(s)
	// Returns true if request is accepted by node, false if it fails
	StopSmeshing(context.Context, *StopSmeshingRequest) (*StopSmeshingResponse, error)
	// Get the current smesher id generated by the node
	SmesherId(context.Context, *empty.Empty) (*SmesherIdResponse, error)
	// Get the current coinbase
	Coinbase(context.Context, *empty.Empty) (*CoinbaseResponse, error)
	// Set the coinbase.
	// Returns true if request succeeds, false if it fails.
	SetCoinbase(context.Context, *SetCoinbaseRequest) (*SetCoinbaseResponse, error)
	// Get the current min gas for including txs in blocks by this smesher
	MinGas(context.Context, *empty.Empty) (*MinGasResponse, error)
	// Set a min gas units for including txs in blocks by this smesher
	// Returns true if request succeeds, false if it fails.
	SetMinGas(context.Context, *SetMinGasRequest) (*SetMinGasResponse, error)
	// Returns post data status from the node
	PostStatus(context.Context, *empty.Empty) (*PostStatusResponse, error)
	// Returns a list of available post compute providers for creating post data
	PostComputeProviders(context.Context, *empty.Empty) (*PostComputeProvidersResponse, error)
	// Starts (or continues) a post init phase. Supports resuming a previously
	// started init session, as well as changing post params (e.g., post data size)
	// after initial setup.
	// Returns true if request is accepted by node, false if it fails
	CreatePostData(context.Context, *CreatePostDataRequest) (*CreatePostDataResponse, error)
	// Stop an ongoing post data init phase and optionally attempt to delete
	// the post data file(s)
	// Returns true if request is accepted by node, false if it fails
	StopPostDataCreationSession(context.Context, *StopPostDataCreationSessionRequest) (*StopPostDataCreationSessionResponse, error)
	// Returns a stream of updates to post data file(s) during the init phase
	PostDataCreationProgressStream(*empty.Empty, SmesherService_PostDataCreationProgressStreamServer) error
}

// UnimplementedSmesherServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSmesherServiceServer struct {
}

func (*UnimplementedSmesherServiceServer) IsSmeshing(context.Context, *empty.Empty) (*IsSmeshingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsSmeshing not implemented")
}
func (*UnimplementedSmesherServiceServer) StartSmeshing(context.Context, *empty.Empty) (*StartSmeshingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartSmeshing not implemented")
}
func (*UnimplementedSmesherServiceServer) StopSmeshing(context.Context, *StopSmeshingRequest) (*StopSmeshingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopSmeshing not implemented")
}
func (*UnimplementedSmesherServiceServer) SmesherId(context.Context, *empty.Empty) (*SmesherIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SmesherId not implemented")
}
func (*UnimplementedSmesherServiceServer) Coinbase(context.Context, *empty.Empty) (*CoinbaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Coinbase not implemented")
}
func (*UnimplementedSmesherServiceServer) SetCoinbase(context.Context, *SetCoinbaseRequest) (*SetCoinbaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCoinbase not implemented")
}
func (*UnimplementedSmesherServiceServer) MinGas(context.Context, *empty.Empty) (*MinGasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MinGas not implemented")
}
func (*UnimplementedSmesherServiceServer) SetMinGas(context.Context, *SetMinGasRequest) (*SetMinGasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMinGas not implemented")
}
func (*UnimplementedSmesherServiceServer) PostStatus(context.Context, *empty.Empty) (*PostStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostStatus not implemented")
}
func (*UnimplementedSmesherServiceServer) PostComputeProviders(context.Context, *empty.Empty) (*PostComputeProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostComputeProviders not implemented")
}
func (*UnimplementedSmesherServiceServer) CreatePostData(context.Context, *CreatePostDataRequest) (*CreatePostDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePostData not implemented")
}
func (*UnimplementedSmesherServiceServer) StopPostDataCreationSession(context.Context, *StopPostDataCreationSessionRequest) (*StopPostDataCreationSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopPostDataCreationSession not implemented")
}
func (*UnimplementedSmesherServiceServer) PostDataCreationProgressStream(*empty.Empty, SmesherService_PostDataCreationProgressStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method PostDataCreationProgressStream not implemented")
}

func RegisterSmesherServiceServer(s *grpc.Server, srv SmesherServiceServer) {
	s.RegisterService(&_SmesherService_serviceDesc, srv)
}

func _SmesherService_IsSmeshing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmesherServiceServer).IsSmeshing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.SmesherService/IsSmeshing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmesherServiceServer).IsSmeshing(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmesherService_StartSmeshing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmesherServiceServer).StartSmeshing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.SmesherService/StartSmeshing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmesherServiceServer).StartSmeshing(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmesherService_StopSmeshing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopSmeshingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmesherServiceServer).StopSmeshing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.SmesherService/StopSmeshing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmesherServiceServer).StopSmeshing(ctx, req.(*StopSmeshingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmesherService_SmesherId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmesherServiceServer).SmesherId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.SmesherService/SmesherId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmesherServiceServer).SmesherId(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmesherService_Coinbase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmesherServiceServer).Coinbase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.SmesherService/Coinbase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmesherServiceServer).Coinbase(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmesherService_SetCoinbase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCoinbaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmesherServiceServer).SetCoinbase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.SmesherService/SetCoinbase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmesherServiceServer).SetCoinbase(ctx, req.(*SetCoinbaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmesherService_MinGas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmesherServiceServer).MinGas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.SmesherService/MinGas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmesherServiceServer).MinGas(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmesherService_SetMinGas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMinGasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmesherServiceServer).SetMinGas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.SmesherService/SetMinGas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmesherServiceServer).SetMinGas(ctx, req.(*SetMinGasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmesherService_PostStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmesherServiceServer).PostStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.SmesherService/PostStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmesherServiceServer).PostStatus(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmesherService_PostComputeProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmesherServiceServer).PostComputeProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.SmesherService/PostComputeProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmesherServiceServer).PostComputeProviders(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmesherService_CreatePostData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmesherServiceServer).CreatePostData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.SmesherService/CreatePostData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmesherServiceServer).CreatePostData(ctx, req.(*CreatePostDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmesherService_StopPostDataCreationSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopPostDataCreationSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmesherServiceServer).StopPostDataCreationSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.SmesherService/StopPostDataCreationSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmesherServiceServer).StopPostDataCreationSession(ctx, req.(*StopPostDataCreationSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmesherService_PostDataCreationProgressStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SmesherServiceServer).PostDataCreationProgressStream(m, &smesherServicePostDataCreationProgressStreamServer{stream})
}

type SmesherService_PostDataCreationProgressStreamServer interface {
	Send(*PostDataCreationProgressStreamResponse) error
	grpc.ServerStream
}

type smesherServicePostDataCreationProgressStreamServer struct {
	grpc.ServerStream
}

func (x *smesherServicePostDataCreationProgressStreamServer) Send(m *PostDataCreationProgressStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _SmesherService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spacemesh.v1.SmesherService",
	HandlerType: (*SmesherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsSmeshing",
			Handler:    _SmesherService_IsSmeshing_Handler,
		},
		{
			MethodName: "StartSmeshing",
			Handler:    _SmesherService_StartSmeshing_Handler,
		},
		{
			MethodName: "StopSmeshing",
			Handler:    _SmesherService_StopSmeshing_Handler,
		},
		{
			MethodName: "SmesherId",
			Handler:    _SmesherService_SmesherId_Handler,
		},
		{
			MethodName: "Coinbase",
			Handler:    _SmesherService_Coinbase_Handler,
		},
		{
			MethodName: "SetCoinbase",
			Handler:    _SmesherService_SetCoinbase_Handler,
		},
		{
			MethodName: "MinGas",
			Handler:    _SmesherService_MinGas_Handler,
		},
		{
			MethodName: "SetMinGas",
			Handler:    _SmesherService_SetMinGas_Handler,
		},
		{
			MethodName: "PostStatus",
			Handler:    _SmesherService_PostStatus_Handler,
		},
		{
			MethodName: "PostComputeProviders",
			Handler:    _SmesherService_PostComputeProviders_Handler,
		},
		{
			MethodName: "CreatePostData",
			Handler:    _SmesherService_CreatePostData_Handler,
		},
		{
			MethodName: "StopPostDataCreationSession",
			Handler:    _SmesherService_StopPostDataCreationSession_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PostDataCreationProgressStream",
			Handler:       _SmesherService_PostDataCreationProgressStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "spacemesh/v1/smesher.proto",
}
