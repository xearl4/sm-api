// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: spacemesh/v1/mesh.proto

package v1

import (
	context "context"
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

var File_spacemesh_v1_mesh_proto protoreflect.FileDescriptor

var file_spacemesh_v1_mesh_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9c, 0x08, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x68, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69,
	0x73, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x54, 0x69, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x54, 0x69,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x0c, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x55, 0x0a, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x45, 0x70, 0x6f, 0x63,
	0x68, 0x12, 0x21, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x45, 0x70, 0x6f, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x09, 0x47, 0x65, 0x6e, 0x65,
	0x73, 0x69, 0x73, 0x49, 0x44, 0x12, 0x1e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x49, 0x44, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x4e,
	0x75, 0x6d, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x23, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x4e, 0x75, 0x6d,
	0x4c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x70, 0x6f,
	0x63, 0x68, 0x4e, 0x75, 0x6d, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0d, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x79, 0x0a,
	0x18, 0x4d, 0x61, 0x78, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x2d, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x78, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x78, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6d, 0x0a, 0x14, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x29, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4d, 0x65, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0b, 0x4c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x20, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x72, 0x0a, 0x15, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x2a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x68, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12,
	0x54, 0x0a, 0x0b, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x20,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61,
	0x79, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_spacemesh_v1_mesh_proto_goTypes = []interface{}{
	(*GenesisTimeRequest)(nil),               // 0: spacemesh.v1.GenesisTimeRequest
	(*CurrentLayerRequest)(nil),              // 1: spacemesh.v1.CurrentLayerRequest
	(*CurrentEpochRequest)(nil),              // 2: spacemesh.v1.CurrentEpochRequest
	(*GenesisIDRequest)(nil),                 // 3: spacemesh.v1.GenesisIDRequest
	(*EpochNumLayersRequest)(nil),            // 4: spacemesh.v1.EpochNumLayersRequest
	(*LayerDurationRequest)(nil),             // 5: spacemesh.v1.LayerDurationRequest
	(*MaxTransactionsPerSecondRequest)(nil),  // 6: spacemesh.v1.MaxTransactionsPerSecondRequest
	(*AccountMeshDataQueryRequest)(nil),      // 7: spacemesh.v1.AccountMeshDataQueryRequest
	(*LayersQueryRequest)(nil),               // 8: spacemesh.v1.LayersQueryRequest
	(*AccountMeshDataStreamRequest)(nil),     // 9: spacemesh.v1.AccountMeshDataStreamRequest
	(*LayerStreamRequest)(nil),               // 10: spacemesh.v1.LayerStreamRequest
	(*GenesisTimeResponse)(nil),              // 11: spacemesh.v1.GenesisTimeResponse
	(*CurrentLayerResponse)(nil),             // 12: spacemesh.v1.CurrentLayerResponse
	(*CurrentEpochResponse)(nil),             // 13: spacemesh.v1.CurrentEpochResponse
	(*GenesisIDResponse)(nil),                // 14: spacemesh.v1.GenesisIDResponse
	(*EpochNumLayersResponse)(nil),           // 15: spacemesh.v1.EpochNumLayersResponse
	(*LayerDurationResponse)(nil),            // 16: spacemesh.v1.LayerDurationResponse
	(*MaxTransactionsPerSecondResponse)(nil), // 17: spacemesh.v1.MaxTransactionsPerSecondResponse
	(*AccountMeshDataQueryResponse)(nil),     // 18: spacemesh.v1.AccountMeshDataQueryResponse
	(*LayersQueryResponse)(nil),              // 19: spacemesh.v1.LayersQueryResponse
	(*AccountMeshDataStreamResponse)(nil),    // 20: spacemesh.v1.AccountMeshDataStreamResponse
	(*LayerStreamResponse)(nil),              // 21: spacemesh.v1.LayerStreamResponse
}
var file_spacemesh_v1_mesh_proto_depIdxs = []int32{
	0,  // 0: spacemesh.v1.MeshService.GenesisTime:input_type -> spacemesh.v1.GenesisTimeRequest
	1,  // 1: spacemesh.v1.MeshService.CurrentLayer:input_type -> spacemesh.v1.CurrentLayerRequest
	2,  // 2: spacemesh.v1.MeshService.CurrentEpoch:input_type -> spacemesh.v1.CurrentEpochRequest
	3,  // 3: spacemesh.v1.MeshService.GenesisID:input_type -> spacemesh.v1.GenesisIDRequest
	4,  // 4: spacemesh.v1.MeshService.EpochNumLayers:input_type -> spacemesh.v1.EpochNumLayersRequest
	5,  // 5: spacemesh.v1.MeshService.LayerDuration:input_type -> spacemesh.v1.LayerDurationRequest
	6,  // 6: spacemesh.v1.MeshService.MaxTransactionsPerSecond:input_type -> spacemesh.v1.MaxTransactionsPerSecondRequest
	7,  // 7: spacemesh.v1.MeshService.AccountMeshDataQuery:input_type -> spacemesh.v1.AccountMeshDataQueryRequest
	8,  // 8: spacemesh.v1.MeshService.LayersQuery:input_type -> spacemesh.v1.LayersQueryRequest
	9,  // 9: spacemesh.v1.MeshService.AccountMeshDataStream:input_type -> spacemesh.v1.AccountMeshDataStreamRequest
	10, // 10: spacemesh.v1.MeshService.LayerStream:input_type -> spacemesh.v1.LayerStreamRequest
	11, // 11: spacemesh.v1.MeshService.GenesisTime:output_type -> spacemesh.v1.GenesisTimeResponse
	12, // 12: spacemesh.v1.MeshService.CurrentLayer:output_type -> spacemesh.v1.CurrentLayerResponse
	13, // 13: spacemesh.v1.MeshService.CurrentEpoch:output_type -> spacemesh.v1.CurrentEpochResponse
	14, // 14: spacemesh.v1.MeshService.GenesisID:output_type -> spacemesh.v1.GenesisIDResponse
	15, // 15: spacemesh.v1.MeshService.EpochNumLayers:output_type -> spacemesh.v1.EpochNumLayersResponse
	16, // 16: spacemesh.v1.MeshService.LayerDuration:output_type -> spacemesh.v1.LayerDurationResponse
	17, // 17: spacemesh.v1.MeshService.MaxTransactionsPerSecond:output_type -> spacemesh.v1.MaxTransactionsPerSecondResponse
	18, // 18: spacemesh.v1.MeshService.AccountMeshDataQuery:output_type -> spacemesh.v1.AccountMeshDataQueryResponse
	19, // 19: spacemesh.v1.MeshService.LayersQuery:output_type -> spacemesh.v1.LayersQueryResponse
	20, // 20: spacemesh.v1.MeshService.AccountMeshDataStream:output_type -> spacemesh.v1.AccountMeshDataStreamResponse
	21, // 21: spacemesh.v1.MeshService.LayerStream:output_type -> spacemesh.v1.LayerStreamResponse
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_spacemesh_v1_mesh_proto_init() }
func file_spacemesh_v1_mesh_proto_init() {
	if File_spacemesh_v1_mesh_proto != nil {
		return
	}
	file_spacemesh_v1_mesh_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spacemesh_v1_mesh_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spacemesh_v1_mesh_proto_goTypes,
		DependencyIndexes: file_spacemesh_v1_mesh_proto_depIdxs,
	}.Build()
	File_spacemesh_v1_mesh_proto = out.File
	file_spacemesh_v1_mesh_proto_rawDesc = nil
	file_spacemesh_v1_mesh_proto_goTypes = nil
	file_spacemesh_v1_mesh_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MeshServiceClient is the client API for MeshService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MeshServiceClient interface {
	// Network genesis time as unix epoch time
	GenesisTime(ctx context.Context, in *GenesisTimeRequest, opts ...grpc.CallOption) (*GenesisTimeResponse, error)
	// Current layer number
	CurrentLayer(ctx context.Context, in *CurrentLayerRequest, opts ...grpc.CallOption) (*CurrentLayerResponse, error)
	// Current epoch number
	CurrentEpoch(ctx context.Context, in *CurrentEpochRequest, opts ...grpc.CallOption) (*CurrentEpochResponse, error)
	// Genesis ID
	GenesisID(ctx context.Context, in *GenesisIDRequest, opts ...grpc.CallOption) (*GenesisIDResponse, error)
	// Number of layers per epoch (a network parameter)
	EpochNumLayers(ctx context.Context, in *EpochNumLayersRequest, opts ...grpc.CallOption) (*EpochNumLayersResponse, error)
	// Layer duration (a network parameter)
	LayerDuration(ctx context.Context, in *LayerDurationRequest, opts ...grpc.CallOption) (*LayerDurationResponse, error)
	// Number of transactions per second (a network parameter)
	MaxTransactionsPerSecond(ctx context.Context, in *MaxTransactionsPerSecondRequest, opts ...grpc.CallOption) (*MaxTransactionsPerSecondResponse, error)
	// Get account data query
	AccountMeshDataQuery(ctx context.Context, in *AccountMeshDataQueryRequest, opts ...grpc.CallOption) (*AccountMeshDataQueryResponse, error)
	// Layers data query
	LayersQuery(ctx context.Context, in *LayersQueryRequest, opts ...grpc.CallOption) (*LayersQueryResponse, error)
	// A stream of transactions and activations from an account.
	// Includes simple coin transactions with the account as the destination.
	AccountMeshDataStream(ctx context.Context, in *AccountMeshDataStreamRequest, opts ...grpc.CallOption) (MeshService_AccountMeshDataStreamClient, error)
	// Layer with blocks, transactions and activations
	// Sent each time layer data changes. Designed for heavy-duty clients.
	LayerStream(ctx context.Context, in *LayerStreamRequest, opts ...grpc.CallOption) (MeshService_LayerStreamClient, error)
}

type meshServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMeshServiceClient(cc grpc.ClientConnInterface) MeshServiceClient {
	return &meshServiceClient{cc}
}

func (c *meshServiceClient) GenesisTime(ctx context.Context, in *GenesisTimeRequest, opts ...grpc.CallOption) (*GenesisTimeResponse, error) {
	out := new(GenesisTimeResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.MeshService/GenesisTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshServiceClient) CurrentLayer(ctx context.Context, in *CurrentLayerRequest, opts ...grpc.CallOption) (*CurrentLayerResponse, error) {
	out := new(CurrentLayerResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.MeshService/CurrentLayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshServiceClient) CurrentEpoch(ctx context.Context, in *CurrentEpochRequest, opts ...grpc.CallOption) (*CurrentEpochResponse, error) {
	out := new(CurrentEpochResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.MeshService/CurrentEpoch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshServiceClient) GenesisID(ctx context.Context, in *GenesisIDRequest, opts ...grpc.CallOption) (*GenesisIDResponse, error) {
	out := new(GenesisIDResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.MeshService/GenesisID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshServiceClient) EpochNumLayers(ctx context.Context, in *EpochNumLayersRequest, opts ...grpc.CallOption) (*EpochNumLayersResponse, error) {
	out := new(EpochNumLayersResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.MeshService/EpochNumLayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshServiceClient) LayerDuration(ctx context.Context, in *LayerDurationRequest, opts ...grpc.CallOption) (*LayerDurationResponse, error) {
	out := new(LayerDurationResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.MeshService/LayerDuration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshServiceClient) MaxTransactionsPerSecond(ctx context.Context, in *MaxTransactionsPerSecondRequest, opts ...grpc.CallOption) (*MaxTransactionsPerSecondResponse, error) {
	out := new(MaxTransactionsPerSecondResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.MeshService/MaxTransactionsPerSecond", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshServiceClient) AccountMeshDataQuery(ctx context.Context, in *AccountMeshDataQueryRequest, opts ...grpc.CallOption) (*AccountMeshDataQueryResponse, error) {
	out := new(AccountMeshDataQueryResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.MeshService/AccountMeshDataQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshServiceClient) LayersQuery(ctx context.Context, in *LayersQueryRequest, opts ...grpc.CallOption) (*LayersQueryResponse, error) {
	out := new(LayersQueryResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.MeshService/LayersQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshServiceClient) AccountMeshDataStream(ctx context.Context, in *AccountMeshDataStreamRequest, opts ...grpc.CallOption) (MeshService_AccountMeshDataStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MeshService_serviceDesc.Streams[0], "/spacemesh.v1.MeshService/AccountMeshDataStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &meshServiceAccountMeshDataStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MeshService_AccountMeshDataStreamClient interface {
	Recv() (*AccountMeshDataStreamResponse, error)
	grpc.ClientStream
}

type meshServiceAccountMeshDataStreamClient struct {
	grpc.ClientStream
}

func (x *meshServiceAccountMeshDataStreamClient) Recv() (*AccountMeshDataStreamResponse, error) {
	m := new(AccountMeshDataStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *meshServiceClient) LayerStream(ctx context.Context, in *LayerStreamRequest, opts ...grpc.CallOption) (MeshService_LayerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MeshService_serviceDesc.Streams[1], "/spacemesh.v1.MeshService/LayerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &meshServiceLayerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MeshService_LayerStreamClient interface {
	Recv() (*LayerStreamResponse, error)
	grpc.ClientStream
}

type meshServiceLayerStreamClient struct {
	grpc.ClientStream
}

func (x *meshServiceLayerStreamClient) Recv() (*LayerStreamResponse, error) {
	m := new(LayerStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MeshServiceServer is the server API for MeshService service.
type MeshServiceServer interface {
	// Network genesis time as unix epoch time
	GenesisTime(context.Context, *GenesisTimeRequest) (*GenesisTimeResponse, error)
	// Current layer number
	CurrentLayer(context.Context, *CurrentLayerRequest) (*CurrentLayerResponse, error)
	// Current epoch number
	CurrentEpoch(context.Context, *CurrentEpochRequest) (*CurrentEpochResponse, error)
	// Genesis ID
	GenesisID(context.Context, *GenesisIDRequest) (*GenesisIDResponse, error)
	// Number of layers per epoch (a network parameter)
	EpochNumLayers(context.Context, *EpochNumLayersRequest) (*EpochNumLayersResponse, error)
	// Layer duration (a network parameter)
	LayerDuration(context.Context, *LayerDurationRequest) (*LayerDurationResponse, error)
	// Number of transactions per second (a network parameter)
	MaxTransactionsPerSecond(context.Context, *MaxTransactionsPerSecondRequest) (*MaxTransactionsPerSecondResponse, error)
	// Get account data query
	AccountMeshDataQuery(context.Context, *AccountMeshDataQueryRequest) (*AccountMeshDataQueryResponse, error)
	// Layers data query
	LayersQuery(context.Context, *LayersQueryRequest) (*LayersQueryResponse, error)
	// A stream of transactions and activations from an account.
	// Includes simple coin transactions with the account as the destination.
	AccountMeshDataStream(*AccountMeshDataStreamRequest, MeshService_AccountMeshDataStreamServer) error
	// Layer with blocks, transactions and activations
	// Sent each time layer data changes. Designed for heavy-duty clients.
	LayerStream(*LayerStreamRequest, MeshService_LayerStreamServer) error
}

// UnimplementedMeshServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMeshServiceServer struct {
}

func (*UnimplementedMeshServiceServer) GenesisTime(context.Context, *GenesisTimeRequest) (*GenesisTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenesisTime not implemented")
}
func (*UnimplementedMeshServiceServer) CurrentLayer(context.Context, *CurrentLayerRequest) (*CurrentLayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentLayer not implemented")
}
func (*UnimplementedMeshServiceServer) CurrentEpoch(context.Context, *CurrentEpochRequest) (*CurrentEpochResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentEpoch not implemented")
}
func (*UnimplementedMeshServiceServer) GenesisID(context.Context, *GenesisIDRequest) (*GenesisIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenesisID not implemented")
}
func (*UnimplementedMeshServiceServer) EpochNumLayers(context.Context, *EpochNumLayersRequest) (*EpochNumLayersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EpochNumLayers not implemented")
}
func (*UnimplementedMeshServiceServer) LayerDuration(context.Context, *LayerDurationRequest) (*LayerDurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LayerDuration not implemented")
}
func (*UnimplementedMeshServiceServer) MaxTransactionsPerSecond(context.Context, *MaxTransactionsPerSecondRequest) (*MaxTransactionsPerSecondResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MaxTransactionsPerSecond not implemented")
}
func (*UnimplementedMeshServiceServer) AccountMeshDataQuery(context.Context, *AccountMeshDataQueryRequest) (*AccountMeshDataQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountMeshDataQuery not implemented")
}
func (*UnimplementedMeshServiceServer) LayersQuery(context.Context, *LayersQueryRequest) (*LayersQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LayersQuery not implemented")
}
func (*UnimplementedMeshServiceServer) AccountMeshDataStream(*AccountMeshDataStreamRequest, MeshService_AccountMeshDataStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AccountMeshDataStream not implemented")
}
func (*UnimplementedMeshServiceServer) LayerStream(*LayerStreamRequest, MeshService_LayerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method LayerStream not implemented")
}

func RegisterMeshServiceServer(s *grpc.Server, srv MeshServiceServer) {
	s.RegisterService(&_MeshService_serviceDesc, srv)
}

func _MeshService_GenesisTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenesisTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServiceServer).GenesisTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.MeshService/GenesisTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServiceServer).GenesisTime(ctx, req.(*GenesisTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshService_CurrentLayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrentLayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServiceServer).CurrentLayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.MeshService/CurrentLayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServiceServer).CurrentLayer(ctx, req.(*CurrentLayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshService_CurrentEpoch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrentEpochRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServiceServer).CurrentEpoch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.MeshService/CurrentEpoch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServiceServer).CurrentEpoch(ctx, req.(*CurrentEpochRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshService_GenesisID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenesisIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServiceServer).GenesisID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.MeshService/GenesisID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServiceServer).GenesisID(ctx, req.(*GenesisIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshService_EpochNumLayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EpochNumLayersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServiceServer).EpochNumLayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.MeshService/EpochNumLayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServiceServer).EpochNumLayers(ctx, req.(*EpochNumLayersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshService_LayerDuration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LayerDurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServiceServer).LayerDuration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.MeshService/LayerDuration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServiceServer).LayerDuration(ctx, req.(*LayerDurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshService_MaxTransactionsPerSecond_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MaxTransactionsPerSecondRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServiceServer).MaxTransactionsPerSecond(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.MeshService/MaxTransactionsPerSecond",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServiceServer).MaxTransactionsPerSecond(ctx, req.(*MaxTransactionsPerSecondRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshService_AccountMeshDataQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountMeshDataQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServiceServer).AccountMeshDataQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.MeshService/AccountMeshDataQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServiceServer).AccountMeshDataQuery(ctx, req.(*AccountMeshDataQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshService_LayersQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LayersQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServiceServer).LayersQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.MeshService/LayersQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServiceServer).LayersQuery(ctx, req.(*LayersQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshService_AccountMeshDataStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AccountMeshDataStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MeshServiceServer).AccountMeshDataStream(m, &meshServiceAccountMeshDataStreamServer{stream})
}

type MeshService_AccountMeshDataStreamServer interface {
	Send(*AccountMeshDataStreamResponse) error
	grpc.ServerStream
}

type meshServiceAccountMeshDataStreamServer struct {
	grpc.ServerStream
}

func (x *meshServiceAccountMeshDataStreamServer) Send(m *AccountMeshDataStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MeshService_LayerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LayerStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MeshServiceServer).LayerStream(m, &meshServiceLayerStreamServer{stream})
}

type MeshService_LayerStreamServer interface {
	Send(*LayerStreamResponse) error
	grpc.ServerStream
}

type meshServiceLayerStreamServer struct {
	grpc.ServerStream
}

func (x *meshServiceLayerStreamServer) Send(m *LayerStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _MeshService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spacemesh.v1.MeshService",
	HandlerType: (*MeshServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenesisTime",
			Handler:    _MeshService_GenesisTime_Handler,
		},
		{
			MethodName: "CurrentLayer",
			Handler:    _MeshService_CurrentLayer_Handler,
		},
		{
			MethodName: "CurrentEpoch",
			Handler:    _MeshService_CurrentEpoch_Handler,
		},
		{
			MethodName: "GenesisID",
			Handler:    _MeshService_GenesisID_Handler,
		},
		{
			MethodName: "EpochNumLayers",
			Handler:    _MeshService_EpochNumLayers_Handler,
		},
		{
			MethodName: "LayerDuration",
			Handler:    _MeshService_LayerDuration_Handler,
		},
		{
			MethodName: "MaxTransactionsPerSecond",
			Handler:    _MeshService_MaxTransactionsPerSecond_Handler,
		},
		{
			MethodName: "AccountMeshDataQuery",
			Handler:    _MeshService_AccountMeshDataQuery_Handler,
		},
		{
			MethodName: "LayersQuery",
			Handler:    _MeshService_LayersQuery_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AccountMeshDataStream",
			Handler:       _MeshService_AccountMeshDataStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LayerStream",
			Handler:       _MeshService_LayerStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "spacemesh/v1/mesh.proto",
}
