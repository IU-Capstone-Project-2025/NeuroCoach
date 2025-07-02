import 'package:android_app/constants/app_text_styles.dart';
import 'package:flutter/material.dart';

import '../constants/app_colors.dart';

class AppTextField extends StatelessWidget {
  const AppTextField({
    super.key,
    required this.controller,
    this.isPassword = false,
    this.prefixIcon,
    this.hint,
    this.inputType,
  });

  final TextEditingController controller;
  final bool isPassword;
  final IconData? prefixIcon;
  final String? hint;
  final TextInputType? inputType;

  @override
  Widget build(BuildContext context) {
    return TextField(
      controller: controller,
      obscureText: isPassword,
      keyboardType: inputType,
      decoration: InputDecoration(
        prefixIcon: Icon(prefixIcon, color: AppColors.grey, size: 16),
        hintText: hint,
        border: OutlineInputBorder(
          borderSide: BorderSide(color: AppColors.grey),
          borderRadius: BorderRadius.circular(16.0),
        ),
        hintStyle: AppTextStyles.textField
      ),
      style: AppTextStyles.textButton,
    );
  }
}
