import 'package:android_app/constants/app_colors.dart';
import 'package:android_app/constants/app_text_styles.dart';
import 'package:android_app/features/path/domain/bloc/workout_path_bloc.dart';
import 'package:android_app/features/path/domain/entities/workout_entity.dart';
import 'package:auto_route/auto_route.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../app/app_router.dart';

enum WorkoutStatus { completed, current, planned }

@RoutePage()
class PathPage extends StatelessWidget {
  const PathPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        automaticallyImplyLeading: false,
        title: Text('Training'),
        titleTextStyle: AppTextStyles.chatTitle,
        centerTitle: true,
        actions: [
          IconButton(
            onPressed: () => context.pushRoute(
              GlossaryRoute(
                exercises: [
                  Exercise(
                    exerciseId: 'ex001',
                    name: 'Push-Up',
                    muscleGroup: 'Chest',
                    sets: 3,
                    reps: 15,
                    restSec: 60,
                    notes: 'Keep your core tight and back straight.',
                    technique:
                        'Lower your body until your chest nearly touches the floor, then push back up.',
                    pictureUrl: 'https://i.pinimg.com/736x/47/72/90/477290deb87fcba55a99c10c41186240.jpg',
                  ),
                  Exercise(
                    exerciseId: 'ex002',
                    name: 'Squat',
                    muscleGroup: 'Legs',
                    sets: 4,
                    reps: 12,
                    restSec: 90,
                    notes:
                        'Keep knees in line with toes. Don’t let heels lift.',
                    technique:
                        'Lower your hips back and down as if sitting in a chair, then drive through your heels to stand.',
                    pictureUrl: 'https://www.shutterstock.com/image-illustration/bodyweight-squat-thighs-exercise-male-600w-2329917681.jpg',
                  ),
                  Exercise(
                    exerciseId: 'ex003',
                    name: 'Plank',
                    muscleGroup: 'Core',
                    sets: 3,
                    reps: 1,
                    // Each rep is a hold
                    restSec: 45,
                    notes: 'Maintain a straight line from head to heels.',
                    technique:
                        'Hold the plank position with elbows under shoulders and core engaged for 30–60 seconds.',
                    pictureUrl: 'https://blog.trainerlist.com/wp-content/uploads/2024/07/plankkk.jpg',
                  ),
                ],
              ),
            ),
            icon: Icon(Icons.book),
          ),
        ],
      ),
      body: BlocBuilder<WorkoutBloc, WorkoutState>(
        builder: (BuildContext context, WorkoutState state) {
          if (state is WorkoutStateLoading || state is WorkoutStateRefresh) {
            if (state is WorkoutStateRefresh) {
              context.read<WorkoutBloc>().add(WorkoutEventGenerate());
            }
            return Center(child: CircularProgressIndicator());
          } else if (state is WorkoutStateError) {
            return Center(
              child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Text(
                    'Something went wrong...',
                    style: AppTextStyles.textButton,
                  ),
                  TextButton(
                    onPressed: () =>
                        context.read<WorkoutBloc>().add(WorkoutEventFetch()),
                    child: Text('Refresh'),
                  ),
                ],
              ),
            );
          } else if (state is WorkoutStateLoaded && state.workout.isNotEmpty) {
            return WorkoutPath(workouts: state.workout);
          } else {
            return Center(
              child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Text(
                    "There's nothing here. Generate your workout",
                    style: AppTextStyles.textButton,
                  ),
                  TextButton(
                    onPressed: () =>
                        context.read<WorkoutBloc>().add(WorkoutEventGenerate()),
                    child: Text('Generate'),
                  ),
                ],
              ),
            );
          }
        },
      ),
    );
  }
}

class WorkoutPath extends StatelessWidget {
  final List<WorkoutEntity> workouts;

  const WorkoutPath({super.key, required this.workouts});

  WorkoutStatus _getStatus(int index) {
    final current = workouts[index];
    final status = current.status;

    if (status == 'done') return WorkoutStatus.completed;

    final prevCompleted = workouts
        .sublist(0, index)
        .where((w) => w.status == 'done')
        .isNotEmpty;

    final hasPreviousPlanned = workouts
        .sublist(0, index)
        .any((w) => w.status == 'planned' || w.status == 'expired');

    if ((!hasPreviousPlanned && prevCompleted) ||
        (index == 0 &&
            (workouts[index].status == 'planned' ||
                workouts[index].status == 'expired'))) {
      return WorkoutStatus.current;
    }

    return WorkoutStatus.planned;
  }

  Color _getColor(WorkoutStatus status) {
    switch (status) {
      case WorkoutStatus.completed:
        return AppColors.lily;
      case WorkoutStatus.current:
        return AppColors.white;
      case WorkoutStatus.planned:
        return AppColors.blackNode;
    }
  }

  CrossAxisAlignment _getAlignment(int index) {
    return index % 2 == 0 ? CrossAxisAlignment.start : CrossAxisAlignment.end;
  }

  EdgeInsets _getPadding(int index) {
    return index % 2 == 0
        ? const EdgeInsets.only(left: 96)
        : const EdgeInsets.only(right: 96);
  }

  @override
  Widget build(BuildContext context) {
    return ListView.builder(
      padding: const EdgeInsets.symmetric(vertical: 24),
      itemCount: workouts.length,
      itemBuilder: (context, index) {
        final workout = workouts[index];
        final status = _getStatus(index);
        final isTappable =
            status == WorkoutStatus.current ||
            status == WorkoutStatus.completed;

        Widget node = Container(
          width: 64,
          height: 64,
          decoration: BoxDecoration(
            color: _getColor(status),
            borderRadius: BorderRadius.circular(16),
          ),
          child: Center(
            child: Icon(
              status == WorkoutStatus.completed
                  ? Icons.check
                  : Icons.fitness_center,
              color: AppColors.black,
            ),
          ),
        );

        if (isTappable) {
          node = GestureDetector(
            onTap: () {
              context.pushRoute(
                WorkoutRoute(
                  name: workout.name,
                  exercises: workout.exercises,
                  isCurrentTraining:
                      workout.status == 'planned' ||
                      workout.status == 'expired',
                  workoutId: workout.workoutId,
                ),
              );
            },
            child: node,
          );
        }

        return Padding(
          padding: const EdgeInsets.symmetric(vertical: 12),
          child: Row(
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              Expanded(
                child: Column(
                  crossAxisAlignment: _getAlignment(index),
                  children: [Padding(padding: _getPadding(index), child: node)],
                ),
              ),
            ],
          ),
        );
      },
    );
  }
}
