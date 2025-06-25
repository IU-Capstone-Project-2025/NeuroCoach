import 'package:android_app/app/app_router.dart';
import 'package:android_app/constants/app_colors.dart';
import 'package:android_app/constants/app_text_styles.dart';
import 'package:android_app/uikit/app_checkbox.dart';
import 'package:android_app/uikit/app_text_field.dart';
import 'package:android_app/uikit/buttons/app_button.dart';
import 'package:auto_route/auto_route.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../domain/bloc/login_bloc.dart';

class LoginPageContents extends StatefulWidget {
  const LoginPageContents({super.key});

  @override
  State<StatefulWidget> createState() => LoginPageContentsState();
}

class LoginPageContentsState extends State<LoginPageContents> {
  late final TextEditingController _emailController;
  late final TextEditingController _passwordController;
  bool _rememberMe = false;

  @override
  void initState() {
    _emailController = TextEditingController();
    _passwordController = TextEditingController();
    super.initState();
  }

  @override
  void dispose() {
    _emailController.dispose();
    _passwordController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Padding(
        padding: const EdgeInsets.symmetric(horizontal: 24),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            AnimatedContainer(
              duration: const Duration(milliseconds: 100),
              curve: Curves.easeInOut,
              height: MediaQuery.of(context).viewInsets.bottom > 0 ? 140 : 250,
              child: Image.asset('assets/robot.png'),
            ),
            SizedBox(height: 36.0),
            AppTextField(
              controller: _emailController,
              prefixIcon: Icons.person,
              inputType: TextInputType.emailAddress,
              hint: 'Email',
            ),
            const SizedBox(height: 8.0),
            AppTextField(
              controller: _passwordController,
              prefixIcon: Icons.lock,
              isPassword: true,
              hint: 'Password',
            ),
            const SizedBox(height: 8.0),
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Row(
                  children: [
                    AppCheckbox(
                      value: _rememberMe,
                      onChanged: (val) {
                        setState(() => _rememberMe = val ?? false);
                      },
                    ),
                    Text('Remember me', style: AppTextStyles.textButton),
                  ],
                ),
              ],
            ),
            const SizedBox(height: 16),
            BlocBuilder<LoginBloc, LoginState>(
              builder: (context, state) {
                final isLoading = state is LoginStateLoading;
                return SizedBox(
                  width: double.infinity,
                  child: AppButton(
                    onPressed: isLoading
                        ? null
                        : () {
                            final email = _emailController.text.trim();
                            final password = _passwordController.text.trim();
                            context.read<LoginBloc>().add(
                              LoginEventLogin(
                                email: email,
                                password: password,
                                rememberMe: _rememberMe,
                              ),
                            );
                          },
                    padding: EdgeInsetsGeometry.symmetric(horizontal: 16),
                    isDisabled: isLoading,
                    text: 'Log in',
                  ),
                );
              },
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Text(
                  "Don't have an account?",
                  style: AppTextStyles.textButton.copyWith(
                    fontWeight: FontWeight.w100,
                  ),
                ),
                TextButton(
                  onPressed: () =>
                      context.router.replace(AuthRoute(isSignUp: true)),
                  child: Text(
                    'Sign up',
                    style: AppTextStyles.textButton.copyWith(
                      color: AppColors.lily,
                      fontWeight: FontWeight.w100,
                    ),
                  ),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
