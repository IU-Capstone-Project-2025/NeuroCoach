import 'package:android_app/features/home/presentation/home_page.dart';
import 'package:android_app/features/init/presentation/init_page.dart';
import 'package:flutter/material.dart';
import 'package:auto_route/auto_route.dart';
import '../features/chat/presentation/screens/chat_page.dart';
import '../features/settings/presentation/screens/settings_page.dart';
import '../features/settings/presentation/screens/settings_main_page.dart';
import '../features/path/presentation/screens/path_page.dart';
import '../features/login/presentation/auth_page.dart';
import '../features/settings/presentation/screens/profile_page.dart';
import '../features/path/presentation/screens/workouts_page.dart';
import '../features/path/presentation/screens/workout_page.dart';
import '../features/path/domain/entities/WorkoutEntity.dart';


part 'app_router.gr.dart';

@AutoRouterConfig()
class AppRouter extends _$AppRouter {

  @override
  RouteType get defaultRouteType => RouteType.custom(
    durationInMilliseconds: 300,
    reverseDurationInMilliseconds: 300,
    transitionsBuilder: (context, animation, secondaryAnimation, child) {
      return FadeTransition(opacity: animation, child: child);
    },
  );

  @override
  List<AutoRoute> get routes => [
    AutoRoute(
      page: InitRoute.page,
      initial: true,
      path: '/',
      type: const RouteType.custom(
        durationInMilliseconds: 0,
        reverseDurationInMilliseconds: 0,
        transitionsBuilder: TransitionsBuilders.noTransition,
      ),
    ),
    AutoRoute(page: AuthRoute.page, path: '/login'),
    AutoRoute(
      page: HomeRoute.page,
      path: '/home',
      children: [
        AutoRoute(page: ChatRoute.page, path: 'chat'),
        AutoRoute(page: WorkoutsRoute.page, path: 'workouts', initial: true, children: [
          AutoRoute(page: WorkoutRoute.page, path: 'workout'),
          AutoRoute(page: PathRoute.page, path: '')
        ]),
        AutoRoute(
          page: SettingsRoute.page,
          path: 'settings',
          children: [
            AutoRoute(page: SettingsMainRoute.page, path: ''), // default
            AutoRoute(page: ProfileRoute.page, path: 'profile'),
          ],
        ),
      ],
    ),
  ];
}
