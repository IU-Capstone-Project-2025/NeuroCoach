import 'package:android_app/features/settings/domain/user_profile_entity.dart';

abstract class ProfileRepository {
  Future<UserProfile?> getUserProfile();
}