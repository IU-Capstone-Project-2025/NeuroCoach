import 'package:android_app/constants/app_text_styles.dart';
import 'package:android_app/features/path/domain/bloc/workout_path_bloc.dart';
import 'package:android_app/features/path/domain/entities/workout_entity.dart';
import 'package:auto_route/auto_route.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../app/app_router.dart';
import '../components/path_node.dart';

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
              GlossaryRoute(),
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

  NodeStatus _getStatus(int index) {
    final current = workouts[index];
    final status = current.status;

    if (status == 'done') return NodeStatus.done;

    final prevCompleted = workouts
        .sublist(0, index)
        .any((w) => w.status == 'done');

    final hasPreviousPlanned = workouts
        .sublist(0, index)
        .any((w) => w.status == 'planned' || w.status == 'expired');

    if ((!hasPreviousPlanned && prevCompleted) ||
        (index == 0 &&
            (status == 'planned' || status == 'expired'))) {
      return NodeStatus.current;
    }

    return NodeStatus.planned;
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
        final isTappable = status == NodeStatus.current || status == NodeStatus.done;

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
                      child: PathNode(
                        status: status,
                        index: index + 1,
                        onTap: isTappable
                            ? () {
                          context.pushRoute(
                            WorkoutRoute(
                              name: workout.name,
                              exercises: workout.exercises,
                              isCurrentTraining: workout.status == 'planned' || workout.status == 'expired',
                              workoutId: workout.workoutId,
                            ),
                          );
                        }
                            : () {}, // fallback for untappable
                      ),
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

