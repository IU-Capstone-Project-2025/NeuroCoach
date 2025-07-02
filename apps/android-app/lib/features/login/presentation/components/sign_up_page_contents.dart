import 'package:android_app/app/app_router.dart';
import 'package:android_app/constants/app_colors.dart';
import 'package:android_app/constants/app_text_styles.dart';
import 'package:android_app/features/login/domain/bloc/sign_up_bloc.dart';
import 'package:android_app/uikit/app_text_field.dart';
import 'package:android_app/uikit/buttons/app_button.dart';
import 'package:auto_route/auto_route.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../domain/bloc/login_bloc.dart';

class SignUpPageContents extends StatefulWidget {
  const SignUpPageContents({super.key});

  @override
  State<StatefulWidget> createState() => SignUpPageContentsState();
}

class SignUpPageContentsState extends State<SignUpPageContents> {
  late final TextEditingController _emailController;
  late final TextEditingController _passwordController;

  late final TextEditingController _ageController;
  late final TextEditingController _heightController;
  late final TextEditingController _weightController;
  late final TextEditingController _minutesController;
  late final TextEditingController _healthIssuesController;

  String? _selectedGoal;
  String? _selectedTimeframe;
  String? _selectedLevel;

  final goals = [
    'weight_loss',
    'muscle_gain',
    'endurance',
    'flexibility',
    'general_fitness',
  ];
  final timeframes = ['1month', '3months', '6months', '1year'];
  final levels = ['beginner', 'intermediate', 'advanced'];

  @override
  void initState() {
    _emailController = TextEditingController();
    _passwordController = TextEditingController();
    _ageController = TextEditingController();
    _heightController = TextEditingController();
    _weightController = TextEditingController();
    _minutesController = TextEditingController();
    _healthIssuesController = TextEditingController();
    super.initState();
  }

  @override
  void dispose() {
    _ageController.dispose();
    _heightController.dispose();
    _weightController.dispose();
    _minutesController.dispose();
    _healthIssuesController.dispose();
    _emailController.dispose();
    _passwordController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Padding(
        padding: const EdgeInsets.symmetric(horizontal: 24),
        child: BlocBuilder<SignUpBloc, SignUpState>(
          builder: (context, state) {
            if (kDebugMode) print('State is ${state.runtimeType}');
            switch (state) {
              case SignUpStateInitial():
                return Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    AnimatedContainer(
                      duration: const Duration(milliseconds: 100),
                      curve: Curves.easeInOut,
                      height: MediaQuery.of(context).viewInsets.bottom > 0
                          ? 140
                          : 250,
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
                    const SizedBox(height: 16),
                    BlocBuilder<SignUpBloc, SignUpState>(
                      builder: (context, state) {
                        final isLoading = state is LoginStateLoading;
                        return SizedBox(
                          width: double.infinity,
                          child: AppButton(
                            onPressed: isLoading
                                ? null
                                : () {
                                    final email = _emailController.text.trim();
                                    final password = _passwordController.text
                                        .trim();
                                    context.read<SignUpBloc>().add(
                                      SignUpEventStartSignup(email, password),
                                    );
                                  },
                            padding: EdgeInsetsGeometry.symmetric(
                              horizontal: 16,
                            ),
                            isDisabled: isLoading,
                            text: 'Sign up',
                          ),
                        );
                      },
                    ),
                    Row(
                      mainAxisAlignment: MainAxisAlignment.center,
                      children: [
                        Text(
                          "Already have an account?",
                          style: AppTextStyles.textButton.copyWith(
                            fontWeight: FontWeight.w100,
                          ),
                        ),
                        TextButton(
                          onPressed: () => context.router.replace(
                            AuthRoute(isSignUp: false),
                          ),
                          child: Text(
                            'Log in',
                            style: AppTextStyles.textButton.copyWith(
                              color: AppColors.lily,
                              fontWeight: FontWeight.w100,
                            ),
                          ),
                        ),
                      ],
                    ),
                  ],
                );
              case SignUpStateGetAge():
                return Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Text(
                      'First, how old are you?',
                      style: AppTextStyles.textButton,
                    ),
                    const SizedBox(height: 12),
                    AppTextField(
                      controller: _ageController,
                      inputType: TextInputType.number,
                      hint: 'Age (13–120)',
                    ),
                    const SizedBox(height: 16),
                    AppButton(
                      text: 'Next',
                      onPressed: () {
                        if (kDebugMode) print('Click');
                        final age = int.tryParse(_ageController.text.trim());
                        if (age != null && age >= 13 && age <= 120) {
                          context.read<SignUpBloc>().add(
                            SignUpEventGetAge(age),
                          );
                        } else {
                          _showError(context, 'Age must be between 13 and 120');
                        }
                      },
                    ),
                  ],
                );
              case SignUpStateGetHeightWeight():
                return Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Text(
                      'Your height & weight',
                      style: AppTextStyles.textButton,
                    ),
                    const SizedBox(height: 12),
                    AppTextField(
                      controller: _heightController,
                      inputType: TextInputType.number,
                      hint: 'Height in cm',
                    ),
                    const SizedBox(height: 8),
                    AppTextField(
                      controller: _weightController,
                      inputType: TextInputType.number,
                      hint: 'Weight in kg',
                    ),
                    const SizedBox(height: 16),
                    AppButton(
                      text: 'Next',
                      onPressed: () {
                        final height = int.tryParse(
                          _heightController.text.trim(),
                        );
                        final weight = int.tryParse(
                          _weightController.text.trim(),
                        );
                        if (height != null &&
                            weight != null &&
                            height > 0 &&
                            weight > 0) {
                          context.read<SignUpBloc>().add(
                            SignUpEventGetHeightWeight(height, weight),
                          );
                        } else {
                          _showError(
                            context,
                            'Please enter valid height and weight',
                          );
                        }
                      },
                    ),
                  ],
                );
              case SignUpStateGetGoalTimeframe():
                return Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Text(
                      'Your goal and timeframe',
                      style: AppTextStyles.textButton,
                    ),
                    const SizedBox(height: 12),
                    DropdownButtonFormField<String>(
                      value: _selectedGoal,
                      decoration: InputDecoration(labelText: 'Goal'),
                      items: goals
                          .map(
                            (g) => DropdownMenuItem(value: g, child: Text(g)),
                          )
                          .toList(),
                      onChanged: (val) => setState(() => _selectedGoal = val),
                    ),
                    const SizedBox(height: 8),
                    DropdownButtonFormField<String>(
                      value: _selectedTimeframe,
                      decoration: InputDecoration(labelText: 'Timeframe'),
                      items: timeframes
                          .map(
                            (t) => DropdownMenuItem(value: t, child: Text(t)),
                          )
                          .toList(),
                      onChanged: (val) =>
                          setState(() => _selectedTimeframe = val),
                    ),
                    const SizedBox(height: 16),
                    AppButton(
                      text: 'Next',
                      onPressed: () {
                        if (_selectedGoal != null &&
                            _selectedTimeframe != null) {
                          context.read<SignUpBloc>().add(
                            SignUpEventGetGoalTimeframe(
                              _selectedGoal!,
                              _selectedTimeframe!,
                            ),
                          );
                        } else {
                          _showError(context, 'Select both goal and timeframe');
                        }
                      },
                    ),
                  ],
                );

              case SignUpStateGetLevelAvailableTime():
                return Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Text(
                      'Fitness level & workout duration',
                      style: AppTextStyles.textButton,
                    ),
                    const SizedBox(height: 12),
                    DropdownButtonFormField<String>(
                      value: _selectedLevel,
                      decoration: InputDecoration(labelText: 'Fitness Level'),
                      items: levels
                          .map(
                            (lvl) =>
                                DropdownMenuItem(value: lvl, child: Text(lvl)),
                          )
                          .toList(),
                      onChanged: (val) => setState(() => _selectedLevel = val),
                    ),
                    const SizedBox(height: 8),
                    AppTextField(
                      controller: _minutesController,
                      inputType: TextInputType.number,
                      hint: 'Minutes per workout (30–1000)',
                    ),
                    const SizedBox(height: 16),
                    AppButton(
                      text: 'Next',
                      onPressed: () {
                        final minutes = int.tryParse(
                          _minutesController.text.trim(),
                        );
                        if (_selectedLevel != null &&
                            minutes != null &&
                            minutes >= 30 &&
                            minutes <= 1000) {
                          context.read<SignUpBloc>().add(
                            SignUpEventGetLevelAvailableTime(
                              _selectedLevel!,
                              minutes,
                            ),
                          );
                        } else {
                          _showError(
                            context,
                            'Please select level and valid minutes',
                          );
                        }
                      },
                    ),
                  ],
                );
              case SignUpStateGetHealthIssues():
                return Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Text(
                      'Do you have any health issues?',
                      style: AppTextStyles.textButton,
                    ),
                    const SizedBox(height: 12),
                    AppTextField(
                      controller: _healthIssuesController,
                      hint: 'List any (comma-separated)',
                    ),
                    const SizedBox(height: 16),
                    AppButton(
                      text: 'Finish Sign Up',
                      onPressed: () {
                        final issues = _healthIssuesController.text
                            .split(',')
                            .map((e) => e.trim())
                            .where((e) => e.isNotEmpty)
                            .toList();
                        context.read<SignUpBloc>().add(
                          SignUpEventFinish(issues),
                        );
                      },
                    ),
                  ],
                );
              case SignUpStateFinish():
                context.read<LoginBloc>().add(
                  LoginEventSignUp(
                    email: state.email,
                    password: state.password,
                    height: state.height,
                    weight: state.weight,
                    age: state.age,
                    goal: state.goal,
                    healthIssues: state.healthIssues,
                    timeframe: state.timeframe,
                    fitnessLevel: state.level,
                    availableMinutes: state.minutes,
                    rememberMe: true,
                  ),
                );
                return Center(child: CircularProgressIndicator());
            }
          },
        ),
      ),
    );
  }

  void _showError(BuildContext context, String msg) {
    ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text(msg)));
  }
}
