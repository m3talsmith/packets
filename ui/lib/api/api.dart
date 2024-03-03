import 'package:grpc/grpc.dart';
import './protos/api.pbgrpc.dart';

class Api {
  String baseUrl = "localhost";

  Api._internal();
  static final Api _instance = Api._internal();
  factory Api() => _instance;
  static Api get instance => _instance;
  
  late ApiClient _apiClient;
  
  init() async {
    _createChannel();
  }

  ApiClient get apiClient => _apiClient;

  _createChannel() {
    final channel = ClientChannel(baseUrl, port: 3000, options: const ChannelOptions(credentials: ChannelCredentials.insecure()));
    _apiClient = ApiClient(channel);
  }
}