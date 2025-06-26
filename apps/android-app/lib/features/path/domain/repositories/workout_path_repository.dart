abstract class WorkoutPathRepository {
  Future<List<Map<String, dynamic>>> fetchWorkouts();
  Future<bool> generate();
}