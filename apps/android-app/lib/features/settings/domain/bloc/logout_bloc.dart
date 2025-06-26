import 'package:android_app/features/settings/domain/repositories/logout_repository.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class LogoutBloc extends Bloc<LogoutEvent, LogoutState> {
  LogoutBloc({required LogoutRepository logoutRepository})
    : _logoutRepository = logoutRepository,
      super(LogoutStateInitial()) {
    on<LogoutEvent>(
        (event, emit) => switch (event) {
          LogoutEventLogout() => _onLogout(event, emit),
        }
    );
  }
  
  Future<void> _onLogout(LogoutEventLogout event, Emitter<LogoutState> emit) async {
    await _logoutRepository.logout() ? emit(LogoutStateLoggedOut()) : emit(LogoutStateError());
  }

  final LogoutRepository _logoutRepository;
}

sealed class LogoutEvent {
  const LogoutEvent();
}

class LogoutEventLogout extends LogoutEvent {
  const LogoutEventLogout();
}

sealed class LogoutState {
  const LogoutState();
}

class LogoutStateInitial extends LogoutState {
  const LogoutStateInitial();
}

class LogoutStateLoggedOut extends LogoutState {
  const LogoutStateLoggedOut();
}

class LogoutStateError extends LogoutState {
  const LogoutStateError();
}
