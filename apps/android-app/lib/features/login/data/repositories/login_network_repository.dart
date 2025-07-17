import 'package:android_app/app/data/services/network_service.dart';
import 'package:android_app/features/login/domain/repositories/login_repository.dart';
import 'package:flutter/foundation.dart';

class LoginNetworkRepository extends LoginRepository {
  @override
  Future<List<String>> login(String email, String password, bool rememberMe) async {
    final response = await NetworkService().request(
      method: 'POST',
      path: '/login',
      body: {'email': email, 'password': password},
    );

    if (response.statusCode == 200) {
      final data = response.data;
      final accessToken = data['access_token'];
      final refreshToken = data['refresh_token'];
      NetworkService().setToken(data['access_token']);
      return [accessToken, refreshToken];
    } else {
      if (kDebugMode) print(response.data['message']);
    }

    return [];
  }

  @override
  Future<List<String>> signUp(
    String email,
    String password,
    int height,
    int weight,
    int age,
    String goal,
    List<String> healthIssues,
    String timeframe,
    String fitnessLevel,
    int availableMinutes,
    bool rememberMe,
  ) async {
    final response = await NetworkService().request(
      method: 'POST',
      path: '/register',
      body: {'email': email, 'password': password},
    );

    if (response.statusCode == 201) {
      final data = response.data;
      if (kDebugMode) print('Got token while signing up: ${data['access_token']}');
      NetworkService().setToken(data['access_token']);
      final addDataResponse = await NetworkService().request(
        method: 'POST',
        path: '/api/profile',
        body: {
          'height': height,
          'weight': weight,
          'age': age,
          'goal': goal,
          'health_issues': healthIssues,
          'timeframe': timeframe,
          'fitness_level': fitnessLevel,
          'available_minutes': availableMinutes,
        }
      );

      if (addDataResponse.statusCode == 200) {
        final data = response.data;
        final accessToken = data['access_token'];
        final refreshToken = data['refresh_token'];
        NetworkService().setToken(accessToken);
        return [accessToken, refreshToken];
      }
    }

    return [];
  }
}
