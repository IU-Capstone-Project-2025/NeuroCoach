import 'package:android_app/features/settings/domain/repositories/logout_repository.dart';
import 'package:shared_preferences/shared_preferences.dart';

class LogoutLocalRepository extends LogoutRepository {

  LogoutLocalRepository(super.prefs) : _preferences = prefs;

  final SharedPreferences _preferences;

  @override
  Future<bool> logout() async {
    return await _preferences.remove('jwt') && await _preferences.remove('email');
  }

}