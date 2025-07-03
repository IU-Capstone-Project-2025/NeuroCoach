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
])
void main() {
  group('InitBloc', () {
    late MockInitRepository mockInitRepository;
    late MockHealthRepository mockHealthRepository;
    late InitBloc bloc;

    setUp(() {
      mockInitRepository = MockInitRepository();
      mockHealthRepository = MockHealthRepository();
      bloc = InitBloc(
        initRepository: mockInitRepository,
        healthRepository: mockHealthRepository,
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
