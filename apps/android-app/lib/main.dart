import 'package:android_app/app/dependencies_factory.dart';
import 'package:android_app/app/presentation/scopes/dependencies_scope.dart';
import 'package:flutter/material.dart';

import 'app/app_router.dart';

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    final appRouter = AppRouter();
    final depScope = DependenciesFactory.build();
    return DependenciesScope(
      appDependencies: depScope,
      child: MaterialApp.router(
        routerDelegate: appRouter.delegate(),
        routeInformationParser: appRouter.defaultRouteParser(),
        theme: ThemeData.dark(useMaterial3: true),
      ),
    );
  }
}

void main() {
  runApp(MyApp());
}