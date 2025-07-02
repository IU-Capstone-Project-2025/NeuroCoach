import 'package:android_app/constants/app_colors.dart';
import 'package:android_app/constants/app_text_styles.dart';
import 'package:flutter/material.dart';

class AppButton extends StatelessWidget {
  const AppButton({super.key, required this.onPressed, required this.text, this.padding = const EdgeInsetsGeometry.symmetric(horizontal: 16.0), this.isDisabled = false});

  final void Function()? onPressed;
  final String text;
  final bool isDisabled;
  final EdgeInsetsGeometry padding;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: padding,
      child: ElevatedButton(
        onPressed: isDisabled ? null : onPressed,
        style: ButtonStyle(
          backgroundColor: WidgetStateProperty.all(AppColors.purple),
          shadowColor: WidgetStateProperty.all(AppColors.black),
        ),
        child: Text(text, style: AppTextStyles.textButton),
      ),
    );
  }
}
