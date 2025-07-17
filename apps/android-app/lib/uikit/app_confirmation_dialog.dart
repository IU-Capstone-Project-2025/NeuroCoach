import 'package:flutter/material.dart';
import 'package:android_app/constants/app_colors.dart';
import 'package:android_app/constants/app_text_styles.dart';

class AppConfirmationDialog {
  static Future<bool> show(
      BuildContext context, {
        required String title,
      }) async {
    return await showDialog<bool>(
      context: context,
      barrierDismissible: false,
      builder: (context) => AlertDialog(
        backgroundColor: AppColors.blackNode,
        title: Text(
          title,
          style: AppTextStyles.textButton,
        ),
        actionsPadding: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
        actionsAlignment: MainAxisAlignment.center,
        actions: [
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
            children: [
              ElevatedButton(
                onPressed: () => Navigator.of(context).pop(true),
                style: ElevatedButton.styleFrom(
                  backgroundColor: AppColors.deepBlue,
                ),
                child: Text('Yes', style: AppTextStyles.textButton),
              ),
              const SizedBox(width: 16),
              ElevatedButton(
                onPressed: () => Navigator.of(context).pop(false),
                style: ElevatedButton.styleFrom(
                  backgroundColor: AppColors.lily,
                ),
                child: Text('No', style: AppTextStyles.textButton),
              ),
            ],
          ),
        ],
      ),
    ) ??
        false; // fallback if dismissed
  }
}
