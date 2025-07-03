import 'package:android_app/app/data/services/network_service.dart';
import 'package:android_app/features/path/domain/entities/workout_entity.dart';
import 'package:android_app/features/path/domain/repositories/workout_path_repository.dart';
import 'package:flutter/foundation.dart';
class WorkoutPathNetworkRepository extends WorkoutPathRepository {
  @override
  Future<List<WorkoutEntity>> fetchWorkouts() async {

    // if (kDebugMode) {
    //   return List.generate(20, (index) {
    //     return {
    //       "workout_id": "workout_$index",
    //       "name": "Workout ${index + 1}",
    //       "description": "Description for workout ${index + 1}",
    //       "status": index < 5 ? "complete" : "planned",
    //       "exercises": [
    //         {
    //           "exercise_id": "exercise_${index}_1",
    //           "name": "Push-ups",
    //           "muscle_group": "chest",
    //           "sets": 3,
    //           "reps": 12,
    //           "rest_sec": 60,
    //           "notes": "Keep core tight",
    //           "technique": "Slow and controlled movement"
    //         },
    //         {
    //           "exercise_id": "exercise_${index}_2",
    //           "name": "Squats",
    //           "muscle_group": "legs",
    //           "sets": 3,
    //           "reps": 15,
    //           "rest_sec": 60,
    //           "notes": "Keep back straight",
    //           "technique": "Go as low as possible"
    //         }
    //       ]
    //     };
    //   });
    // }

    final response = await NetworkService().request(
      method: 'GET',
      path: '/api/workout-plan',
    );



    if (response.statusCode == 200) {
      if (response.data['workouts'] is List) {
        final List<dynamic> workouts = response.data['workouts'];
        if (kDebugMode) print(workouts);
        return workouts.map((item) => WorkoutEntity.fromJson(item)).toList();
      }
    }

    return <WorkoutEntity>[];
  }

  @override
  Future<bool> generate() async {
    final response = await NetworkService().request(
      method: 'POST',
      path: '/api/generate-plan',
      body: {"regenerate": false},
    );

    return response.statusCode == 200;
  }
}
