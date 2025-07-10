import 'package:android_app/app/app_router.dart';
import 'package:android_app/features/path/domain/bloc/finish_workout_bloc.dart';
import 'package:android_app/features/path/domain/bloc/workout_path_bloc.dart';
import 'package:android_app/features/settings/domain/bloc/logout_bloc.dart';
import 'package:android_app/features/settings/domain/bloc/profile_bloc.dart';
import 'package:auto_route/auto_route.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../app/presentation/scopes/dependencies_scope.dart';
import '../../chat/domain/bloc/chat_bloc.dart';

class HomeScope extends StatelessWidget {
  const HomeScope({super.key, required this.child});

  final Widget child;

  @override
  Widget build(BuildContext context) {
    final depScope = DependenciesScope.findAppDependenciesOf(context);
    return MultiBlocProvider(
      providers: [
        BlocProvider(
          create: (_) =>
              ChatBloc(chatMessagesRepository: depScope.chatMessagesRepository)
                ..add(ChatEventRetrieve()),
        ),
        BlocProvider(
          create: (_) =>
              LogoutBloc(logoutRepository: depScope.logoutRepository),
        ),
        BlocProvider(
          create: (_) =>
              WorkoutBloc(workoutPathRepository: depScope.workoutPathRepository)
                ..add(WorkoutEventFetch()),
        ),
        BlocProvider(
          create: (_) =>
              ProfileBloc(profileRepository: depScope.profileRepository)
                ..add(ProfileEventLoad()),
        ),
        BlocProvider(
          create: (_) => FinishWorkoutBloc(depScope.finishWorkoutRepository),
        ),
      ],
      child: MultiBlocListener(
        listeners: [
          BlocListener<LogoutBloc, LogoutState>(
            listener: (BuildContext context, LogoutState state) {
              if (state is LogoutStateLoggedOut) {
                context.router.replace(InitRoute());
              } else if (state is LogoutStateError) {
                ScaffoldMessenger.of(
                  context,
                ).showSnackBar(SnackBar(content: Text('Something went wrong')));
              }
            },
          ),
          BlocListener<FinishWorkoutBloc, FinishWorkoutState>(
            listener: (BuildContext context, FinishWorkoutState state) {
              if (state is FinishWorkoutStateLoaded) {
                context.read<WorkoutBloc>().add(WorkoutEventFetch());
              }
              if (state is FinishWorkoutStateError) {
                ScaffoldMessenger.of(
                  context,
                ).showSnackBar(SnackBar(content: Text('Something went wrong')));
              }
            },
          ),
        ],
        child: child,
      ),
    );
  }
}
