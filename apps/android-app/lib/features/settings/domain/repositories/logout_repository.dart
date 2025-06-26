import 'package:shared_preferences/shared_preferences.dart';

abstract class LogoutRepository {

  LogoutRepository(SharedPreferences prefs);

  Future<bool> logout();
}