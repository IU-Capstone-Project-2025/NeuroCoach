import 'package:android_app/app/app_router.dart';
import 'package:android_app/app/presentation/scopes/dependencies_scope.dart';
import 'package:android_app/features/login/presentation/login_page_contents/login_page_contents.dart';
import 'package:auto_route/auto_route.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../domain/bloc/login_bloc.dart';

@RoutePage()
class LoginPage extends StatelessWidget {
  const LoginPage({super.key});

  @override
  Widget build(BuildContext context) {
    final depScope = DependenciesScope.findAppDependenciesOf(context);

    return BlocProvider<LoginBloc>(
      create: (_) => LoginBloc(loginRepository: depScope.loginRepository),
      child: BlocListener<LoginBloc, LoginState>(
        listener: (context, state) {
          if (state is LoginStateLoaded) {
            context.router.push(HomeRoute());
          } else if (state is LoginStateError) {
            ScaffoldMessenger.of(
              context,
            ).showSnackBar(const SnackBar(content: Text('Error logging in')));
          }
        },
        child: LoginPageContents(),
      ),
    );
  }
}
