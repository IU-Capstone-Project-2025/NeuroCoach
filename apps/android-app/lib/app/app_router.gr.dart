// GENERATED CODE - DO NOT MODIFY BY HAND

// **************************************************************************
// AutoRouterGenerator
// **************************************************************************

// ignore_for_file: type=lint
// coverage:ignore-file

part of 'app_router.dart';

abstract class _$AppRouter extends RootStackRouter {
  // ignore: unused_element
  _$AppRouter({super.navigatorKey});

  @override
  final Map<String, PageFactory> pagesMap = {
    AuthRoute.name: (routeData) {
      final args =
          routeData.argsAs<AuthRouteArgs>(orElse: () => const AuthRouteArgs());
      return AutoRoutePage<dynamic>(
        routeData: routeData,
        child: AuthPage(
          key: args.key,
          isSignUp: args.isSignUp,
        ),
      );
    },
    ChatRoute.name: (routeData) {
      return AutoRoutePage<dynamic>(
        routeData: routeData,
        child: const ChatPage(),
      );
    },
    HomeRoute.name: (routeData) {
      return AutoRoutePage<dynamic>(
        routeData: routeData,
        child: const HomePage(),
      );
    },
    InitRoute.name: (routeData) {
      return AutoRoutePage<dynamic>(
        routeData: routeData,
        child: const InitPage(),
      );
    },
    PathRoute.name: (routeData) {
      return AutoRoutePage<dynamic>(
        routeData: routeData,
        child: const PathPage(),
      );
    },
    ProfileRoute.name: (routeData) {
      return AutoRoutePage<dynamic>(
        routeData: routeData,
        child: const ProfilePage(),
      );
    },
    SettingsMainRoute.name: (routeData) {
      return AutoRoutePage<dynamic>(
        routeData: routeData,
        child: const SettingsMainPage(),
      );
    },
    SettingsRoute.name: (routeData) {
      return AutoRoutePage<dynamic>(
        routeData: routeData,
        child: const SettingsPage(),
      );
    },
    WorkoutRoute.name: (routeData) {
      final args = routeData.argsAs<WorkoutRouteArgs>();
      return AutoRoutePage<dynamic>(
        routeData: routeData,
        child: WorkoutPage(
          key: args.key,
          workoutId: args.workoutId,
          name: args.name,
          exercises: args.exercises,
          isCurrentTraining: args.isCurrentTraining,
        ),
      );
    },
    WorkoutsRoute.name: (routeData) {
      return AutoRoutePage<dynamic>(
        routeData: routeData,
        child: const WorkoutsPage(),
      );
    },
  };
}

/// generated route for
/// [AuthPage]
class AuthRoute extends PageRouteInfo<AuthRouteArgs> {
  AuthRoute({
    Key? key,
    bool isSignUp = false,
    List<PageRouteInfo>? children,
  }) : super(
          AuthRoute.name,
          args: AuthRouteArgs(
            key: key,
            isSignUp: isSignUp,
          ),
          initialChildren: children,
        );

  static const String name = 'AuthRoute';

  static const PageInfo<AuthRouteArgs> page = PageInfo<AuthRouteArgs>(name);
}

class AuthRouteArgs {
  const AuthRouteArgs({
    this.key,
    this.isSignUp = false,
  });

  final Key? key;

  final bool isSignUp;

  @override
  String toString() {
    return 'AuthRouteArgs{key: $key, isSignUp: $isSignUp}';
  }
}

/// generated route for
/// [ChatPage]
class ChatRoute extends PageRouteInfo<void> {
  const ChatRoute({List<PageRouteInfo>? children})
      : super(
          ChatRoute.name,
          initialChildren: children,
        );

  static const String name = 'ChatRoute';

  static const PageInfo<void> page = PageInfo<void>(name);
}

/// generated route for
/// [HomePage]
class HomeRoute extends PageRouteInfo<void> {
  const HomeRoute({List<PageRouteInfo>? children})
      : super(
          HomeRoute.name,
          initialChildren: children,
        );

  static const String name = 'HomeRoute';

  static const PageInfo<void> page = PageInfo<void>(name);
}

/// generated route for
/// [InitPage]
class InitRoute extends PageRouteInfo<void> {
  const InitRoute({List<PageRouteInfo>? children})
      : super(
          InitRoute.name,
          initialChildren: children,
        );

  static const String name = 'InitRoute';

  static const PageInfo<void> page = PageInfo<void>(name);
}

/// generated route for
/// [PathPage]
class PathRoute extends PageRouteInfo<void> {
  const PathRoute({List<PageRouteInfo>? children})
      : super(
          PathRoute.name,
          initialChildren: children,
        );

  static const String name = 'PathRoute';

  static const PageInfo<void> page = PageInfo<void>(name);
}

/// generated route for
/// [ProfilePage]
class ProfileRoute extends PageRouteInfo<void> {
  const ProfileRoute({List<PageRouteInfo>? children})
      : super(
          ProfileRoute.name,
          initialChildren: children,
        );

  static const String name = 'ProfileRoute';

  static const PageInfo<void> page = PageInfo<void>(name);
}

/// generated route for
/// [SettingsMainPage]
class SettingsMainRoute extends PageRouteInfo<void> {
  const SettingsMainRoute({List<PageRouteInfo>? children})
      : super(
          SettingsMainRoute.name,
          initialChildren: children,
        );

  static const String name = 'SettingsMainRoute';

  static const PageInfo<void> page = PageInfo<void>(name);
}

/// generated route for
/// [SettingsPage]
class SettingsRoute extends PageRouteInfo<void> {
  const SettingsRoute({List<PageRouteInfo>? children})
      : super(
          SettingsRoute.name,
          initialChildren: children,
        );

  static const String name = 'SettingsRoute';

  static const PageInfo<void> page = PageInfo<void>(name);
}

/// generated route for
/// [WorkoutPage]
class WorkoutRoute extends PageRouteInfo<WorkoutRouteArgs> {
  WorkoutRoute({
    Key? key,
    required String workoutId,
    required String name,
    required List<Exercise> exercises,
    required bool isCurrentTraining,
    List<PageRouteInfo>? children,
  }) : super(
          WorkoutRoute.name,
          args: WorkoutRouteArgs(
            key: key,
            workoutId: workoutId,
            name: name,
            exercises: exercises,
            isCurrentTraining: isCurrentTraining,
          ),
          initialChildren: children,
        );

  static const String name = 'WorkoutRoute';

  static const PageInfo<WorkoutRouteArgs> page =
      PageInfo<WorkoutRouteArgs>(name);
}

class WorkoutRouteArgs {
  const WorkoutRouteArgs({
    this.key,
    required this.workoutId,
    required this.name,
    required this.exercises,
    required this.isCurrentTraining,
  });

  final Key? key;

  final String workoutId;

  final String name;

  final List<Exercise> exercises;

  final bool isCurrentTraining;

  @override
  String toString() {
    return 'WorkoutRouteArgs{key: $key, workoutId: $workoutId, name: $name, exercises: $exercises, isCurrentTraining: $isCurrentTraining}';
  }
}

/// generated route for
/// [WorkoutsPage]
class WorkoutsRoute extends PageRouteInfo<void> {
  const WorkoutsRoute({List<PageRouteInfo>? children})
      : super(
          WorkoutsRoute.name,
          initialChildren: children,
        );

  static const String name = 'WorkoutsRoute';

  static const PageInfo<void> page = PageInfo<void>(name);
}
