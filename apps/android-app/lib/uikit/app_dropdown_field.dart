import 'package:flutter/material.dart';
import '../constants/app_colors.dart';
import '../constants/app_text_styles.dart';

class AppDropdownField<T> extends StatelessWidget {
  const AppDropdownField({
    super.key,
    required this.items,
    required this.value,
    required this.onChanged,
    this.hint,
    this.prefixIcon,
  });

  final List<T> items;
  final T? value;
  final void Function(T?) onChanged;
  final String? hint;
  final IconData? prefixIcon;

  @override
  Widget build(BuildContext context) {
    return InputDecorator(
      decoration: InputDecoration(
        prefixIcon: prefixIcon != null
            ? Icon(prefixIcon, color: AppColors.grey, size: 16)
            : null,
        hintText: hint,
        border: OutlineInputBorder(
          borderSide: BorderSide(color: AppColors.grey),
          borderRadius: BorderRadius.circular(16.0),
        ),
        hintStyle: AppTextStyles.textField,
        contentPadding: const EdgeInsets.symmetric(horizontal: 12, vertical: 4),
      ),
      child: DropdownButtonHideUnderline(
        child: DropdownButton<T>(
          isExpanded: true,
          value: value,
          hint: hint != null ? Text(hint!, style: AppTextStyles.textField) : null,
          icon: Icon(Icons.arrow_drop_down, color: AppColors.grey),
          dropdownColor: AppColors.black,
          borderRadius: BorderRadius.circular(16.0),
          items: items.map((T item) {
            return DropdownMenuItem<T>(
              value: item,
              child: Text(item.toString(), style: AppTextStyles.textField),
            );
          }).toList(),
          onChanged: onChanged,
        ),
      ),
    );
  }
}
