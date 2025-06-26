import 'package:android_app/app/data/services/network_service.dart';
import 'package:android_app/features/settings/domain/repositories/profile_repository.dart';
import 'package:android_app/features/settings/domain/user_profile_entity.dart';

class ProfileNetworkRepository extends ProfileRepository {
  @override
  Future<UserProfile?> getUserProfile() async {
    final response = await NetworkService().request(
      method: 'GET',
      path: '/api/profile',
    );

    if (response.statusCode == 200) {
      return UserProfile.fromJson(response.data);
    }

    return null;
  }
}
