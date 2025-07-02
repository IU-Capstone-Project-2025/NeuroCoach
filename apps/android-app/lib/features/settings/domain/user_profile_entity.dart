class UserProfile {
  final double height;
  final double weight;
  final int age;
  final String goal;
  final List<String> healthIssues;
  final String timeframe;
  final String fitnessLevel;
  final int availableMinutes;
  final DateTime updatedAt;

  const UserProfile({
    required this.height,
    required this.weight,
    required this.age,
    required this.goal,
    required this.healthIssues,
    required this.timeframe,
    required this.fitnessLevel,
    required this.availableMinutes,
    required this.updatedAt,
  });

  factory UserProfile.fromJson(Map<String, dynamic> json) {
    return UserProfile(
      height: (json['height'] as num).toDouble(),
      weight: (json['weight'] as num).toDouble(),
      age: json['age'] as int,
      goal: json['goal'] as String,
      healthIssues: List<String>.from(json['health_issues'] ?? []),
      timeframe: json['timeframe'] as String,
      fitnessLevel: json['fitness_level'] as String,
      availableMinutes: json['available_minutes'] as int,
      updatedAt: DateTime.parse(json['updated_at']),
    );
  }
}
