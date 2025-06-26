import 'package:android_app/constants/app_colors.dart';
import 'package:flutter/material.dart';

class AppCheckbox extends StatelessWidget {
  const AppCheckbox({super.key, required this.onChanged, required this.value});

  final bool value;
  final void Function(bool?) onChanged;

  @override
  Widget build(BuildContext context) {
    return Checkbox(
      value: value,
      onChanged: onChanged,
      activeColor: AppColors.purple,
      side: BorderSide(color: AppColors.grey),
    );
  }
}
