import 'package:flutter/foundation.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class SignUpBloc extends Bloc<SignUpEvent, SignUpState> {
  SignUpBloc() : super(const SignUpStateInitial(0)) {
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

        SignUpEventGetHealthIssues() => _onGetHealthIssues(event, emit),

        SignUpEventFinish() => _onFinish(event, emit),

        SignUpEventReset() => _onReset(event, emit), //Reset if error
      },
    );
  }

  void _onStartSignup(SignUpEventStartSignup event, Emitter<SignUpState> emit) {
    if (kDebugMode) print('Got ${event.runtimeType} event');
    registrationStep++;
    emit(SignUpStateGetAge(registrationStep));
  }

  void _onGetAge(SignUpEventGetAge event, Emitter<SignUpState> emit) {
    if (kDebugMode) print('Got ${event.runtimeType} event');
    age = event.age;
    registrationStep++;
    emit(SignUpStateGetHeightWeight(registrationStep));
  }

  void _onGetHeightWeight(
    SignUpEventGetHeightWeight event,
    Emitter<SignUpState> emit,
  ) {
    if (kDebugMode) print('Got ${event.runtimeType} event');
    height = event.height;
    weight = event.weight;
    registrationStep++;
    emit(SignUpStateGetGoalTimeframe(registrationStep));
  }

  void _onGetGoalTimeframe(
    SignUpEventGetGoalTimeframe event,
    Emitter<SignUpState> emit,
  ) {
    if (kDebugMode) print('Got ${event.runtimeType} event');
    goal = event.goal;
    timeframe = event.timeframe;
    registrationStep++;
    emit(SignUpStateGetLevelAvailableTime(registrationStep));
  }

  void _onGetLevelAvailableTime(
    SignUpEventGetLevelAvailableTime event,
    Emitter<SignUpState> emit,
  ) {
    if (kDebugMode) print('Got ${event.runtimeType} event');
    level = event.level;
    minutes = event.availableTime;
    registrationStep++;
    emit(SignUpStateGetHealthIssues(registrationStep));
  }

  void _onGetHealthIssues(SignUpEventGetHealthIssues event, Emitter<SignUpState> emit) {
    if (kDebugMode) print('Got ${event.runtimeType} event');
    healthIssues = event.healthIssues;
    registrationStep++;
    emit(SignUpStateGetCredentials(registrationStep));
  }


  void _onFinish(SignUpEventFinish event, Emitter<SignUpState> emit) {
    if (kDebugMode) print('Got ${event.runtimeType} event');
    email = event.email;
    password = event.password;
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
          registrationStep
      ),
    );
  }

  void _onReset(SignUpEventReset event, Emitter<SignUpState> emit) {
    registrationStep = 0;
    emit(SignUpStateInitial(registrationStep));
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
  int registrationStep = 0;
}

sealed class SignUpEvent {
  const SignUpEvent();
}

class SignUpEventReset extends SignUpEvent {
  const SignUpEventReset();
}

class SignUpEventStartSignup extends SignUpEvent {
  const SignUpEventStartSignup();
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

class SignUpEventGetHealthIssues extends SignUpEvent {
  SignUpEventGetHealthIssues(this.healthIssues);
  final List<String> healthIssues;
}

class SignUpEventFinish extends SignUpEvent {
  const SignUpEventFinish(this.email, this.password);

  final String email;
  final String password;
}

sealed class SignUpState {
  const SignUpState(this.step);

  final int step;
}

class SignUpStateInitial extends SignUpState {
  const SignUpStateInitial(super.step);
}

class SignUpStateGetAge extends SignUpState {
  const SignUpStateGetAge(super.step);
}

class SignUpStateGetHeightWeight extends SignUpState {
  const SignUpStateGetHeightWeight(super.step);
}

class SignUpStateGetGoalTimeframe extends SignUpState {
  const SignUpStateGetGoalTimeframe(super.step);
}

class SignUpStateGetLevelAvailableTime extends SignUpState {
  const SignUpStateGetLevelAvailableTime(super.step);
}

class SignUpStateGetHealthIssues extends SignUpState {
  const SignUpStateGetHealthIssues(super.step);
}

class SignUpStateGetCredentials extends SignUpState {
  const SignUpStateGetCredentials(super.step);
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
    super.step,
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
