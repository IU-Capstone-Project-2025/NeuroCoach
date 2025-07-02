import 'package:android_app/features/path/domain/entities/WorkoutEntity.dart';

abstract class WorkoutPathRepository {
  Future<List<WorkoutEntity>> fetchWorkouts();
  Future<bool> generate();
}