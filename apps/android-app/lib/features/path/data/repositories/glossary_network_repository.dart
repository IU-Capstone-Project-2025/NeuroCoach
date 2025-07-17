import 'package:android_app/app/data/services/network_service.dart';
import 'package:android_app/features/path/domain/entities/workout_entity.dart';
import 'package:android_app/features/path/domain/repositories/glossary_repository.dart';
import 'package:flutter/foundation.dart';

class GlossaryNetworkRepository extends GlossaryRepository {
  @override
  Future<List<SimplifiedExerciseEntity>> getGlossary() async {
    final response = await NetworkService().request(
      method: 'GET',
      path: '/api/media',
    );

    if (response.statusCode == 200) {
      final data = response.data;
      if (kDebugMode) print(data);
      if (data is List) {
        final parsedExercises = data
            .map((item) => SimplifiedExerciseEntity.fromJson(item))
            .toList();
        return parsedExercises;
      }
    }

    return [];
  }
}
