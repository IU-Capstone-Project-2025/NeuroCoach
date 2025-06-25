import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../app/presentation/scopes/dependencies_scope.dart';
import '../../chat/domain/bloc/chat_bloc.dart';

class HomeScope extends StatelessWidget {
  const HomeScope({super.key, required this.child});

  final Widget child;

  @override
  Widget build(BuildContext context) {
    final depScope = DependenciesScope.findAppDependenciesOf(context);
    return MultiBlocProvider(
      providers: [
        BlocProvider(
          create: (_) =>
              ChatBloc(chatMessagesRepository: depScope.chatMessagesRepository)
                ..add(ChatEventRetrieve()),
        ),
      ],
      child: child,
    );
  }
}
