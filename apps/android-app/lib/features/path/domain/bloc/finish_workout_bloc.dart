import 'package:android_app/features/path/domain/repositories/finish_workout_repository.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class FinishWorkoutBloc extends Bloc<FinishWorkoutEvent, FinishWorkoutState> {
  FinishWorkoutBloc(FinishWorkoutRepository finishWorkoutRepository)
    : _finishWorkoutRepository = finishWorkoutRepository,
      super(FinishWorkoutStateInitial()) {
    on<FinishWorkoutEvent>(
      (event, emit) => switch (event) {
        FinishWorkoutEventFinish() => _onWorkoutFinish(event, emit),
      },
    );
  }

  void _onWorkoutFinish(
    FinishWorkoutEventFinish event,
    Emitter<FinishWorkoutState> emit,
  ) async {
    try {
      await _finishWorkoutRepository.finishWorkout(event.workoutId);
      emit(const FinishWorkoutStateLoaded());
    } on Exception catch (e, s) {
      if (kDebugMode) print('$e, $s');
      emit(FinishWorkoutStateError(e));
    }
  }

  final FinishWorkoutRepository _finishWorkoutRepository;
}

sealed class FinishWorkoutEvent {
  const FinishWorkoutEvent();
}

class FinishWorkoutEventFinish extends FinishWorkoutEvent {
  const FinishWorkoutEventFinish(this.workoutId);

  final String workoutId;
}

sealed class FinishWorkoutState {
  const FinishWorkoutState();
}

class FinishWorkoutStateInitial extends FinishWorkoutState {
  const FinishWorkoutStateInitial();
}

class FinishWorkoutStateLoaded extends FinishWorkoutState {
  const FinishWorkoutStateLoaded();
}

class FinishWorkoutStateError extends FinishWorkoutState {
  const FinishWorkoutStateError(this.err);

  final Exception err;
}
