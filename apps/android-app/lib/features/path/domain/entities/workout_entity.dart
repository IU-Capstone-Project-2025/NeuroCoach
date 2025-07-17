import 'package:equatable/equatable.dart';

class WorkoutEntity implements EquatableMixin {
  final String workoutId;
  final String name;
  final String description;
  final String status;
  final List<Exercise> exercises;

  WorkoutEntity({
    required this.workoutId,
    required this.name,
    required this.description,
    required this.status,
    required this.exercises,
  });

  factory WorkoutEntity.fromJson(Map<String, dynamic> json) {
    return WorkoutEntity(
      workoutId: json['workout_id'],
      name: json['name'],
      description: json['description'],
      status: json['status'],
      exercises: (json['exercises'] as List<dynamic>)
          .map((e) => Exercise.fromJson(e))
          .toList(),
    );
  }

  @override
  List<Object?> get props => [workoutId, name, description, status, exercises];

  @override
  bool? get stringify => true;
}

class Exercise implements EquatableMixin {
  final String exerciseId;
  final String name;
  final String muscleGroup;
  final int sets;
  final int reps;
  final int restSec;
  final String notes;
  final String technique;
  final String? pictureUrl;

  Exercise({
    required this.exerciseId,
    required this.name,
    required this.muscleGroup,
    required this.sets,
    required this.reps,
    required this.restSec,
    required this.notes,
    required this.technique,
    this.pictureUrl,
  });

  factory Exercise.fromJson(Map<String, dynamic> json) {
    return Exercise(
      exerciseId: json['exercise_id'],
      name: json['name'],
      muscleGroup: json['muscle_group'],
      sets: json['sets'],
      reps: json['reps'],
      restSec: json['rest_sec'],
      notes: json['notes'],
      technique: json['technique'],
      pictureUrl: json['url'] != null ? json['url'] as String : null,
    );
  }

  @override
  List<Object?> get props => [
    exerciseId,
    name,
    muscleGroup,
    sets,
    reps,
    restSec,
    notes,
    technique,
    pictureUrl,
  ];

  @override
  bool? get stringify => true;
}

class SimplifiedExerciseEntity implements EquatableMixin {
  final String id;
  final String name;
  final String imageUrl;
  final String description;
  final DateTime createdAt;

  SimplifiedExerciseEntity({
    required this.id,
    required this.name,
    required this.imageUrl,
    required this.description,
    required this.createdAt,
  });

  factory SimplifiedExerciseEntity.fromJson(Map<String, dynamic> json) {
    return SimplifiedExerciseEntity(
      id: json['id'],
      name: json['name'],
      imageUrl: json['image_url'],
      description: json['description'],
      createdAt: DateTime.parse(json['created_at']),
    );
  }

  @override
  List<Object?> get props => [id, name, imageUrl, description, createdAt];

  @override
  bool? get stringify => true;
}
