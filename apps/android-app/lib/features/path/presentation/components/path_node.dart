import 'package:android_app/constants/app_colors.dart';
import 'package:android_app/constants/app_text_styles.dart';
import 'package:android_app/uikit/inner_shadow.dart';
import 'package:flutter/material.dart';

class PathNode extends StatelessWidget {
  final NodeStatus status;
  final VoidCallback onTap;
  final int index;

  const PathNode({
    super.key,
    required this.status,
    required this.onTap,
    required this.index,
  });

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: onTap,
      child: InnerShadow(
        shadows: [
          Shadow(
            color: AppColors.black,
            offset: Offset(-2, 2),
            blurRadius: 7,
          )
        ],
        child: Container(
          height: 64,
          width: 64,
          decoration: BoxDecoration(
            color: _getNodeColor(),
            borderRadius: BorderRadius.circular(16.0),
            border: Border.all(
              width: 0.5,
              color: AppColors.white,
            ),
            boxShadow: [
              BoxShadow(
                  color: AppColors.deepBlue,
                  blurRadius: 10
              )
            ]
          ),
          child: Center(
            child: Text(
              index.toString(),
              style: AppTextStyles.chatTitle.copyWith(
                fontSize: 32,
                color: _getTextColor(),
              ),
            ),
          ),
        ),
      ),
    );
  }

  Color _getNodeColor() {
    return switch (status) {
      NodeStatus.current => AppColors.white,
      NodeStatus.planned => AppColors.blackNode,
      NodeStatus.done => AppColors.lily,
    };
  }

  Color _getTextColor() {
    return switch (status) {
      NodeStatus.current => AppColors.lily,
      NodeStatus.planned => AppColors.black,
      NodeStatus.done => AppColors.white,
    };
  }
}

enum NodeStatus { current, planned, done }
