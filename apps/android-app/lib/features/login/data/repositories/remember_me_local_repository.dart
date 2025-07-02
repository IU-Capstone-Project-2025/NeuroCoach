import 'package:android_app/features/login/domain/repositories/remember_me_repository.dart';
import 'package:shared_preferences/shared_preferences.dart';

class RememberMeLocalRepository extends RememberMeRepository {
  RememberMeLocalRepository({required SharedPreferences preferences})
    : _preferences = preferences;

  final SharedPreferences _preferences;

  @override
  Future<void> rememberUser({required String jwtToken, required String email}) async {
    await _preferences.setString('jwt', jwtToken);
    await _preferences.setString('email', email);
  }
}
