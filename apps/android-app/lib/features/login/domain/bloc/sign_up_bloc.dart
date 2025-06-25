import 'package:flutter/foundation.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class SignUpBloc extends Bloc<SignUpEvent, SignUpState> {
  SignUpBloc() : super(const SignUpStateInitial()) {
    on<SignUpEvent>(
      (event, emit) => switch (event) {
        SignUpEventStartSignup() => _onStartSignup(event, emit),

        SignUpEventGetAge() => _onGetAge(event, emit),

        SignUpEventGetHeightWeight() => _onGetHeightWeight(event, emit),

        SignUpEventGetGoalTimeframe() => _onGetGoalTimeframe(event, emit),

        SignUpEventGetLevelAvailableTime() => _onGetLevelAvailableTime(
          event,
          emit,
        ),

        SignUpEventFinish() => _onFinish(event, emit),
      },
    );
  }

  void _onStartSignup(SignUpEventStartSignup event, Emitter<SignUpState> emit) {
    if (kDebugMode) print('Got ${event.runtimeType} event');
    email = event.email;
    password = event.password;
    emit(const SignUpStateGetAge());
  }

  void _onGetAge(SignUpEventGetAge event, Emitter<SignUpState> emit) {
    if (kDebugMode) print('Got ${event.runtimeType} event');
    age = event.age;
    emit(const SignUpStateGetHeightWeight());
  }

  void _onGetHeightWeight(
    SignUpEventGetHeightWeight event,
    Emitter<SignUpState> emit,
  ) {
    if (kDebugMode) print('Got ${event.runtimeType} event');
    height = event.height;
    weight = event.weight;
    emit(const SignUpStateGetGoalTimeframe());
  }

  void _onGetGoalTimeframe(
    SignUpEventGetGoalTimeframe event,
    Emitter<SignUpState> emit,
  ) {
    if (kDebugMode) print('Got ${event.runtimeType} event');
    goal = event.goal;
    timeframe = event.timeframe;
    emit(const SignUpStateGetLevelAvailableTime());
  }

  void _onGetLevelAvailableTime(
    SignUpEventGetLevelAvailableTime event,
    Emitter<SignUpState> emit,
  ) {
    if (kDebugMode) print('Got ${event.runtimeType} event');
    level = event.level;
    minutes = event.availableTime;
    emit(const SignUpStateGetHealthIssues());
  }

  void _onFinish(SignUpEventFinish event, Emitter<SignUpState> emit) {
    if (kDebugMode) print('Got ${event.runtimeType} event');
    healthIssues = event.healthIssues;
    emit(
      SignUpStateFinish(
        email!,
        password!,
        age!,
        height!,
        weight!,
        goal!,
        timeframe!,
        level!,
        minutes!,
        healthIssues,
      ),
    );
  }

  String? email;
  String? password;
  int? age;
  int? height;
  int? weight;
  String? goal;
  String? timeframe;
  String? level;
  int? minutes;
  List<String> healthIssues = [];
}

sealed class SignUpEvent {
  const SignUpEvent();
}

class SignUpEventStartSignup extends SignUpEvent {
  const SignUpEventStartSignup(this.email, this.password);

  final String email;
  final String password;
}

class SignUpEventGetAge extends SignUpEvent {
  const SignUpEventGetAge(this.age);

  final int age;
}

class SignUpEventGetHeightWeight extends SignUpEvent {
  const SignUpEventGetHeightWeight(this.height, this.weight);

  final int height;
  final int weight;
}

class SignUpEventGetGoalTimeframe extends SignUpEvent {
  const SignUpEventGetGoalTimeframe(this.goal, this.timeframe);

  final String goal;
  final String timeframe;
}

class SignUpEventGetLevelAvailableTime extends SignUpEvent {
  const SignUpEventGetLevelAvailableTime(this.level, this.availableTime);

  final String level;
  final int availableTime;
}

class SignUpEventFinish extends SignUpEvent {
  const SignUpEventFinish(this.healthIssues);

  final List<String> healthIssues;
}

sealed class SignUpState {
  const SignUpState();
}

class SignUpStateInitial extends SignUpState {
  const SignUpStateInitial();
}

class SignUpStateGetAge extends SignUpState {
  const SignUpStateGetAge();
}

class SignUpStateGetHeightWeight extends SignUpState {
  const SignUpStateGetHeightWeight();
}

class SignUpStateGetGoalTimeframe extends SignUpState {
  const SignUpStateGetGoalTimeframe();
}

class SignUpStateGetLevelAvailableTime extends SignUpState {
  const SignUpStateGetLevelAvailableTime();
}

class SignUpStateGetHealthIssues extends SignUpState {
  const SignUpStateGetHealthIssues();
}

class SignUpStateFinish extends SignUpState {
  const SignUpStateFinish(
    this.email,
    this.password,
    this.age,
    this.height,
    this.weight,
    this.goal,
    this.timeframe,
    this.level,
    this.minutes,
    this.healthIssues,
  );

  final String email;
  final String password;
  final int age;
  final int height;
  final int weight;
  final String goal;
  final String timeframe;
  final String level;
  final int minutes;
  final List<String> healthIssues;
}
