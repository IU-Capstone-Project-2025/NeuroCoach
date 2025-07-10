import 'package:android_app/constants/app_text_styles.dart';
import 'package:android_app/features/path/domain/entities/workout_entity.dart';
import 'package:android_app/features/path/presentation/components/exercise_overview.dart';
import 'package:auto_route/annotations.dart';
import 'package:flutter/material.dart';

@RoutePage()
class GlossaryPage extends StatefulWidget {
  const GlossaryPage({super.key, required this.exercises});

  final List<Exercise> exercises;

  @override
  State<StatefulWidget> createState() => _GlossaryPageState();

}

class _GlossaryPageState extends State<GlossaryPage> {

  int _expandedIndex = 0;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Glossary'),
        titleTextStyle: AppTextStyles.chatTitle,
        centerTitle: true,
      ),
      body: ExerciseOverview(
        exercises: widget.exercises,
        performedSets: {1: 0, 2: 0, 3: 0},
        expandedIndex: _expandedIndex,
        onExpansionChanged: (index) {
          setState(() {
            if (index != null) {
              _expandedIndex = index;
            }
          });
        },
        isTrainingStarted: false,
        onStartTraining: () => {},
        isCurrentTraining: false,
      ),);
  }

}