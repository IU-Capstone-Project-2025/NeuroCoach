import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'package:android_app/features/path/domain/bloc/workout_path_bloc.dart';
import 'package:android_app/features/path/domain/entities/workout_entity.dart';
import 'package:android_app/features/path/domain/repositories/workout_path_repository.dart';

import 'workout_path_bloc_test.mocks.dart';

@GenerateMocks([
  WorkoutPathRepository,
])
void main() {
  group('WorkoutBloc', () {
    late MockWorkoutPathRepository mockWorkoutPathRepository;
    late WorkoutBloc bloc;

    setUp(() {
      mockWorkoutPathRepository = MockWorkoutPathRepository();
      bloc = WorkoutBloc(workoutPathRepository: mockWorkoutPathRepository);
    });

    final workouts = [
      WorkoutEntity(
        workoutId: 'w1',
        name: 'Workout 1',
        description: 'Description 1',
        status: 'active',
        exercises: [
          Exercise(
            exerciseId: 'e1',
            name: 'Push Up',
            muscleGroup: 'Chest',
            sets: 3,
            reps: 12,
            restSec: 60,
            notes: 'Keep back straight',
            technique: 'Standard',
          ),
        ],
      ),
      WorkoutEntity(
        workoutId: 'w2',
        name: 'Workout 2',
        description: 'Description 2',
        status: 'planned',
        exercises: [
          Exercise(
            exerciseId: 'e2',
            name: 'Squat',
            muscleGroup: 'Legs',
            sets: 4,
            reps: 10,
            restSec: 90,
            notes: 'Go low',
            technique: 'Standard',
          ),
        ],
      ),
    ];

    test('emits [Loading, Loaded] on fetch success', () async {
      when(mockWorkoutPathRepository.fetchWorkouts()).thenAnswer((_) async => workouts);

      expectLater(
        bloc.stream,
        emitsInOrder([
          isA<WorkoutStateLoading>(),
          isA<WorkoutStateLoaded>().having((s) => s.workout, 'workout', workouts),
        ]),
      );

      bloc.add(const WorkoutEventFetch());
    });

    test('emits [Loading, Error] on fetch failure', () async {
      when(mockWorkoutPathRepository.fetchWorkouts()).thenThrow(Exception('fetch error'));

      expectLater(
        bloc.stream,
        emitsInOrder([
          isA<WorkoutStateLoading>(),
          isA<WorkoutStateError>(),
        ]),
      );

      bloc.add(const WorkoutEventFetch());
    });

    test('emits [Loading, Loaded] on generate success', () async {
      when(mockWorkoutPathRepository.generate()).thenAnswer((_) async => true);
      when(mockWorkoutPathRepository.fetchWorkouts()).thenAnswer((_) async => workouts);

      expectLater(
        bloc.stream,
        emitsInOrder([
          isA<WorkoutStateLoading>(),
          isA<WorkoutStateLoaded>().having((s) => s.workout, 'workout', workouts),
        ]),
      );

      bloc.add(const WorkoutEventGenerate());
    });

    test('emits [Loading, Error] on generate returns false', () async {
      when(mockWorkoutPathRepository.generate()).thenAnswer((_) async => false);

      expectLater(
        bloc.stream,
        emitsInOrder([
          isA<WorkoutStateLoading>(),
          isA<WorkoutStateError>(),
        ]),
      );

      bloc.add(const WorkoutEventGenerate());
    });

    test('emits [Loading, Error] on generate throws', () async {
      when(mockWorkoutPathRepository.generate()).thenThrow(Exception('generate error'));

      expectLater(
        bloc.stream,
        emitsInOrder([
          isA<WorkoutStateLoading>(),
          isA<WorkoutStateError>(),
        ]),
      );

      bloc.add(const WorkoutEventGenerate());
    });
  });
}
