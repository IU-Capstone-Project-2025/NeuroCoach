import 'package:android_app/features/login/domain/repositories/login_repository.dart';
import 'package:android_app/features/login/domain/repositories/remember_me_repository.dart';
import 'package:flutter/material.dart';

@immutable
class AppDependencies {
  const AppDependencies({
    required this.loginRepository,
    required this.rememberMeRepository,
  });

  final LoginRepository loginRepository;
  final RememberMeRepository rememberMeRepository;
}
