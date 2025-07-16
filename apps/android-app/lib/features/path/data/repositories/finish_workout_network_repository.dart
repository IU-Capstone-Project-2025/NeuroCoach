import 'package:android_app/app/data/services/network_service.dart';
import 'package:android_app/features/path/domain/repositories/finish_workout_repository.dart';

class FinishWorkoutNetworkRepository extends FinishWorkoutRepository {
  @override
  Future<void> finishWorkout(String workoutId) async {
    final response = await NetworkService().request(
      method: 'POST',
      path: '/api/complete-workout',
      body: {'workout_id': workoutId}
    );

    if (response.statusCode == 200) {
      return;
    }
  }
}
