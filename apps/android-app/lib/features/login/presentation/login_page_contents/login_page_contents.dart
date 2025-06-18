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
      appBar: AppBar(title: const Text('Login')),
      body: Padding(
        padding: const EdgeInsets.all(24),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            TextField(
              controller: _emailController,
              decoration: const InputDecoration(
                labelText: 'Email',
              ),
              keyboardType: TextInputType.emailAddress,
            ),
            const SizedBox(height: 16),
            TextField(
              controller: _passwordController,
              decoration: const InputDecoration(
                labelText: 'Password',
              ),
              obscureText: true,
            ),
            const SizedBox(height: 16),
            Row(
              children: [
                Checkbox(
                  value: _rememberMe,
                  onChanged: (val) {
                    setState(() => _rememberMe = val ?? false);
                  },
                ),
                const Text('Remember me'),
              ],
            ),
            const SizedBox(height: 24),
            BlocBuilder<LoginBloc, LoginState>(
              builder: (context, state) {
                final isLoading = state is LoginStateLoading;
                return SizedBox(
                  width: double.infinity,
                  child: ElevatedButton(
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
                    child: isLoading
                        ? const CircularProgressIndicator()
                        : const Text('Log In'),
                  ),
                );
              },
            ),
          ],
        ),
      ),
    );
  }

}