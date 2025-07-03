import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:mocktail/mocktail.dart';
import 'package:bloc_test/bloc_test.dart';
import 'package:android_app/features/path/presentation/screens/path_page.dart';
import 'package:android_app/features/path/domain/bloc/workout_path_bloc.dart';
import 'package:android_app/features/path/domain/entities/workout_entity.dart';

class MockWorkoutBloc extends Mock implements WorkoutBloc {}

void main() {
  setUpAll(() {
    registerFallbackValue(WorkoutStateInitial());
    registerFallbackValue(WorkoutEventFetch());
  });

  Widget makeTestable(Widget child, WorkoutBloc bloc) =>
      MaterialApp(home: BlocProvider.value(value: bloc, child: child));

  testWidgets('shows loading indicator when state is WorkoutStateLoading', (tester) async {
    final bloc = MockWorkoutBloc();
    when(() => bloc.state).thenReturn(WorkoutStateLoading());
    whenListen(bloc, Stream<WorkoutState>.empty(), initialState: WorkoutStateLoading());
    await tester.pumpWidget(makeTestable(const PathPage(), bloc));
    expect(find.byType(CircularProgressIndicator), findsOneWidget);
  });

  testWidgets('shows error message and refresh button when state is WorkoutStateError', (tester) async {
    final bloc = MockWorkoutBloc();
    when(() => bloc.state).thenReturn(WorkoutStateError());
    whenListen(bloc, Stream<WorkoutState>.empty(), initialState: WorkoutStateError());

    await tester.pumpWidget(makeTestable(const PathPage(), bloc));
    expect(find.text('Something went wrong...'), findsOneWidget);
    expect(find.text('Refresh'), findsOneWidget);
  });

  testWidgets('shows WorkoutPath when state is WorkoutStateLoaded with workouts', (tester) async {
    final bloc = MockWorkoutBloc();
    final workouts = [
      WorkoutEntity(
        workoutId: 'w1',
        name: 'Workout 1',
        description: 'desc',
        status: 'planned',
        exercises: [],
      ),
    ];
    when(() => bloc.state).thenReturn(WorkoutStateLoaded(workout: workouts));
    whenListen(bloc, Stream<WorkoutState>.empty(), initialState: WorkoutStateLoaded(workout: workouts));

    await tester.pumpWidget(makeTestable(const PathPage(), bloc));
    expect(find.byType(WorkoutPath), findsOneWidget);
    expect(find.text('Workout 1'), findsNothing); // Name is not directly shown
  });

  testWidgets('shows generate message and button when state is initial/empty', (tester) async {
    final bloc = MockWorkoutBloc();
    when(() => bloc.state).thenReturn(WorkoutStateInitial());
    whenListen(bloc, Stream<WorkoutState>.empty(), initialState: WorkoutStateInitial());

    await tester.pumpWidget(makeTestable(const PathPage(), bloc));
    expect(find.text("There's nothing here. Generate your workout"), findsOneWidget);
    expect(find.text('Generate'), findsOneWidget);
  });
}
