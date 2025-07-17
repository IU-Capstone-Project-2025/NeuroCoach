import 'package:android_app/app/presentation/scopes/app_config_scope.dart';
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
              child: Container(
                width: double.infinity,
                decoration: BoxDecoration(
                  color: AppColors.grey.withAlpha(99),
                  borderRadius: BorderRadius.circular(16.0),
                ),
                child: Padding(
                  padding: const EdgeInsets.symmetric(
                    horizontal: 16.0,
                    vertical: 16.0,
                  ),
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Row(
                        children: [
                          CircleAvatar(
                            backgroundColor: AppColors.grey,
                            radius: 12.0,
                            child: Icon(
                              Icons.person,
                              color: AppColors.messageGrey,
                              size: 16,
                            ),
                          ),
                          const SizedBox(width: 8),
                          Text(
                            AppConfigScope.of(context).email!,
                            style: AppTextStyles.textButton.copyWith(
                              fontSize: 20,
                            ),
                          ),
                        ],
                      ),
                      _buildInfoRow('Height', '${profile.height.ceil()} cm'),
                      _buildInfoRow('Weight', '${profile.weight.ceil()} kg'),
                      _buildInfoRow('Age', '${profile.age}'),
                      _buildInfoRow(
                        'Health Issues',
                        _formatList(profile.healthIssues),
                      ),
                      _buildInfoRow(
                        'Fitness Level',
                        _capitalize(profile.fitnessLevel),
                      ),
                    ],
                  ),
                ),
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
      padding: const EdgeInsets.symmetric(vertical: 8.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            title,
            style: AppTextStyles.textField.copyWith(color: AppColors.white),
          ),
          const SizedBox(height: 2),
          Text(value, style: AppTextStyles.textButton),
        ],
      ),
    );
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
