import 'package:android_app/app/app_router.dart';
import 'package:android_app/app/presentation/scopes/dependencies_scope.dart';
import 'package:android_app/features/login/domain/bloc/sign_up_bloc.dart';
import 'package:android_app/features/login/presentation/components/login_page_contents.dart';
import 'package:auto_route/auto_route.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../domain/bloc/login_bloc.dart';
import 'components/sign_up_page_contents.dart';

@RoutePage()
class AuthPage extends StatelessWidget {
  final bool isSignUp;

  const AuthPage({super.key, this.isSignUp = false});

  @override
  Widget build(BuildContext context) {
    final depScope = DependenciesScope.findAppDependenciesOf(context);

    return MultiBlocProvider(
      providers: [
        BlocProvider<LoginBloc>(
          create: (_) => LoginBloc(
            loginRepository: depScope.loginRepository,
            rememberMeRepository: depScope.rememberMeRepository,
          ),
        ),
        BlocProvider<SignUpBloc>(
          create: (_) => SignUpBloc(),
        ),
      ],
      child: MultiBlocListener(
        listeners: [
          BlocListener<LoginBloc, LoginState>(
            listener: (context, state) {
              if (state is LoginStateLoaded) {
                context.router.replace(HomeRoute());
              } else if (state is LoginStateError) {
                if (kDebugMode) {
                  context.router.push(HomeRoute());
                } else {
                  ScaffoldMessenger.of(context).showSnackBar(
                    const SnackBar(content: Text('Error logging in')),
                  );
                  context.read<LoginBloc>().add(LoginEventReset());
                }
              } else if (state is LoginStateInitial) {
                context.router.replace(AuthRoute(isSignUp: false));
              }
            },
          ),
        ],
        child: isSignUp
            ? const SignUpPageContents()
            : const LoginPageContents(),
      ),
    );
  }
}
