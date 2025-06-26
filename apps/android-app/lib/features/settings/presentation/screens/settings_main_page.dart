import 'package:auto_route/annotations.dart';
import 'package:auto_route/auto_route.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../app/app_router.dart';
import '../../../../app/presentation/scopes/app_config_scope.dart';
import '../../../../constants/app_colors.dart';
import '../../../../constants/app_text_styles.dart';
import '../../domain/bloc/logout_bloc.dart';

@RoutePage()
class SettingsMainPage extends StatelessWidget {
  const SettingsMainPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        automaticallyImplyLeading: false,
        title: Text('Settings'),
        titleTextStyle: AppTextStyles.chatTitle,
        centerTitle: true,
      ),
      body: Padding(
        padding: const EdgeInsets.symmetric(horizontal: 12, vertical: 24),
        child: Column(
          children: [
            GestureDetector(
              onTap: () => context.router.push(const ProfileRoute()),
              child: Container(
                decoration: BoxDecoration(
                  color: AppColors.lily,
                  border: Border.fromBorderSide(
                    BorderSide(color: Colors.transparent),
                  ),
                  borderRadius: BorderRadiusGeometry.circular(16),
                ),
                child: Padding(
                  padding: EdgeInsetsGeometry.symmetric(
                    horizontal: 6,
                    vertical: 8,
                  ),
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Row(
                        children: [
                          SizedBox(
                            height: 48,
                            width: 48,
                            child: CircleAvatar(
                              radius: 24,
                              backgroundColor: AppColors.grey,
                              child: Icon(Icons.person, size: 36),
                            ),
                          ),
                          const SizedBox(width: 8.0),
                          Column(
                            children: [
                              Text(
                                AppConfigScope.of(context).email!,
                                style: AppTextStyles.textButton.copyWith(
                                  color: AppColors.black,
                                ),
                              ),
                            ],
                          ),
                        ],
                      ),
                      Icon(Icons.arrow_forward),
                    ],
                  ),
                ),
              ),
            ),
            const SizedBox(height: 48),
            TextButton(
              onPressed: () => context.read<LogoutBloc>().add(LogoutEventLogout()),
              child: Text('Logout'),
            ),
          ],
        ),
      ),
    );
  }
}
