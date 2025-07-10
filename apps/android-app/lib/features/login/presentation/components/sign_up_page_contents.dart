import 'package:android_app/app/app_router.dart';
import 'package:android_app/constants/app_colors.dart';
import 'package:android_app/constants/app_text_styles.dart';
import 'package:android_app/features/login/domain/bloc/sign_up_bloc.dart';
import 'package:android_app/uikit/app_dropdown_field.dart';
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
  late final TextEditingController _hoursController;
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
    _hoursController = TextEditingController();
    _healthIssuesController = TextEditingController();
    super.initState();
  }

  @override
  void dispose() {
    _ageController.dispose();
    _heightController.dispose();
    _weightController.dispose();
    _minutesController.dispose();
    _hoursController.dispose();
    _healthIssuesController.dispose();
    _emailController.dispose();
    _passwordController.dispose();
    super.dispose();
  }

  void _showError(BuildContext context, String msg) {
    ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text(msg)));
  }

  bool _validateEmail(String email) {
    final emailRegExp = RegExp(
      r'^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$',
    );
    return emailRegExp.hasMatch(email);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        automaticallyImplyLeading: false,
        centerTitle: true,
        titleTextStyle: AppTextStyles.chatTitle,
        backgroundColor: Colors.transparent,
        actions: [
          BlocBuilder<SignUpBloc, SignUpState>(
            builder: (context, state) {
              if ((state is! SignUpStateInitial) &&
                  (state is! SignUpStateFinish)) {
                return TextButton(
                  onPressed: () =>
                      context.router.replace(AuthRoute(isSignUp: false)),
                  child: Text('To Login'),
                );
              } else {
                return const SizedBox.shrink();
              }
            },
          ),
        ],
        title: BlocBuilder<SignUpBloc, SignUpState>(
          builder: (context, state) {
            if ((state is! SignUpStateInitial) &&
                (state is! SignUpStateFinish)) {
              final int currentStep = state.step;
              final int totalSteps = 6;
              final double progress = currentStep / totalSteps;
              final String percentage = '${(progress * 100).round()}%';

              return Container(
                padding: const EdgeInsets.symmetric(
                  horizontal: 12,
                  vertical: 4,
                ),
                constraints: const BoxConstraints(minWidth: 180, maxWidth: 250),
                child: Stack(
                  alignment: Alignment.center,
                  children: [
                    ClipRRect(
                      borderRadius: BorderRadius.circular(30),
                      child: TweenAnimationBuilder<double>(
                        tween: Tween<double>(begin: 0, end: progress),
                        duration: const Duration(milliseconds: 400),
                        curve: Curves.easeInOut,
                        builder: (context, animatedProgress, _) {
                          return LinearProgressIndicator(
                            value: animatedProgress,
                            backgroundColor: Theme.of(
                              context,
                            ).scaffoldBackgroundColor,
                            valueColor: AlwaysStoppedAnimation<Color>(
                              AppColors.lily,
                            ),
                            minHeight: 20,
                          );
                        },
                      ),
                    ),
                    Positioned.fill(
                      child: Center(
                        child: Text(
                          percentage,
                          style: AppTextStyles.chatTitle.copyWith(height: 1),
                        ),
                      ),
                    ),
                  ],
                ),
              );
            } else {
              return const SizedBox.shrink();
            }
          },
        ),
      ),
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
                    Text(
                      'Begin your journey into sport',
                      style: AppTextStyles.textButton,
                    ),
                    const SizedBox(height: 16),
                    BlocBuilder<SignUpBloc, SignUpState>(
                      builder: (context, state) {
                        final isLoading = state is LoginStateLoading;
                        return SizedBox(
                          width: double.infinity,
                          child: AppButton(
                            onPressed: () => context.read<SignUpBloc>().add(
                              SignUpEventStartSignup(),
                            ),
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
                    AppDropdownField<String>(
                      value: _selectedGoal != null
                          ? LabelValueMapper.toGoalLabel(_selectedGoal!)
                          : null,
                      items: LabelValueMapper.goalMap.values.toList(),
                      onChanged: (label) => setState(() {
                        _selectedGoal = LabelValueMapper.toGoalValue(label!);
                      }),
                      hint: 'Select your goal',
                    ),
                    const SizedBox(height: 8),
                    AppDropdownField<String>(
                      value: _selectedTimeframe != null
                          ? LabelValueMapper.toTimeframeLabel(
                              _selectedTimeframe!,
                            )
                          : null,
                      items: LabelValueMapper.timeframeMap.values.toList(),
                      onChanged: (label) => setState(() {
                        _selectedTimeframe = LabelValueMapper.toTimeframeValue(
                          label!,
                        );
                      }),
                      hint: 'Select the timeframe',
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
                      'Fitness level & available time per week',
                      style: AppTextStyles.textButton,
                    ),
                    const SizedBox(height: 12),
                    AppDropdownField<String>(
                      value: _selectedLevel != null
                          ? LabelValueMapper.toLevelLabel(_selectedLevel!)
                          : null,
                      items: LabelValueMapper.levelMap.values.toList(),
                      onChanged: (label) => setState(() {
                        _selectedLevel = LabelValueMapper.toLevelValue(label!);
                      }),
                      hint: 'Select your level',
                    ),
                    const SizedBox(height: 8),
                    Row(
                      children: [
                        Expanded(
                          child: AppTextField(
                            controller: _hoursController,
                            inputType: TextInputType.number,
                            hint: 'Hours (0–16)',
                          ),
                        ),
                        const SizedBox(width: 12),
                        Expanded(
                          child: AppTextField(
                            controller: _minutesController,
                            inputType: TextInputType.number,
                            hint: 'Minutes (0–59)',
                          ),
                        ),
                      ],
                    ),
                    const SizedBox(height: 16),
                    AppButton(
                      text: 'Next',
                      onPressed: () {
                        final hours =
                            int.tryParse(_hoursController.text.trim()) ?? 0;
                        final minutes =
                            int.tryParse(_minutesController.text.trim()) ??
                            0;
                        final totalMinutes = hours * 60 + minutes;

                        if (_selectedLevel != null &&
                            totalMinutes >= 30 &&
                            totalMinutes <= 1000) {
                          context.read<SignUpBloc>().add(
                            SignUpEventGetLevelAvailableTime(
                              _selectedLevel!,
                              totalMinutes,
                            ),
                          );
                        } else {
                          _showError(
                            context,
                            'Please select a level and enter between 30 and 1000 total minutes.',
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
                      text: 'Next',
                      onPressed: () {
                        final issues = _healthIssuesController.text
                            .split(',')
                            .map((e) => e.trim())
                            .where((e) => e.isNotEmpty)
                            .toList();
                        context.read<SignUpBloc>().add(
                          SignUpEventGetHealthIssues(issues),
                        );
                      },
                    ),
                  ],
                );
              case SignUpStateGetCredentials():
                return Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Text(
                      'Finally, let us get your credentials',
                      style: AppTextStyles.textButton,
                    ),
                    const SizedBox(height: 12),
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
                                    if (_emailController.text.isNotEmpty &&
                                        _passwordController.text.isNotEmpty) {
                                      if (_validateEmail(
                                        _emailController.text,
                                      )) {
                                        final email = _emailController.text
                                            .trim();
                                        final password = _passwordController
                                            .text
                                            .trim();
                                        context.read<SignUpBloc>().add(
                                          SignUpEventFinish(email, password),
                                        );
                                      } else {
                                        _showError(
                                          context,
                                          'Please enter valid email',
                                        );
                                      }
                                    } else {
                                      _showError(
                                        context,
                                        'Please enter both email and password',
                                      );
                                    }
                                  },
                            padding: EdgeInsetsGeometry.symmetric(
                              horizontal: 16,
                            ),
                            isDisabled: isLoading,
                            text: 'Finish sign up',
                          ),
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
}

class LabelValueMapper {
  static final Map<String, String> goalMap = {
    'weight_loss': 'Weight Loss',
    'muscle_gain': 'Muscle Gain',
    'endurance': 'Endurance',
    'flexibility': 'Flexibility',
    'general_fitness': 'General Fitness',
  };

  static final Map<String, String> timeframeMap = {
    '1month': '1 Month',
    '3months': '3 Months',
    '6months': '6 Months',
    '1year': '1 Year',
  };

  static final Map<String, String> levelMap = {
    'beginner': 'Beginner',
    'intermediate': 'Intermediate',
    'advanced': 'Advanced',
  };

  static String toGoalLabel(String value) => goalMap[value] ?? value;

  static String toTimeframeLabel(String value) => timeframeMap[value] ?? value;

  static String toLevelLabel(String value) => levelMap[value] ?? value;

  static String toGoalValue(String label) =>
      goalMap.entries.firstWhere((e) => e.value == label).key;

  static String toTimeframeValue(String label) =>
      timeframeMap.entries.firstWhere((e) => e.value == label).key;

  static String toLevelValue(String label) =>
      levelMap.entries.firstWhere((e) => e.value == label).key;
}
