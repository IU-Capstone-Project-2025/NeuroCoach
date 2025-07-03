import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'package:android_app/features/login/domain/bloc/login_bloc.dart';
import 'package:android_app/features/login/domain/repositories/login_repository.dart';
import 'package:android_app/features/login/domain/repositories/remember_me_repository.dart';

import 'login_bloc_test.mocks.dart';

@GenerateMocks([
  LoginRepository,
  RememberMeRepository,
])
void main() {
  group('LoginBloc', () {
    late MockLoginRepository mockLoginRepository;
    late MockRememberMeRepository mockRememberMeRepository;
    late LoginBloc bloc;

    setUp(() {
      mockLoginRepository = MockLoginRepository();
      mockRememberMeRepository = MockRememberMeRepository();
      bloc = LoginBloc(
        loginRepository: mockLoginRepository,
        rememberMeRepository: mockRememberMeRepository,
      );
    });

    const email = 'test@example.com';
    const password = 'password';
    const token = 'token';
    const rememberMe = true;

    test('emits [Loading, Loaded] and calls rememberUser on login success', () async {
      when(mockLoginRepository.login(email, password, rememberMe)).thenAnswer((_) async => token);
      when(mockRememberMeRepository.rememberUser(jwtToken: token, email: email)).thenAnswer((_) async {});

      expectLater(
        bloc.stream,
        emitsInOrder([
          isA<LoginStateLoading>(),
          isA<LoginStateLoaded>().having((s) => s.email, 'email', email),
        ]),
      );

      bloc.add(const LoginEventLogin(email: email, password: password, rememberMe: rememberMe));
      await untilCalled(mockRememberMeRepository.rememberUser(jwtToken: token, email: email));
      verify(mockRememberMeRepository.rememberUser(jwtToken: token, email: email)).called(1);
    });

    test('emits [Loading, Loaded] and does not call rememberUser on login with empty token', () async {
      when(mockLoginRepository.login(email, password, rememberMe)).thenAnswer((_) async => '');

      expectLater(
        bloc.stream,
        emitsInOrder([
          isA<LoginStateLoading>(),
          isA<LoginStateLoaded>().having((s) => s.email, 'email', email),
        ]),
      );

      bloc.add(const LoginEventLogin(email: email, password: password, rememberMe: rememberMe));
      await untilCalled(mockLoginRepository.login(email, password, rememberMe));
      verifyNever(mockRememberMeRepository.rememberUser(jwtToken: anyNamed('jwtToken'), email: anyNamed('email')));
    });

    test('emits [Loading, Error] on login failure', () async {
      when(mockLoginRepository.login(email, password, rememberMe)).thenThrow(Exception('Login failed'));

      expectLater(
        bloc.stream,
        emitsInOrder([
          isA<LoginStateLoading>(),
          isA<LoginStateError>(),
        ]),
      );

      bloc.add(const LoginEventLogin(email: email, password: password, rememberMe: rememberMe));
    });

    test('emits [Loading, Loaded] and calls rememberUser on signup success', () async {
      when(mockLoginRepository.signUp(
        email,
        password,
        170,
        70,
        25,
        'Lose Weight',
        ['None'],
        '3 months',
        'Beginner',
        30,
        rememberMe,
      )).thenAnswer((_) async => token);
      when(mockRememberMeRepository.rememberUser(jwtToken: token, email: email)).thenAnswer((_) async {});

      expectLater(
        bloc.stream,
        emitsInOrder([
          isA<LoginStateLoading>(),
          isA<LoginStateLoaded>().having((s) => s.email, 'email', email),
        ]),
      );

      bloc.add(LoginEventSignUp(
        email: email,
        password: password,
        height: 170,
        weight: 70,
        age: 25,
        goal: 'Lose Weight',
        healthIssues: ['None'],
        timeframe: '3 months',
        fitnessLevel: 'Beginner',
        availableMinutes: 30,
        rememberMe: rememberMe,
      ));
      await untilCalled(mockRememberMeRepository.rememberUser(jwtToken: token, email: email));
      verify(mockRememberMeRepository.rememberUser(jwtToken: token, email: email)).called(1);
    });

    test('emits [Loading, Loaded] and does not call rememberUser on signup with empty token', () async {
      when(mockLoginRepository.signUp(
        email,
        password,
        170,
        70,
        25,
        'Lose Weight',
        ['None'],
        '3 months',
        'Beginner',
        30,
        rememberMe,
      )).thenAnswer((_) async => '');

      expectLater(
        bloc.stream,
        emitsInOrder([
          isA<LoginStateLoading>(),
          isA<LoginStateLoaded>().having((s) => s.email, 'email', email),
        ]),
      );

      bloc.add(LoginEventSignUp(
        email: email,
        password: password,
        height: 170,
        weight: 70,
        age: 25,
        goal: 'Lose Weight',
        healthIssues: ['None'],
        timeframe: '3 months',
        fitnessLevel: 'Beginner',
        availableMinutes: 30,
        rememberMe: rememberMe,
      ));
      await untilCalled(mockLoginRepository.signUp(
        email,
        password,
        170,
        70,
        25,
        'Lose Weight',
        ['None'],
        '3 months',
        'Beginner',
        30,
        rememberMe,
      ));
      verifyNever(mockRememberMeRepository.rememberUser(jwtToken: anyNamed('jwtToken'), email: anyNamed('email')));
    });

    test('emits [Loading, Error] on signup failure', () async {
      when(mockLoginRepository.signUp(
        email,
        password,
        170,
        70,
        25,
        'Lose Weight',
        ['None'],
        '3 months',
        'Beginner',
        30,
        rememberMe,
      )).thenThrow(Exception('Signup failed'));

      expectLater(
        bloc.stream,
        emitsInOrder([
          isA<LoginStateLoading>(),
          isA<LoginStateError>(),
        ]),
      );

      bloc.add(LoginEventSignUp(
        email: email,
        password: password,
        height: 170,
        weight: 70,
        age: 25,
        goal: 'Lose Weight',
        healthIssues: ['None'],
        timeframe: '3 months',
        fitnessLevel: 'Beginner',
        availableMinutes: 30,
        rememberMe: rememberMe,
      ));
    });

    test('emits [Initial] on reset', () async {
      expectLater(
        bloc.stream,
        emitsInOrder([
          isA<LoginStateInitial>(),
        ]),
      );
      bloc.add(const LoginEventReset());
    });
  });
}
