import 'package:flutter/material.dart';
import '../../domain/entities/app_config.dart';
import 'app_config_scope.dart';

class AppConfigScopeContainer extends StatefulWidget {
  final AppConfig initialConfig;
  final Widget child;

  const AppConfigScopeContainer({
    super.key,
    required this.initialConfig,
    required this.child,
  });

  static AppConfigScopeContainerState of(BuildContext context) {
    final state = context.findAncestorStateOfType<AppConfigScopeContainerState>();
    assert(state != null, 'No AppConfigScopeContainer found in context');
    return state!;
  }

  @override
  AppConfigScopeContainerState createState() => AppConfigScopeContainerState();
}

class AppConfigScopeContainerState extends State<AppConfigScopeContainer> {
  late AppConfig _config;

  @override
  void initState() {
    super.initState();
    _config = widget.initialConfig;
  }

  void updateEmail(String email) {
    setState(() {
      _config = AppConfig(email: email);
    });
  }

  @override
  Widget build(BuildContext context) {
    return AppConfigScope(
      appConfig: _config,
      child: widget.child,
    );
  }
}
