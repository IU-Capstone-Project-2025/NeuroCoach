import 'package:android_app/app/app_dependencies.dart';
import 'package:android_app/features/login/data/repositories/login_network_repository.dart';

class DependenciesFactory {
  static AppDependencies build() {
    return AppDependencies(loginRepository: LoginNetworkRepository());
  }
}