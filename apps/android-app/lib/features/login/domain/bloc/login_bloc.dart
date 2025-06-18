import 'package:flutter/foundation.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../repositories/login_repository.dart';

class LoginBloc extends Bloc<LoginEvent, LoginState> {
  LoginBloc({required LoginRepository loginRepository})
    : _loginRepository = loginRepository,
      super(const LoginStateInitial()) {
    on<LoginEvent>(
      (event, emit) => switch (event) {
        LoginEventLogin() => _onLogin(event, emit),
        LoginEventSignUp() => _onSignUp(event, emit),
      },
    );
  }

  Future<void> _onLogin(LoginEventLogin event, Emitter<LoginState> emit) async {
    emit(const LoginStateLoading());
    try {
      _loginRepository.login();
      emit(LoginStateLoaded());
    } on Object catch (e, s) {
      if (kDebugMode) print('$e, $s');
      emit(LoginStateError());
    }
  }

  Future<void> _onSignUp(
    LoginEventSignUp event,
    Emitter<LoginState> emit,
  ) async {
    emit(const LoginStateLoading());
    try {
      _loginRepository.signUp();
      emit(LoginStateLoaded());
    } on Object catch (e, s) {
      if (kDebugMode) print('$e, $s');
      emit(LoginStateError());
    }
  }

  final LoginRepository _loginRepository;
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
  const LoginEventSignUp(this.email, this.password);

  final String email;
  final String password;
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
