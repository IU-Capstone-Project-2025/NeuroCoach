import 'package:android_app/features/home/presentation/home_page.dart';
import 'package:android_app/features/init/presentation/init_page.dart';
import 'package:flutter/material.dart';
import 'package:auto_route/auto_route.dart';
import '../features/chat/presentation/screens/chat_page.dart';
import '../features/settings/presentation/screens/settings_page.dart';
import '../features/path/presentation/screens/path_page.dart';
import '../features/login/presentation/auth_page.dart';

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
        AutoRoute(page: PathRoute.page, path: 'path', initial: true),
        AutoRoute(page: SettingsRoute.page, path: 'settings'),
      ],
    ),
  ];
}
