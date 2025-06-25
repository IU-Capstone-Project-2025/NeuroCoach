import 'package:android_app/features/init/domain/repositories/health_repository.dart';
import 'package:android_app/features/init/domain/repositories/init_repository.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class InitBloc extends Bloc<InitEvent, InitState> {

  InitBloc(
      {required InitRepository initRepository, required HealthRepository healthRepository})
      : _initRepository = initRepository,
        _healthRepository = healthRepository,
        super(const InitStateInitial()) {
    on<InitEvent>(
        (event, emit) => switch (event) {
          InitEventCheck() => _onCheck(event, emit),
        }
    );
  }

  final InitRepository _initRepository;
  final HealthRepository _healthRepository;

  Future<void> _onCheck(InitEventCheck event, Emitter<InitState> emit) async {
    final String token = await _initRepository.getJWTToken();
    if (kDebugMode) print('Got JWT token: $token');
    if (token.isEmpty) {
      emit(const InitStateUnauthenticated());
      return;
    }
    await _healthRepository.checkToken(token) ? emit(const InitStateAuthenticated()) : emit(const InitStateUnauthenticated());
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