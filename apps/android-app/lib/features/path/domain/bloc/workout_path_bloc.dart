import 'package:android_app/features/path/domain/entities/WorkoutEntity.dart';
import 'package:android_app/features/path/domain/repositories/workout_path_repository.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class WorkoutBloc extends Bloc<WorkoutEvent, WorkoutState> {
  WorkoutBloc({required WorkoutPathRepository workoutPathRepository})
    : _workoutPathRepository = workoutPathRepository,
      super(WorkoutStateInitial()) {
    on<WorkoutEvent>(
      (event, emit) => switch (event) {
        WorkoutEventFetch() => _onWorkoutLoad(event, emit),
        WorkoutEventGenerate() => _onWorkoutGenerate(event, emit),
      },
    );
  }

  Future<void> _onWorkoutLoad(
    WorkoutEventFetch event,
    Emitter<WorkoutState> emit,
  ) async {
    emit(WorkoutStateLoading());
    try {
      final workouts = await _workoutPathRepository.fetchWorkouts();
      emit(WorkoutStateLoaded(workout: workouts));
    } on Object catch (e, s) {
      if (kDebugMode) print('$e, $s');
      emit(WorkoutStateError());
    }
  }

  Future<void> _onWorkoutGenerate(
    WorkoutEventGenerate event,
    Emitter<WorkoutState> emit,
  ) async {
    emit(WorkoutStateLoading());
    try {
      if (await _workoutPathRepository.generate()) {
        final workouts = await _workoutPathRepository.fetchWorkouts();
        emit(WorkoutStateLoaded(workout: workouts));
      } else {
        emit(WorkoutStateError());
      }
    } on Object catch (e, s) {
      if (kDebugMode) print('$e, $s');
      emit(WorkoutStateError());
    }
  }

  final WorkoutPathRepository _workoutPathRepository;
}

sealed class WorkoutEvent {
  const WorkoutEvent();
}

class WorkoutEventFetch extends WorkoutEvent {
  const WorkoutEventFetch();
}

class WorkoutEventGenerate extends WorkoutEvent {
  const WorkoutEventGenerate();
}

sealed class WorkoutState {
  const WorkoutState();
}

class WorkoutStateInitial extends WorkoutState {
  const WorkoutStateInitial();
}

class WorkoutStateLoading extends WorkoutState {
  const WorkoutStateLoading();
}

class WorkoutStateLoaded extends WorkoutState {
  const WorkoutStateLoaded({required this.workout});

  final List<WorkoutEntity> workout;
}

class WorkoutStateError extends WorkoutState {
  const WorkoutStateError();
}
