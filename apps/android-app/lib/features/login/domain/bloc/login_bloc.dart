import 'package:android_app/features/login/domain/repositories/remember_me_repository.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../repositories/login_repository.dart';

class LoginBloc extends Bloc<LoginEvent, LoginState> {
  LoginBloc({
    required LoginRepository loginRepository,
    required RememberMeRepository rememberMeRepository,
  }) : _loginRepository = loginRepository,
       _rememberMeRepository = rememberMeRepository,
       super(const LoginStateInitial()) {
    on<LoginEvent>(
      (event, emit) => switch (event) {
        LoginEventLogin() => _onLogin(event, emit),
        LoginEventSignUp() => _onSignUp(event, emit),
        LoginEventReset() => emit(const LoginStateInitial()),
      },
    );
  }

  final LoginRepository _loginRepository;
  final RememberMeRepository _rememberMeRepository;

  Future<void> _onLogin(LoginEventLogin event, Emitter<LoginState> emit) async {
    emit(const LoginStateLoading());
    try {
      final String token = await _loginRepository.login(
        event.email,
        event.password,
        event.rememberMe,
      );

      if (event.rememberMe && token.isNotEmpty) {
        _rememberMeRepository.rememberUser(jwtToken: token);
      }
      emit(const LoginStateLoaded());
    } on Object catch (e, s) {
      if (kDebugMode) print('$e, $s');
      emit(const LoginStateError());
    }
  }

  Future<void> _onSignUp(
    LoginEventSignUp event,
    Emitter<LoginState> emit,
  ) async {
    emit(const LoginStateLoading());
    try {
      final String token = await _loginRepository.signUp(
        event.email,
        event.password,
        event.height,
        event.weight,
        event.age,
        event.goal,
        event.healthIssues,
        event.timeframe,
        event.fitnessLevel,
        event.availableMinutes,
        event.rememberMe,
      );

      if (event.rememberMe && token.isNotEmpty) {
        _rememberMeRepository.rememberUser(jwtToken: token);
      }
      emit(const LoginStateLoaded());
    } on Object catch (e, s) {
      if (kDebugMode) print('$e, $s');
      emit(LoginStateError());
    }
  }
}

sealed class LoginEvent {
  const LoginEvent();
}

class LoginEventLogin extends LoginEvent {
  const LoginEventLogin({
    required this.email,
    required this.password,
    required this.rememberMe,
  });

  final String email;
  final String password;
  final bool rememberMe;
}

class LoginEventSignUp extends LoginEvent {
  const LoginEventSignUp({
    required this.email,
    required this.password,
    required this.height,
    required this.weight,
    required this.age,
    required this.goal,
    required this.healthIssues,
    required this.timeframe,
    required this.fitnessLevel,
    required this.availableMinutes,
    required this.rememberMe,
  });

  final String email;
  final String password;
  final int height;
  final int weight;
  final int age;
  final String goal;
  final List<String> healthIssues;
  final String timeframe;
  final String fitnessLevel;
  final int availableMinutes;
  final bool rememberMe;
}

class LoginEventReset extends LoginEvent {
  const LoginEventReset();
}

sealed class LoginState {
  const LoginState();
}

class LoginStateInitial extends LoginState {
  const LoginStateInitial();
}

class LoginStateLoading extends LoginState {
  const LoginStateLoading();
}

class LoginStateLoaded extends LoginState {
  const LoginStateLoaded();
}

class LoginStateError extends LoginState {
  const LoginStateError();
}
