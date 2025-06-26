import 'package:flutter/material.dart';

import '../../domain/entities/app_config.dart';

class AppConfigScope extends InheritedWidget {

  const AppConfigScope({
    super.key,
    required this.appConfig,
    required super.child,
  });

  final AppConfig appConfig;

  static AppConfig of(BuildContext context) {
    final scope =
    context.dependOnInheritedWidgetOfExactType<AppConfigScope>();
    assert(scope != null, 'No AppConfigScope found in context');
    return scope!.appConfig;
  }

  @override
  bool updateShouldNotify(covariant AppConfigScope oldWidget) {
    return oldWidget.appConfig != appConfig;
  }
}
