import 'package:android_app/features/init/domain/repositories/health_repository.dart';
import 'package:android_app/features/init/domain/repositories/init_repository.dart';
import 'package:android_app/features/login/domain/repositories/login_repository.dart';
import 'package:android_app/features/login/domain/repositories/remember_me_repository.dart';
import 'package:flutter/material.dart';

@immutable
class AppDependencies {
  const AppDependencies({
    required this.loginRepository,
    required this.rememberMeRepository,
    required this.initRepository,
    required this.healthRepository,
  });

  final LoginRepository loginRepository;
  final RememberMeRepository rememberMeRepository;
  final InitRepository initRepository;
  final HealthRepository healthRepository;
}
