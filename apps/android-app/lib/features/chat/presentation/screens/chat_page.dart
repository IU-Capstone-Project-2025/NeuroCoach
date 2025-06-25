import 'package:android_app/constants/app_text_styles.dart';
import 'package:android_app/features/chat/domain/bloc/chat_bloc.dart';
import 'package:auto_route/annotations.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../constants/app_colors.dart';
import '../modals/chat_bubble.dart';

@RoutePage()
class ChatPage extends StatefulWidget {
  const ChatPage({super.key});

  @override
  State<StatefulWidget> createState() => ChatPageState();
}

class ChatPageState extends State<ChatPage> with TickerProviderStateMixin {
  final TextEditingController _controller = TextEditingController();
  final ScrollController _scrollController = ScrollController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        automaticallyImplyLeading: false,
        centerTitle: true,
        title: Text('Coach'),
        titleTextStyle: AppTextStyles.chatTitle,
        backgroundColor: Theme.of(context).appBarTheme.backgroundColor,
      ),
      body: Padding(
        padding: const EdgeInsets.only(left: 16.0, right: 16.0, top: 32.0),
        child: Container(
          decoration: BoxDecoration(
            color: AppColors.grey.withAlpha(54),
            border: BoxBorder.fromBorderSide(
              BorderSide(color: Colors.transparent),
            ),
            borderRadius: BorderRadiusGeometry.all(Radius.circular(16.0)),
          ),
          child: Padding(
            padding: const EdgeInsets.all(14.0),
            child: SizedBox.expand(
              child: Column(
                children: [
                  Expanded(
                    child: BlocConsumer<ChatBloc, ChatState>(
                      listener: (_, state) {
                        _scrollToBottom();
                      },
                      builder: (context, state) {
                        if (state is ChatStateLoading) {
                          return const Center(child: CircularProgressIndicator());
                        } else if (state is ChatStateLoaded) {
                          if (state.messages.isEmpty) {
                            return Center(
                              child: Text(
                                'No messages',
                                style: AppTextStyles.textButton,
                              ),
                            );
                          }
                          return ListView.builder(
                            controller: _scrollController,
                            itemCount: state.messages.length,
                            itemBuilder: (_, index) {
                              final msg = state.messages[index];
                              return ChatMessageBubble(message: msg);
                            },
                          );
                        } else if (state is ChatStateError) {
                          return ListView.builder(
                            controller: _scrollController,
                            itemCount: state.messages.length,
                            itemBuilder: (_, index) {
                              final msg = state.messages[index];
                              return ChatMessageBubble(message: msg);
                            },
                          );
                        } else {
                          return ListView.builder(
                            controller: _scrollController,
                            itemCount: state.messages.length,
                            itemBuilder: (_, index) {
                              final msg = state.messages[index];
                              return ChatMessageBubble(message: msg);
                            },
                          );
                        }
                      },
                    ),
                  ),
                  Padding(
                    padding: const EdgeInsets.all(16.0),
                    child: Container(
                      decoration: BoxDecoration(
                        color: AppColors.grey.withAlpha(54),
                        border: Border.fromBorderSide(
                          BorderSide(color: Colors.transparent),
                        ),
                        borderRadius: BorderRadius.all(Radius.circular(32.0))
                      ),
                      child: Row(
                        children: [
                          Expanded(
                            child: Padding(
                              padding: const EdgeInsets.only(left: 16.0),
                              child: TextField(
                                controller: _controller,
                                decoration: InputDecoration(
                                  border: InputBorder.none,
                                  hintText: 'Ask anything',
                                  hintStyle: AppTextStyles.textField.copyWith(
                                    color: Color(0xFFFFFFFF),
                                  ),
                                ),
                                style: AppTextStyles.textButton,
                                maxLines: 3,
                                minLines: 1,
                                keyboardType: TextInputType.multiline,
                              ),
                            ),
                          ),
                          IconButton(
                            icon: const Icon(
                              Icons.send,
                              color: Color(0xFF3B4168),
                            ),
                            onPressed: () {
                              final text = _controller.text.trim();
                              if (text.isNotEmpty) {
                                context.read<ChatBloc>().add(
                                  ChatEventSend(message: text),
                                );
                                _controller.clear();
                              }
                            },
                          ),
                        ],
                      ),
                    ),
                  ),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }

  void _scrollToBottom() {
    WidgetsBinding.instance.addPostFrameCallback((_) {
      _scrollController.animateTo(
        _scrollController.position.maxScrollExtent,
        duration: const Duration(milliseconds: 300),
        curve: Curves.easeOut,
      );
    });
  }
}
