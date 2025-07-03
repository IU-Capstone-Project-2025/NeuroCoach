import 'package:android_app/features/chat/domain/bloc/chat_bloc.dart';
import 'package:android_app/features/path/domain/bloc/workout_path_bloc.dart';
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:mocktail/mocktail.dart';
import 'package:bloc_test/bloc_test.dart';
import 'package:android_app/features/settings/presentation/screens/profile_page.dart';
import 'package:android_app/features/settings/presentation/screens/settings_main_page.dart';
import 'package:android_app/features/settings/domain/bloc/profile_bloc.dart';
import 'package:android_app/features/settings/domain/bloc/logout_bloc.dart';
import 'package:android_app/features/settings/domain/user_profile_entity.dart';
import 'package:android_app/app/presentation/scopes/app_config_scope.dart';
import 'package:android_app/app/domain/entities/app_config.dart';

class MockProfileBloc extends Mock implements ProfileBloc {}
class MockLogoutBloc extends Mock implements LogoutBloc {}
class MockChatBloc extends Mock implements ChatBloc {}
class MockWorkoutBloc extends Mock implements WorkoutBloc {}

void main() {
  setUpAll(() {
    registerFallbackValue(ProfileStateInitial());
    registerFallbackValue(LogoutEventLogout());
    registerFallbackValue(LogoutStateInitial());
  });

  group('ProfilePage', () {
    testWidgets('shows loading indicator for loading/initial state', (tester) async {
      final bloc = MockProfileBloc();
      when(() => bloc.state).thenReturn(ProfileStateLoading());
      whenListen(bloc, Stream<ProfileState>.empty(), initialState: ProfileStateLoading());
      final chatBloc = MockChatBloc();
      final logoutBloc = MockLogoutBloc();
      final workoutBloc = MockWorkoutBloc();
      await tester.pumpWidget(
        MultiBlocProvider(
          providers: [
            BlocProvider<ProfileBloc>.value(value: bloc),
            BlocProvider<ChatBloc>.value(value: chatBloc),
            BlocProvider<LogoutBloc>.value(value: logoutBloc),
            BlocProvider<WorkoutBloc>.value(value: workoutBloc),
          ],
          child: MaterialApp(home: const ProfilePage()),
        ),
      );
      expect(find.byType(CircularProgressIndicator), findsOneWidget);
    });
    testWidgets('shows error message for error state', (tester) async {
      final bloc = MockProfileBloc();
      final chatBloc = MockChatBloc();
      final logoutBloc = MockLogoutBloc();
      final workoutBloc = MockWorkoutBloc();
      when(() => bloc.state).thenReturn(ProfileStateError());
      whenListen(bloc, Stream<ProfileState>.empty(), initialState: ProfileStateError());
      await tester.pumpWidget(
        MultiBlocProvider(
          providers: [
            BlocProvider<ProfileBloc>.value(value: bloc),
            BlocProvider<ChatBloc>.value(value: chatBloc),
            BlocProvider<LogoutBloc>.value(value: logoutBloc),
            BlocProvider<WorkoutBloc>.value(value: workoutBloc),
          ],
          child: MaterialApp(home: const ProfilePage()),
        ),
      );
      expect(find.text('Failed to load profile.'), findsOneWidget);
    });
    testWidgets('shows profile info for loaded state', (tester) async {
      final bloc = MockProfileBloc();
      final chatBloc = MockChatBloc();
      final logoutBloc = MockLogoutBloc();
      final workoutBloc = MockWorkoutBloc();
      final userProfile = UserProfile(
        height: 180.0,
        weight: 75.0,
        age: 30,
        goal: 'weight_loss',
        healthIssues: ['asthma'],
        timeframe: '3 months',
        fitnessLevel: 'beginner',
        availableMinutes: 45,
        updatedAt: DateTime.now(),
      );
      when(() => bloc.state).thenReturn(ProfileStateLoaded(userProfile: userProfile));
      whenListen(bloc, Stream<ProfileState>.empty(), initialState: ProfileStateLoaded(userProfile: userProfile));
      await tester.pumpWidget(
        MultiBlocProvider(
          providers: [
            BlocProvider<ProfileBloc>.value(value: bloc),
            BlocProvider<ChatBloc>.value(value: chatBloc),
            BlocProvider<LogoutBloc>.value(value: logoutBloc),
            BlocProvider<WorkoutBloc>.value(value: workoutBloc),
          ],
          child: MaterialApp(home: const ProfilePage()),
        ),
      );
      expect(find.text('180.0 cm'), findsOneWidget);
      expect(find.text('75.0 kg'), findsOneWidget);
      expect(find.text('30'), findsOneWidget);
      expect(find.text('Weight Loss'), findsOneWidget);
      expect(find.text('Asthma'), findsOneWidget);
      expect(find.text('3 months'), findsOneWidget);
      expect(find.text('Beginner'), findsOneWidget);
      expect(find.text('45 minutes'), findsOneWidget);
    });
  });

  group('SettingsMainPage', () {
    testWidgets('shows user email and profile tap', (tester) async {
      final logoutBloc = MockLogoutBloc();
      final chatBloc = MockChatBloc();
      final profileBloc = MockProfileBloc();
      final workoutBloc = MockWorkoutBloc();
      await tester.pumpWidget(
        MultiBlocProvider(
          providers: [
            BlocProvider<LogoutBloc>.value(value: logoutBloc),
            BlocProvider<ProfileBloc>.value(value: profileBloc),
            BlocProvider<ChatBloc>.value(value: chatBloc),
            BlocProvider<WorkoutBloc>.value(value: workoutBloc),
          ],
          child: MaterialApp(
            home: AppConfigScope(
              appConfig: const AppConfig(email: 'test@example.com'),
              child: const SettingsMainPage(),
            ),
          ),
        ),
      );
      expect(find.text('test@example.com'), findsOneWidget);
      expect(find.byKey(const Key('profile_tile')), findsOneWidget);
    });
    testWidgets('shows logout button and triggers event', (tester) async {
      final chatBloc = MockChatBloc();
      final profileBloc = MockProfileBloc();
      final workoutBloc = MockWorkoutBloc();
      final logoutBloc = MockLogoutBloc();
      when(() => logoutBloc.state).thenReturn(LogoutStateInitial());
      when(() => logoutBloc.stream).thenAnswer((_) => const Stream.empty());
      when(() => profileBloc.stream).thenAnswer((_) => const Stream.empty());
      when(() => chatBloc.stream).thenAnswer((_) => const Stream.empty());
      when(() => workoutBloc.stream).thenAnswer((_) => const Stream.empty());
      await tester.pumpWidget(
        MultiBlocProvider(
          providers: [
            BlocProvider<LogoutBloc>.value(value: logoutBloc),
            BlocProvider<ProfileBloc>.value(value: profileBloc),
            BlocProvider<ChatBloc>.value(value: chatBloc),
            BlocProvider<WorkoutBloc>.value(value: workoutBloc),
          ],
          child: MaterialApp(
            home: AppConfigScope(
              appConfig: const AppConfig(email: 'test@example.com'),
              child: const SettingsMainPage(),
            ),
          ),
        ),
      );
      await tester.tap(find.text('Logout'));
      await tester.pump();
      verify(() => logoutBloc.add(any(that: isA<LogoutEventLogout>()))).called(1);
    });
  });
}
