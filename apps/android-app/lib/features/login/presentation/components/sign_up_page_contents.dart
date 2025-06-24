import 'package:android_app/app/app_router.dart';
import 'package:android_app/constants/app_colors.dart';
import 'package:android_app/constants/app_text_styles.dart';
import 'package:android_app/uikit/app_text_field.dart';
import 'package:auto_route/auto_route.dart';
import 'package:flutter/material.dart';


class SignUpPageContents extends StatefulWidget {
  const SignUpPageContents({super.key});

  @override
  State<StatefulWidget> createState() => LoginPageContentsState();
}

class LoginPageContentsState extends State<SignUpPageContents> {
  late final TextEditingController _emailController;
  late final TextEditingController _passwordController;

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
              mainAxisAlignment: MainAxisAlignment.end,
              children: [
                TextButton(
                  onPressed: () => context.router.replace(LoginRoute()),
                  child: Text(
                    'Log in',
                    style: AppTextStyles.textButton.copyWith(
                      color: AppColors.lily,
                    ),
                  ),
                ),
              ],
            ),
            const SizedBox(height: 16),
          ],
        ),
      ),
    );
  }
}
