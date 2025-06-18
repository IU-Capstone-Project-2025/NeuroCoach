import 'package:android_app/features/login/domain/repositories/login_repository.dart';
import 'package:flutter/material.dart';

@immutable
class AppDependencies {
  const AppDependencies({required this.loginRepository});

  final LoginRepository loginRepository;
}
