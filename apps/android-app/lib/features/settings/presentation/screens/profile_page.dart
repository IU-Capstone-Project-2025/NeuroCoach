import 'package:auto_route/annotations.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../constants/app_colors.dart';
import '../../../../constants/app_text_styles.dart';
import '../../domain/bloc/profile_bloc.dart';

@RoutePage()
class ProfilePage extends StatelessWidget {
  const ProfilePage({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<ProfileBloc, ProfileState>(
      builder: (context, state) {
        if (state is ProfileStateLoading || state is ProfileStateInitial) {
          return const Scaffold(
            body: Center(child: CircularProgressIndicator()),
          );
        } else if (state is ProfileStateError) {
          return Scaffold(
            appBar: AppBar(
              title: const Text('Profile'),
              titleTextStyle: AppTextStyles.chatTitle,
              centerTitle: true,
            ),
            body: Center(
              child: Text(
                'Failed to load profile.',
                style: AppTextStyles.textButton,
              ),
            ),
          );
        } else if (state is ProfileStateLoaded) {
          final profile = state.userProfile;

          return Scaffold(
            appBar: AppBar(
              title: const Text('Profile'),
              titleTextStyle: AppTextStyles.chatTitle,
              centerTitle: true,
            ),
            body: SingleChildScrollView(
              padding: const EdgeInsets.all(20.0),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  _buildInfoRow('Height', '${profile.height} cm'),
                  _buildInfoRow('Weight', '${profile.weight} kg'),
                  _buildInfoRow('Age', '${profile.age}'),
                  _buildInfoRow('Goal', _formatGoal(profile.goal)),
                  _buildInfoRow('Health Issues', _formatList(profile.healthIssues)),
                  _buildInfoRow('Timeframe', profile.timeframe),
                  _buildInfoRow('Fitness Level', _capitalize(profile.fitnessLevel)),
                  _buildInfoRow('Available Time/Day', '${profile.availableMinutes} minutes'),
                ],
              ),
            ),
          );
        } else {
          return const SizedBox.shrink(); // fallback
        }
      },
    );
  }

  Widget _buildInfoRow(String title, String value) {
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 12.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(title, style: AppTextStyles.textField.copyWith(color: AppColors.grey)),
          const SizedBox(height: 4),
          Text(value, style: AppTextStyles.textButton),
        ],
      ),
    );
  }

  String _formatGoal(String goal) {
    switch (goal) {
      case 'weight_loss':
        return 'Weight Loss';
      case 'muscle_gain':
        return 'Muscle Gain';
      case 'endurance':
        return 'Endurance';
      default:
        return _capitalize(goal);
    }
  }

  String _formatList(List<String> items) {
    if (items.isEmpty) return 'None';
    return items.map(_capitalize).join(', ');
  }

  String _capitalize(String input) {
    if (input.isEmpty) return input;
    return input[0].toUpperCase() + input.substring(1).replaceAll('_', ' ');
  }
}
