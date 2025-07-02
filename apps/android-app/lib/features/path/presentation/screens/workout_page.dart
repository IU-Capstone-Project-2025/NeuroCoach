import 'package:android_app/features/path/domain/entities/workout_entity.dart';
import 'package:auto_route/auto_route.dart';
import 'package:flutter/material.dart';

import '../../../../constants/app_colors.dart';
import '../../../../constants/app_text_styles.dart';

@RoutePage()
class WorkoutPage extends StatelessWidget {
  final String name;
  final List<Exercise> exercises;

  const WorkoutPage({super.key, required this.name, required this.exercises});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(name),
        titleTextStyle: AppTextStyles.chatTitle,
        centerTitle: true,
      ),
      body: ListView.builder(
        padding: const EdgeInsets.all(16),
        itemCount: exercises.length,
        itemBuilder: (context, index) {
          final ex = exercises[index];
          return Padding(
            padding: const EdgeInsets.symmetric(vertical: 6.0),
            child: Container(
              decoration: BoxDecoration(
                color: AppColors.grey.withAlpha(90),
                borderRadius: BorderRadius.circular(16),
              ),
              child: Theme(
                data: Theme.of(
                  context,
                ).copyWith(dividerColor: Colors.transparent),
                child: ExpansionTile(
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(16.0),
                    side: BorderSide(color: Colors.transparent),
                  ),
                  collapsedShape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(16.0),
                    side: BorderSide(color: Colors.transparent),
                  ),
                  iconColor: AppColors.white,
                  collapsedIconColor: AppColors.white,
                  title: Text(ex.name, style: AppTextStyles.textButton),
                  children: [
                    Padding(
                      padding: const EdgeInsets.all(12.0),
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children:
                            [
                                  Text(
                                    "Muscle Group: ${ex.muscleGroup}",
                                    style: AppTextStyles.textButton,
                                  ),
                                  Text(
                                    "Sets: ${ex.sets}",
                                    style: AppTextStyles.textButton,
                                  ),
                                  Text(
                                    "Reps: ${ex.reps}",
                                    style: AppTextStyles.textButton,
                                  ),
                                  Text(
                                    "Rest: ${ex.restSec} sec",
                                    style: AppTextStyles.textButton,
                                  ),
                                  Text(
                                    "Technique: ${ex.technique}",
                                    style: AppTextStyles.textButton,
                                  ),
                                  Text(
                                    "Notes: ${ex.notes}",
                                    style: AppTextStyles.textButton,
                                  ),
                                ]
                                .map(
                                  (e) => Padding(
                                    padding: const EdgeInsets.symmetric(
                                      vertical: 2,
                                    ),
                                    child: e,
                                  ),
                                )
                                .toList(),
                      ),
                    ),
                  ],
                ),
              ),
            ),
          );
        },
      ),
    );
  }
}
