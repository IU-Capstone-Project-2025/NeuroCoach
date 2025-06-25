import 'package:android_app/app/data/services/network_service.dart';
import 'package:android_app/features/login/domain/repositories/login_repository.dart';

class LoginNetworkRepository extends LoginRepository {
  @override
  Future<String> login(String email, String password, bool rememberMe) async {
    final response = await NetworkService().request(
      method: 'POST',
      path: '/login',
      body: {'email': email, 'password': password},
    );

    if (response.statusCode == 200) {
      final data = response.data;
      NetworkService().setToken(data['token']);
      return data['token'];
    }

    return '';
  }

  @override
  Future<String> signUp(
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
      path: '/signup',
      body: {
        'email': email,
        'password': password,
        'height': height,
        'weight': weight,
        'age': age,
        'goal': goal,
        'health_issues': healthIssues,
        'timeframe': timeframe,
        'fitness_level': fitnessLevel,
        'available_minutes': availableMinutes,
      },
    );

    if (response.statusCode == 200) {
      final data = response.data;
      NetworkService().setToken(data['token']);
      return data['token'];
    }

    return '';
  }
}
