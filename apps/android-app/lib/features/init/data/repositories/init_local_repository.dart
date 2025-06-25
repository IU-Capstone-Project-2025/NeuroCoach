import 'package:android_app/features/init/domain/repositories/init_repository.dart';
import 'package:flutter/foundation.dart';
import 'package:shared_preferences/shared_preferences.dart';

class InitLocalRepository extends InitRepository {

  InitLocalRepository({required SharedPreferences preferences})
      : _preferences = preferences;

  final SharedPreferences _preferences;

  @override
  Future<String> getJWTToken() async {
    if (kDebugMode) print('Started fetching JWT token in repo');
    final String? token = _preferences.getString('jwt');
    return token ?? '';
  }

  @override
  Future<bool> removeJWTToken() async {
    if (kDebugMode) print('Deleting JWT from shared prefs');
    return await _preferences.remove('jwt');
  }

}