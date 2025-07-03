import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'package:android_app/features/settings/domain/bloc/logout_bloc.dart';
import 'package:android_app/features/settings/domain/repositories/logout_repository.dart';

import 'logout_bloc_test.mocks.dart';

@GenerateMocks([
  LogoutRepository,
])
void main() {
  group('LogoutBloc', () {
    late MockLogoutRepository mockLogoutRepository;
    late LogoutBloc bloc;

    setUp(() {
      mockLogoutRepository = MockLogoutRepository();
      bloc = LogoutBloc(logoutRepository: mockLogoutRepository);
    });

    test('emits [LoggedOut] when logout returns true', () async {
      when(mockLogoutRepository.logout()).thenAnswer((_) async => true);

      expectLater(
        bloc.stream,
        emitsInOrder([
          isA<LogoutStateLoggedOut>(),
        ]),
      );

      bloc.add(const LogoutEventLogout());
    });

    test('emits [Error] when logout returns false', () async {
      when(mockLogoutRepository.logout()).thenAnswer((_) async => false);

      expectLater(
        bloc.stream,
        emitsInOrder([
          isA<LogoutStateError>(),
        ]),
      );

      bloc.add(const LogoutEventLogout());
    });
  });
}
