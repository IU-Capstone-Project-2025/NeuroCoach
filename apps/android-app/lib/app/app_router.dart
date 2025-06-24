import 'package:android_app/features/home/presentation/home_page.dart';
import 'package:auto_route/auto_route.dart';
import '../features/chat/presentation/chat_page.dart';
import '../features/login/presentation/login_page.dart';
import '../features/login/presentation/sign_up_page.dart';

part 'app_router.gr.dart';

@AutoRouterConfig()
class AppRouter extends _$AppRouter {
  @override
  List<AutoRoute> get routes => [
    AutoRoute(page: LoginRoute.page, initial: true),
    AutoRoute(page: SignUpRoute.page),
    AutoRoute(page: ChatRoute.page),
    AutoRoute(page: HomeRoute.page),
  ];
  
}