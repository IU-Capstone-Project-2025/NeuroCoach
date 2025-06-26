import 'package:android_app/app/domain/entities/app_config.dart';
import 'package:shared_preferences/shared_preferences.dart';

class AppConfigFactory {
  static AppConfig build(SharedPreferences prefs) {
    return AppConfig(email: prefs.getString('email') ?? '');
  }
}
