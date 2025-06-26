import 'package:android_app/constants/app_colors.dart';
import 'package:android_app/constants/app_text_styles.dart';
import 'package:android_app/features/path/domain/bloc/workout_path_bloc.dart';
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
      ),
      body: BlocBuilder<WorkoutBloc, WorkoutState>(
        builder: (BuildContext context, WorkoutState state) {
          if (state is WorkoutStateLoading) {
            return CircularProgressIndicator();
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
  final List<Map<String, dynamic>> workouts;

  const WorkoutPath({super.key, required this.workouts});

  WorkoutStatus _getStatus(int index) {
    final current = workouts[index];
    final status = current['status'] as String;

    if (status == 'complete') return WorkoutStatus.completed;

    final prevCompleted = workouts
        .sublist(0, index)
        .where((w) => w['status'] == 'complete')
        .isNotEmpty;

    final hasPreviousPlanned = workouts
        .sublist(0, index)
        .any((w) => w['status'] == 'planned');

    if (!hasPreviousPlanned && prevCompleted) return WorkoutStatus.current;

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
        final isTappable = status == WorkoutStatus.current || status == WorkoutStatus.completed;

        Widget node = Container(
          width: 64,
          height: 64,
          decoration: BoxDecoration(
            color: _getColor(status),
            borderRadius: BorderRadius.circular(16),
          ),
          child: Center(
            child: Icon(
              status == WorkoutStatus.completed ? Icons.check : Icons.fitness_center,
              color: AppColors.black,
            ),
          ),
        );

        if (isTappable) {
          node = GestureDetector(
            onTap: () {
              context.pushRoute(
                WorkoutRoute(
                  name: workout['name'],
                  exercises: List<Map<String, dynamic>>.from(workout['exercises']),
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
                  children: [
                    Padding(
                      padding: _getPadding(index),
                      child: node
                    ),
                  ],
                ),
              ),
            ],
          ),
        );
      },
    );
  }
}

