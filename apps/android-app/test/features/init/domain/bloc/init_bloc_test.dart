import 'package:android_app/features/login/domain/repositories/remember_me_repository.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'package:android_app/features/init/domain/bloc/init_bloc.dart';
import 'package:android_app/features/init/domain/repositories/init_repository.dart';
import 'package:android_app/features/init/domain/repositories/health_repository.dart';

import 'init_bloc_test.mocks.dart';

@GenerateMocks([
  InitRepository,
  HealthRepository,
  RememberMeRepository,
])
void main() {
  group('InitBloc', () {
    late MockInitRepository mockInitRepository;
    late MockHealthRepository mockHealthRepository;
    late MockRememberMeRepository mockRememberMeRepository;
    late InitBloc bloc;

    setUp(() {
      mockInitRepository = MockInitRepository();
      mockHealthRepository = MockHealthRepository();
      mockRememberMeRepository = MockRememberMeRepository();
      bloc = InitBloc(
        initRepository: mockInitRepository,
        healthRepository: mockHealthRepository,
        rememberMeRepository: mockRememberMeRepository
      );
    });

    test('emits [InitStateUnauthenticated] when token is empty', () async {
      when(mockInitRepository.getJWTToken()).thenAnswer((_) async => '');

      expectLater(
        bloc.stream,
        emitsInOrder([
          isA<InitStateUnauthenticated>(),
        ]),
      );

      bloc.add(const InitEventCheck());
    });

    test('emits [InitStateAuthenticated] when token is valid', () async {
      when(mockInitRepository.getJWTToken()).thenAnswer((_) async => 'token');
      when(mockHealthRepository.checkToken('token')).thenAnswer((_) async => true);

      expectLater(
        bloc.stream,
        emitsInOrder([
          isA<InitStateAuthenticated>(),
        ]),
      );

      bloc.add(const InitEventCheck());
    });

    test('emits [InitStateUnauthenticated] and removes token when token is invalid', () async {
      when(mockInitRepository.getJWTToken()).thenAnswer((_) async => 'token');
      when(mockHealthRepository.checkToken('token')).thenAnswer((_) async => false);
      when(mockInitRepository.removeJWTToken()).thenAnswer((_) async => true);

      expectLater(
        bloc.stream,
        emitsInOrder([
          isA<InitStateUnauthenticated>(),
        ]),
      );

      bloc.add(const InitEventCheck());
      await untilCalled(mockInitRepository.removeJWTToken());
      verify(mockInitRepository.removeJWTToken()).called(1);
    });
  });
}
