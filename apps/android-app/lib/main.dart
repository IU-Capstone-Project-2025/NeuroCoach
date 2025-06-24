import 'package:android_app/app/app_dependencies.dart';
import 'package:android_app/app/dependencies_factory.dart';
import 'package:android_app/app/presentation/scopes/dependencies_scope.dart';
import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';

import 'app/app_router.dart';


final _appRouter = AppRouter();

class MyApp extends StatelessWidget {
  const MyApp({super.key, required this.dependencies});

  final AppDependencies dependencies;

  @override
  Widget build(BuildContext context) {
    final appRouter = _appRouter;
    final depScope = dependencies;
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

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();
  SharedPreferences prefs = await SharedPreferences.getInstance();
  final depScope = DependenciesFactory.build(prefs);
  runApp(MyApp(dependencies: depScope));
}