//
//  Generated code. Do not modify.
//  source: protos/api.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'package:protobuf/protobuf.dart' as $pb;

import 'api.pb.dart' as $0;

export 'api.pb.dart';

@$pb.GrpcServiceName('api.Api')
class ApiClient extends $grpc.Client {
  static final _$getDevices = $grpc.ClientMethod<$0.GetDevicesRequest, $0.GetDevicesResponse>(
      '/api.Api/GetDevices',
      ($0.GetDevicesRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.GetDevicesResponse.fromBuffer(value));

  ApiClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseFuture<$0.GetDevicesResponse> getDevices($0.GetDevicesRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getDevices, request, options: options);
  }
}

@$pb.GrpcServiceName('api.Api')
abstract class ApiServiceBase extends $grpc.Service {
  $core.String get $name => 'api.Api';

  ApiServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.GetDevicesRequest, $0.GetDevicesResponse>(
        'GetDevices',
        getDevices_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.GetDevicesRequest.fromBuffer(value),
        ($0.GetDevicesResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.GetDevicesResponse> getDevices_Pre($grpc.ServiceCall call, $async.Future<$0.GetDevicesRequest> request) async {
    return getDevices(call, await request);
  }

  $async.Future<$0.GetDevicesResponse> getDevices($grpc.ServiceCall call, $0.GetDevicesRequest request);
}
