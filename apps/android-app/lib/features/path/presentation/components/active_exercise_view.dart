import '../../../../constants/app_colors.dart';
import '../../../../constants/app_text_styles.dart';
import 'package:flutter/material.dart';

import '../../../../uikit/buttons/app_button.dart';
import '../../domain/entities/workout_entity.dart';

class ActiveExerciseView extends StatelessWidget {
  const ActiveExerciseView({
    super.key,
    required this.exercise,
    required this.performedSets,
    required this.onDoneWithSet,
  });

  final Exercise exercise;
  final int performedSets;
  final VoidCallback onDoneWithSet;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 24.0, vertical: 36),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Center(
            child: Text(
              "Exercise: ${exercise.name}",
              style: AppTextStyles.textButton.copyWith(fontSize: 22),
              textAlign: TextAlign.center,
            ),
          ),
          const SizedBox(height: 16),

          Row(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Text(
                "Set: ${performedSets + 1} / ${exercise.sets}",
                style: AppTextStyles.textButton.copyWith(fontSize: 18),
              ),
            ],
          ),

          const SizedBox(height: 24),
          Divider(color: AppColors.grey.withAlpha(100)),
          const SizedBox(height: 8),

          _infoRow("Reps", "${exercise.reps}"),
          if (exercise.technique.isNotEmpty)
            _infoRow("Technique", exercise.technique),
          if (exercise.notes.isNotEmpty) _infoRow("Notes", exercise.notes),

          const Spacer(),

          Center(
            child: AppButton(onPressed: onDoneWithSet, text: "Done with set"),
          ),
        ],
      ),
    );
  }

  Widget _infoRow(String title, String value) {
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 4.0),
      child: Row(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            "$title: ",
            style: AppTextStyles.textButton.copyWith(
              fontWeight: FontWeight.w600,
            ),
          ),
          Expanded(child: Text(value, style: AppTextStyles.textButton)),
        ],
      ),
    );
  }
}
