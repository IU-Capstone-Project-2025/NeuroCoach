import 'package:android_app/app/app_router.dart';
import 'package:android_app/app/presentation/scopes/dependencies_scope.dart';
import 'package:android_app/features/init/domain/bloc/init_bloc.dart';
import 'package:auto_route/auto_route.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

@RoutePage()
class InitPage extends StatelessWidget {
  const InitPage({super.key});

  @override
  Widget build(BuildContext context) {
    final depScope = DependenciesScope.findAppDependenciesOf(context);

    return BlocProvider(
      create: (_) => InitBloc(
        initRepository: depScope.initRepository,
        healthRepository: depScope.healthRepository,
      )..add(InitEventCheck()),
      child: BlocListener<InitBloc, InitState>(
        listener: (context, state) {
          if (state is InitStateAuthenticated) {
            context.router.replace(const HomeRoute());
          } else if (state is InitStateUnauthenticated) {
            context.router.replace(AuthRoute(isSignUp: false));
          }
        },
        child: const Scaffold(body: Center(child: CircularProgressIndicator())),
      ),
    );
  }
}
