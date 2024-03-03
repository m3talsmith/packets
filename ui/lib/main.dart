import 'dart:developer';

import 'package:flutter/material.dart';
import 'package:grpc/grpc.dart';
import 'package:ui/api/api.dart';
import 'package:ui/api/protos/api.pbgrpc.dart';

void main() {
  WidgetsFlutterBinding.ensureInitialized();
  Api().init();

  runApp(const App());
}

class App extends StatelessWidget {
  const App({super.key});

  @override
  Widget build(context) {
    return MaterialApp(
      title: 'Packets',
      theme: ThemeData(
        useMaterial3: true,
        colorSchemeSeed: Colors.amber
      ),
      home: const Home(),
    );
  }
}

class Home extends StatefulWidget {
  const Home({super.key});

  @override
  State<Home> createState() => _HomeState();
}

class _HomeState extends State<Home> {
  List<Device> devices = [];

  getDevices() async {
    try {
      var resp = await Api.instance.apiClient.getDevices(GetDevicesRequest());
      log('devices: ${resp.devices}');
      setState(() => devices = resp.devices);
    } on GrpcError catch (e) {
      log('grpc error: $e');
    } catch (e) {
      log('general error: $e');
    }
  }

  @override
  Widget build(context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Home'),),
      body:ListView(
            children: [
              ElevatedButton(onPressed: getDevices, child: const Text('Get Devices')),
              ...devices.map((e) => ListTile(
                dense: true,
              title: Text(e.name),
              subtitle: Text(e.description),
              )),
            ]
          )
    );
  }
}