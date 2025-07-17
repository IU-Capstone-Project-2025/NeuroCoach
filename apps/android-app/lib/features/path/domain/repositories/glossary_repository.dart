import 'package:android_app/features/path/domain/entities/workout_entity.dart';

abstract class GlossaryRepository {
  Future<List<SimplifiedExerciseEntity>> getGlossary();
}