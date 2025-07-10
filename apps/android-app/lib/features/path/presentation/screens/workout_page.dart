import 'package:android_app/constants/app_text_styles.dart';
import 'package:android_app/features/path/domain/bloc/finish_workout_bloc.dart';
import 'package:android_app/features/path/domain/entities/workout_entity.dart';
import 'package:android_app/features/path/presentation/components/rest_timer.dart';
import 'package:android_app/uikit/buttons/app_button.dart';
import 'package:auto_route/auto_route.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../components/active_exercise_view.dart';
import '../components/exercise_overview.dart';

@RoutePage()
class WorkoutPage extends StatefulWidget {
  const WorkoutPage({
    super.key,
    required this.workoutId,
    required this.name,
    required this.exercises,
    required this.isCurrentTraining,
  });

  final String workoutId;
  final String name;
  final bool isCurrentTraining;
  final List<Exercise> exercises;

  @override
  State<WorkoutPage> createState() => _WorkoutPageState();
}

class _WorkoutPageState extends State<WorkoutPage> {
  int? _expandedIndex;
  bool isTrainingStarted = false;
  bool isTrainingFinished = false;
  bool isRest = false;
  int exerciseIndex = 0;

  final Map<int, int> performedSets = {}; // index -> completed sets

  void _startTraining() {
    setState(() {
      isTrainingStarted = true;
      exerciseIndex = 0;
      performedSets.clear();
    });
  }

  void _onRestFinished() {
    final setsCompleted = performedSets[exerciseIndex] ?? 0;
    performedSets[exerciseIndex] = setsCompleted + 1;

    final nextIndex = _chooseNextValidExercise();

    if (nextIndex != -1) {
      setState(() {
        exerciseIndex = nextIndex;
        isRest = false;
      });
    } else {
      setState(() {
        isTrainingFinished = true;
      });
    }
  }

  int _chooseNextValidExercise() {
    final total = widget.exercises.length;
    for (int offset = 1; offset <= total; offset++) {
      final nextIndex = (exerciseIndex + offset) % total;
      final setsDone = performedSets[nextIndex] ?? 0;
      if (setsDone < widget.exercises[nextIndex].sets) {
        return nextIndex;
      }
    }
    return -1;
  }

  void _goToRest() {
    if (_chooseNextValidExercise() == -1) {
      _onRestFinished();
      return;
    }
    setState(() {
      isRest = true;
    });
  }

  int _chooseRestTime() {
    final currentExercise = widget.exercises[exerciseIndex];
    return exerciseIndex == widget.exercises.length - 1
        ? 300
        : currentExercise.restSec;
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.name),
        titleTextStyle: AppTextStyles.chatTitle,
        centerTitle: true,
        automaticallyImplyLeading: !isTrainingStarted,
      ),
      body: Center(
        child: isTrainingFinished
            ? Center(
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Text(
                      'You completed all the exercises!',
                      style: AppTextStyles.textButton,
                    ),
                    const SizedBox(height: 16),
                    AppButton(
                      onPressed: () => {
                        context.read<FinishWorkoutBloc>().add(
                          FinishWorkoutEventFinish(widget.workoutId),
                        ),
                        context.router.maybePop(),
                      },
                      text: 'Finish',
                    ),
                  ],
                ),
              )
            : isTrainingStarted
            ? isRest
                  ? RestTimer(
                      restTime: _chooseRestTime(),
                      onFinished: _onRestFinished,
                    )
                  : ActiveExerciseView(
                      exercise: widget.exercises[exerciseIndex],
                      performedSets: performedSets[exerciseIndex] ?? 0,
                      onDoneWithSet: _goToRest,
                    )
            : ExerciseOverview(
                exercises: widget.exercises,
                performedSets: performedSets,
                expandedIndex: _expandedIndex,
                onExpansionChanged: (index) {
                  setState(() {
                    _expandedIndex = index;
                  });
                },
                isTrainingStarted: isTrainingStarted,
                onStartTraining: _startTraining,
                isCurrentTraining: widget.isCurrentTraining,
              ),
      ),
    );
  }
}
