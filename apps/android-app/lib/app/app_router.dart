import 'package:android_app/features/home/presentation/home_page.dart';
import 'package:android_app/features/init/presentation/init_route.dart';
import 'package:flutter/material.dart';
import 'package:auto_route/auto_route.dart';
import '../features/chat/presentation/chat_page.dart';
import '../features/login/presentation/auth_page.dart';

part 'app_router.gr.dart';

@AutoRouterConfig()
class AppRouter extends _$AppRouter {
  @override
  List<AutoRoute> get routes => [
    AutoRoute(page: InitRoute.page, initial: true),
    AutoRoute(page: AuthRoute.page),
    AutoRoute(page: ChatRoute.page),
    AutoRoute(page: HomeRoute.page),
  ];
  
}