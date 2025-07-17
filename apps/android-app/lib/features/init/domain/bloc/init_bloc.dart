import 'package:android_app/app/data/services/network_service.dart';
import 'package:android_app/features/init/domain/repositories/health_repository.dart';
import 'package:android_app/features/init/domain/repositories/init_repository.dart';
import 'package:android_app/features/login/domain/repositories/remember_me_repository.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class InitBloc extends Bloc<InitEvent, InitState> {
  InitBloc({
    required InitRepository initRepository,
    required HealthRepository healthRepository,
    required RememberMeRepository rememberMeRepository,
  }) : _initRepository = initRepository,
       _healthRepository = healthRepository,
       _rememberMeRepository = rememberMeRepository,
       super(const InitStateInitial()) {
    on<InitEvent>(
      (event, emit) => switch (event) {
        InitEventCheck() => _onCheck(event, emit),
      },
    );
  }

  final InitRepository _initRepository;
  final HealthRepository _healthRepository;
  final RememberMeRepository _rememberMeRepository;

  Future<void> _onCheck(InitEventCheck event, Emitter<InitState> emit) async {
    final String token = await _initRepository.getJWTToken();
    final String refreshToken = await _initRepository.getRefreshToken();
    if (kDebugMode) print('Got JWT token: $token');
    if (kDebugMode) print('Got refresh token: $refreshToken');
    if (token.isEmpty || refreshToken.isEmpty) {
      NetworkService().removeToken();
      await _initRepository.removeJWTToken();
      await _initRepository.removeRefreshToken();
      emit(const InitStateUnauthenticated());
      return;
    }

    if (await _healthRepository.checkToken(token)) {
      emit(const InitStateAuthenticated());
      return;
    } else {
      final tokens = await _healthRepository.refresh(refreshToken);
      if (tokens.isNotEmpty) {
        _rememberMeRepository.rememberUser(
          jwtToken: tokens[0],
          refreshToken: tokens[1],
          email: tokens[2],
        );
      }
      _initRepository.removeJWTToken();
      emit(const InitStateUnauthenticated());
    }
  }
}

sealed class InitEvent {
  const InitEvent();
}

class InitEventCheck extends InitEvent {
  const InitEventCheck();
}

sealed class InitState {
  const InitState();
}

class InitStateInitial extends InitState {
  const InitStateInitial();
}

class InitStateAuthenticated extends InitState {
  const InitStateAuthenticated();
}

class InitStateUnauthenticated extends InitState {
  const InitStateUnauthenticated();
}
