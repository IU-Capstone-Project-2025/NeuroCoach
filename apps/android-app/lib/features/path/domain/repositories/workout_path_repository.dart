import 'package:android_app/features/path/domain/entities/workout_entity.dart';

abstract class WorkoutPathRepository {
  Future<List<WorkoutEntity>> fetchWorkouts();
  Future<bool> generate();
}