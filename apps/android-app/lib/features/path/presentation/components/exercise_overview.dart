import 'package:flutter/material.dart';

import '../../../../constants/app_colors.dart';
import '../../../../constants/app_text_styles.dart';
import '../../../../uikit/buttons/app_button.dart';
import '../../domain/entities/workout_entity.dart';

class ExerciseOverview extends StatelessWidget {
  const ExerciseOverview({
    super.key,
    required this.exercises,
    required this.performedSets,
    required this.expandedIndex,
    required this.onExpansionChanged,
    required this.isTrainingStarted,
    required this.onStartTraining,
    required this.isCurrentTraining,
  });

  final List<Exercise> exercises;
  final Map<int, int> performedSets;
  final int? expandedIndex;
  final Function(int?) onExpansionChanged;
  final bool isTrainingStarted;
  final VoidCallback onStartTraining;
  final bool isCurrentTraining;

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Expanded(
          child: ListView.builder(
            padding: const EdgeInsets.all(16),
            itemCount: exercises.length,
            itemBuilder: (context, index) {
              final ex = exercises[index];
              final setsDone = performedSets[index] ?? 0;
              final isDone = setsDone >= ex.sets;

              final isDisabled = isTrainingStarted;
              final tileColor = isDone
                  ? AppColors.grey.withAlpha(50)
                  : AppColors.grey.withAlpha(90);

              return Padding(
                padding: const EdgeInsets.symmetric(vertical: 6.0),
                child: Container(
                  decoration: BoxDecoration(
                    color: tileColor,
                    borderRadius: BorderRadius.circular(16),
                  ),
                  child: IgnorePointer(
                    ignoring: isDisabled,
                    child: Theme(
                      data: Theme.of(
                        context,
                      ).copyWith(dividerColor: Colors.transparent),
                      child: ExpansionTile(
                        key: ValueKey(
                          'expansion_tile_$index${expandedIndex == index ? '_expanded' : ''}',
                        ),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(16.0),
                          side: const BorderSide(color: Colors.transparent),
                        ),
                        collapsedShape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(16.0),
                          side: const BorderSide(color: Colors.transparent),
                        ),
                        iconColor: AppColors.white,
                        collapsedIconColor: AppColors.white,
                        title: Text(
                          ex.name,
                          style: AppTextStyles.textButton.copyWith(
                            color: isDone ? Colors.grey : AppColors.white,
                          ),
                        ),
                        initiallyExpanded: expandedIndex == index,
                        onExpansionChanged: (expanded) =>
                            onExpansionChanged(expanded ? index : null),
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
                ),
              );
            },
          ),
        ),
        if (isCurrentTraining)
          Padding(
            padding: const EdgeInsets.symmetric(vertical: 48.0),
            child: AppButton(
              onPressed: onStartTraining,
              text: 'Start training',
            ),
          ),
      ],
    );
  }
}
