import 'package:android_app/app/app_dependencies.dart';
import 'package:android_app/features/chat/data/repositories/chat_messages_network_repository.dart';
import 'package:android_app/features/init/data/repositories/health_network_repository.dart';
import 'package:android_app/features/init/data/repositories/init_local_repository.dart';
import 'package:android_app/features/login/data/repositories/login_network_repository.dart';
import 'package:android_app/features/login/data/repositories/remember_me_local_repository.dart';
import 'package:android_app/features/path/data/repositories/workout_path_network_repository.dart';
import 'package:android_app/features/settings/data/profile_network_repository.dart';
import 'package:android_app/features/settings/data/repositories/logout_local_repository.dart';
import 'package:shared_preferences/shared_preferences.dart';

class DependenciesFactory {
  static AppDependencies build(SharedPreferences prefs) {
    return AppDependencies(
      loginRepository: LoginNetworkRepository(),
      rememberMeRepository: RememberMeLocalRepository(preferences: prefs),
      initRepository: InitLocalRepository(preferences: prefs),
      healthRepository: HealthNetworkRepository(),
      chatMessagesRepository: ChatMessagesNetworkRepository(),
      logoutRepository: LogoutLocalRepository(prefs),
      workoutPathRepository: WorkoutPathNetworkRepository(),
      profileRepository: ProfileNetworkRepository(),
    );
  }
}
