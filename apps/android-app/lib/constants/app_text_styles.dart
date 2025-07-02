import 'package:android_app/constants/app_colors.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

class AppTextStyles {
  static TextStyle get textButton {
    return GoogleFonts.arOneSans(
      color: Color(0xFFFFFFFF),
      fontWeight: FontWeight.w400,
      fontSize: 14,
      height: 36 / 14,
    );
  }

  static TextStyle get textField {
    return GoogleFonts.arOneSans(
      color: AppColors.grey,
      fontWeight: FontWeight.w200,
      fontSize: 12,
      height: 18 / 12,
    );
  }

  static TextStyle get chatTitle {
    return GoogleFonts.anta(
      color: Color(0xFFFFFFFF),
      fontWeight: FontWeight.w400,
      fontSize: 18,
      height: 36 / 18
    );
  }

}