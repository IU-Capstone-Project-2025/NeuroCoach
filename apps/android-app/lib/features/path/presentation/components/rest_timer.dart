import 'dart:async';

import 'package:android_app/constants/app_text_styles.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';

class RestTimer extends StatefulWidget {
  const RestTimer({
    super.key,
    required this.restTime,
    required this.onFinished,
  });

  final int restTime;
  final VoidCallback onFinished;

  @override
  State<StatefulWidget> createState() => _RestTimerState();
}

class _RestTimerState extends State<RestTimer> {
  late int _remainingTime;
  Timer? _timer;

  @override
  void initState() {
    super.initState();
    _remainingTime = widget.restTime;

    _timer = Timer.periodic(const Duration(seconds: 1), (timer) {
      if (_remainingTime <= 1) {
        timer.cancel();
        widget.onFinished();
      } else {
        setState(() {
          _remainingTime--;
        });
      }
    });
  }

  @override
  void dispose() {
    _timer?.cancel();
    super.dispose();
  }

  String _formatTime(int seconds) {
    if (seconds < 60) {
      // Just seconds
      return '$seconds';
    } else {
      // MM:SS
      final minutes = seconds ~/ 60;
      final secs = seconds % 60;
      final secsStr = secs.toString().padLeft(2, '0');
      return '$minutes:$secsStr';
    }
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        Text('Time to rest', style: AppTextStyles.textButton),
        const SizedBox(height: 16.0),
        Text(
          _formatTime(_remainingTime),
          style: AppTextStyles.textButton.copyWith(fontSize: 32.0),
        ),
        if (kDebugMode) ...[
          const SizedBox(height: 24),
          ElevatedButton(
            onPressed: () {
              setState(() {
                _remainingTime = 5;
              });
            },
            child: const Text('Set to 5 seconds (debug)'),
          ),
        ],
      ],
    );
  }
}
