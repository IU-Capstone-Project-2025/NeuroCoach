import 'package:android_app/constants/app_text_styles.dart';
import 'package:flutter/material.dart';
import 'package:flutter_markdown/flutter_markdown.dart';

import '../../../../constants/app_colors.dart';
import '../../domain/entities/chat_message_entity.dart';

class ChatMessageBubble extends StatelessWidget {
  final ChatMessage message;

  const ChatMessageBubble({super.key, required this.message});

  @override
  Widget build(BuildContext context) {
    final isUser = message.isUser;

    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 8.0, vertical: 6.0),
      child: Row(
        mainAxisAlignment:
        isUser ? MainAxisAlignment.end : MainAxisAlignment.start,
        crossAxisAlignment: CrossAxisAlignment.end,
        children: [
          if (!isUser)
            Padding(
              padding: const EdgeInsets.only(right: 8.0),
              child: ClipRRect(
                borderRadius: BorderRadius.circular(12),
                child: Image.asset(
                  'assets/robot.png',
                  width: 24,
                  height: 24,
                  fit: BoxFit.cover,
                ),
              ),
            ),
          Flexible(
            child: Padding(
              padding: const EdgeInsets.only(bottom: 4.0),
              child: Container(
                padding: const EdgeInsets.symmetric(vertical: 4.0, horizontal: 16.0),
                decoration: BoxDecoration(
                  color: isUser ? AppColors.lily : AppColors.messageGrey,
                  borderRadius: BorderRadius.circular(12),
                ),
                child: isUser
                    ? Text(
                  message.message,
                  style: AppTextStyles.textButton,
                )
                    : MarkdownBody(
                  data: message.response ?? '',
                  styleSheet: MarkdownStyleSheet.fromTheme(Theme.of(context)).copyWith(
                    p: AppTextStyles.textButton,
                    h1: AppTextStyles.textButton.copyWith(fontSize: 24, fontWeight: FontWeight.bold),
                    h2: AppTextStyles.textButton.copyWith(fontSize: 20, fontWeight: FontWeight.bold),
                    h3: AppTextStyles.textButton.copyWith(fontSize: 18, fontWeight: FontWeight.bold),
                    h4: AppTextStyles.textButton.copyWith(fontSize: 16, fontWeight: FontWeight.bold),
                    h5: AppTextStyles.textButton.copyWith(fontSize: 14, fontWeight: FontWeight.bold),
                    h6: AppTextStyles.textButton.copyWith(fontSize: 12, fontWeight: FontWeight.bold),
                    listBullet: AppTextStyles.textButton,
                    listBulletPadding: const EdgeInsets.only(left: 4),
                    unorderedListAlign: WrapAlignment.start,
                    listIndent: 24,
                  ),
                ),
              ),
            ),
          ),
          if (isUser)
            Padding(
              padding: const EdgeInsets.only(left: 8.0),
              child: Container(
                width: 24,
                height: 24,
                decoration: BoxDecoration(
                  borderRadius: BorderRadius.circular(12),
                  color: AppColors.grey,
                ),
                child: const Icon(Icons.person, color: Colors.white, size: 16),
              ),
            ),
        ],
      ),
    );
  }
}
