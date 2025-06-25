import 'dart:io';

import 'package:android_app/app/dependencies_factory.dart';
import 'package:android_app/app/presentation/scopes/dependencies_scope.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:shared_preferences/shared_preferences.dart';

import 'app/app_dependencies.dart';
import 'app/app_router.dart';
import 'constants/app_colors.dart';

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
        theme: ThemeData(
          splashColor: Colors.transparent,
          highlightColor: Colors.transparent,
          scaffoldBackgroundColor: AppColors.black,
          appBarTheme: AppBarTheme(
            backgroundColor: AppColors.black,
            iconTheme: IconThemeData(
              color: AppColors.white
            ),
          ),
          bottomNavigationBarTheme: BottomNavigationBarThemeData(
            backgroundColor: AppColors.black,
            selectedIconTheme: IconThemeData(
              color: AppColors.lily
            ),
              unselectedIconTheme: IconThemeData(
                  color: AppColors.white
              ),
          )
        ),
      ),
    );
  }
}

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();
  if (kDebugMode) print(Directory.current.path);
  // await dotenv.load(fileName: '/home/blazz1t/Projects/NeuroCoach/apps/android-app/.env');
  SharedPreferences prefs = await SharedPreferences.getInstance();
  final depScope = DependenciesFactory.build(prefs);
  runApp(MyApp(dependencies: depScope));
}