import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'package:android_app/features/settings/domain/bloc/profile_bloc.dart';
import 'package:android_app/features/settings/domain/repositories/profile_repository.dart';
import 'package:android_app/features/settings/domain/user_profile_entity.dart';

import 'profile_bloc_test.mocks.dart';

@GenerateMocks([
  ProfileRepository,
])
void main() {
  group('ProfileBloc', () {
    late MockProfileRepository mockProfileRepository;
    late ProfileBloc bloc;

    setUp(() {
      mockProfileRepository = MockProfileRepository();
      bloc = ProfileBloc(profileRepository: mockProfileRepository);
    });

    final userProfile = UserProfile(
      height: 180.0,
      weight: 75.0,
      age: 30,
      goal: 'Build Muscle',
      healthIssues: ['None'],
      timeframe: '6 months',
      fitnessLevel: 'Intermediate',
      availableMinutes: 45,
      updatedAt: DateTime.parse('2024-01-01T12:00:00Z'),
    );

    test('emits [Loading, Loaded] when getUserProfile returns a profile', () async {
      when(mockProfileRepository.getUserProfile()).thenAnswer((_) async => userProfile);

      expectLater(
        bloc.stream,
        emitsInOrder([
          isA<ProfileStateLoading>(),
          isA<ProfileStateLoaded>().having((s) => s.userProfile, 'userProfile', userProfile),
        ]),
      );

      bloc.add(const ProfileEventLoad());
    });

    test('emits [Loading, Error] when getUserProfile returns null', () async {
      when(mockProfileRepository.getUserProfile()).thenAnswer((_) async => null);

      expectLater(
        bloc.stream,
        emitsInOrder([
          isA<ProfileStateLoading>(),
          isA<ProfileStateError>(),
        ]),
      );

      bloc.add(const ProfileEventLoad());
    });

    test('emits [Loading, Error] when getUserProfile throws', () async {
      when(mockProfileRepository.getUserProfile()).thenThrow(Exception('error'));

      expectLater(
        bloc.stream,
        emitsInOrder([
          isA<ProfileStateLoading>(),
          isA<ProfileStateError>(),
        ]),
      );

      bloc.add(const ProfileEventLoad());
    });
  });
}
