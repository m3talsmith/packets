//
//  Generated code. Do not modify.
//  source: protos/api.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use getDevicesRequestDescriptor instead')
const GetDevicesRequest$json = {
  '1': 'GetDevicesRequest',
};

/// Descriptor for `GetDevicesRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getDevicesRequestDescriptor = $convert.base64Decode(
    'ChFHZXREZXZpY2VzUmVxdWVzdA==');

@$core.Deprecated('Use getDevicesResponseDescriptor instead')
const GetDevicesResponse$json = {
  '1': 'GetDevicesResponse',
  '2': [
    {'1': 'devices', '3': 1, '4': 3, '5': 11, '6': '.api.Device', '10': 'devices'},
  ],
};

/// Descriptor for `GetDevicesResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getDevicesResponseDescriptor = $convert.base64Decode(
    'ChJHZXREZXZpY2VzUmVzcG9uc2USJQoHZGV2aWNlcxgBIAMoCzILLmFwaS5EZXZpY2VSB2Rldm'
    'ljZXM=');

@$core.Deprecated('Use addressDescriptor instead')
const Address$json = {
  '1': 'Address',
  '2': [
    {'1': 'ip', '3': 1, '4': 1, '5': 9, '10': 'ip'},
    {'1': 'netmask', '3': 2, '4': 1, '5': 9, '10': 'netmask'},
    {'1': 'broadcast', '3': 3, '4': 1, '5': 9, '10': 'broadcast'},
    {'1': 'p2p', '3': 4, '4': 1, '5': 9, '10': 'p2p'},
  ],
};

/// Descriptor for `Address`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List addressDescriptor = $convert.base64Decode(
    'CgdBZGRyZXNzEg4KAmlwGAEgASgJUgJpcBIYCgduZXRtYXNrGAIgASgJUgduZXRtYXNrEhwKCW'
    'Jyb2FkY2FzdBgDIAEoCVIJYnJvYWRjYXN0EhAKA3AycBgEIAEoCVIDcDJw');

@$core.Deprecated('Use deviceDescriptor instead')
const Device$json = {
  '1': 'Device',
  '2': [
    {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
    {'1': 'description', '3': 2, '4': 1, '5': 9, '10': 'description'},
    {'1': 'flags', '3': 3, '4': 1, '5': 5, '10': 'flags'},
    {'1': 'addresses', '3': 4, '4': 3, '5': 11, '6': '.api.Address', '10': 'addresses'},
  ],
};

/// Descriptor for `Device`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deviceDescriptor = $convert.base64Decode(
    'CgZEZXZpY2USEgoEbmFtZRgBIAEoCVIEbmFtZRIgCgtkZXNjcmlwdGlvbhgCIAEoCVILZGVzY3'
    'JpcHRpb24SFAoFZmxhZ3MYAyABKAVSBWZsYWdzEioKCWFkZHJlc3NlcxgEIAMoCzIMLmFwaS5B'
    'ZGRyZXNzUglhZGRyZXNzZXM=');

