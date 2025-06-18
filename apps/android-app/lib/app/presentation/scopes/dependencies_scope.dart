import 'package:android_app/app/app_dependencies.dart';
import 'package:flutter/material.dart';

class DependenciesScope extends InheritedWidget {

  @override
  bool updateShouldNotify(covariant InheritedWidget oldWidget) {
    final old = oldWidget as DependenciesScope;
    return appDependencies != old.appDependencies;
  }

  const DependenciesScope({
    super.key,
    required this.appDependencies,
    required super.child,
  });

  static AppDependencies findAppDependenciesOf(BuildContext context) => of(context);

  static AppDependencies of(BuildContext context) {
    final scope = context.dependOnInheritedWidgetOfExactType<DependenciesScope>();
    assert(scope != null, 'No DependenciesScope found in context');
    return scope!.appDependencies;
  }

  final AppDependencies appDependencies;

}