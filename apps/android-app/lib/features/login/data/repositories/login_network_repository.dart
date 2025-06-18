import 'package:android_app/features/login/domain/repositories/login_repository.dart';

class LoginNetworkRepository extends LoginRepository {
  @override
  Future<void> login() async {
    await Future.delayed(Duration(seconds: 1));
    return;
  }

  @override
  Future<void> signUp() async {
    await Future.delayed(Duration(seconds: 1));
    return;
  }

}