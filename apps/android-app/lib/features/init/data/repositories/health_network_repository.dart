import 'package:android_app/app/data/services/network_service.dart';
import 'package:android_app/features/init/domain/repositories/health_repository.dart';
import 'package:flutter/foundation.dart';

class HealthNetworkRepository extends HealthRepository {
  @override
  Future<bool> checkToken(String token) async {
    NetworkService().setToken(token);
    try {
      final response = await NetworkService().request(method: 'GET', path: '/api/profile');
      if (response.statusCode == 200) {
        return true;
      } else {
        NetworkService().removeToken();
        return false;
      }
    } on Object catch (e, s) {
      if (kDebugMode) print('$e, $s');
      return false;
    }
  }
}