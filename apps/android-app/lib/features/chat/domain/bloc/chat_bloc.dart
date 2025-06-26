import 'package:android_app/features/chat/domain/entities/chat_message_entity.dart';
import 'package:android_app/features/chat/domain/repositories/chat_messages_repository.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class ChatBloc extends Bloc<ChatEvent, ChatState> {
  ChatBloc({required ChatMessagesRepository chatMessagesRepository})
    : _chatMessagesRepository = chatMessagesRepository,
      super(ChatStateInitial(messages: [])) {
    on<ChatEvent>(
      (event, emit) => switch (event) {
        ChatEventSend() => _onMessageSent(event, emit),
        ChatEventRetrieve() => _onChatRetrieve(event, emit),
      },
    );
  }

  Future<void> _onMessageSent(
    ChatEventSend event,
    Emitter<ChatState> emit,
  ) async {
    final userMessage = ChatMessage(
      id: DateTime.now().millisecondsSinceEpoch.toString(),
      message: event.message,
      response: null,
      isUser: true,
      createdAt: DateTime.now(),
    );
    _messages.add(userMessage);
    emit(ChatStateLoaded(messages: [..._messages]));
    try {
      final response = await _chatMessagesRepository.sendMessage(event.message);
      final botMessage = ChatMessage(
        id: (DateTime.now().millisecondsSinceEpoch + 1).toString(),
        message: event.message,
        response: response,
        isUser: false,
        createdAt: DateTime.now(),
      );
      _messages.add(botMessage);
      emit(ChatStateLoaded(messages: [..._messages]));
    } on Object catch (e, s) {
      if (kDebugMode) print('$e, $s');
      _messages.add(
        ChatMessage(
          id: (DateTime.now().millisecondsSinceEpoch + 1).toString(),
          message: event.message,
          response: 'Something went wrong....',
          isUser: false,
          createdAt: DateTime.now(),
        ),
      );
      emit(ChatStateError(messages: _messages));
    }
  }

  Future<void> _onChatRetrieve(
      ChatEventRetrieve event,
      Emitter<ChatState> emit,
      ) async {
    emit(ChatStateLoading(messages: _messages));
    try {
      final rawMessages = await _chatMessagesRepository.fetchMessages();

      _messages.clear();

      for (final raw in rawMessages) {

        _messages.add(ChatMessage(
          id: raw.id,
          message: raw.message,
          response: raw.response,
          isUser: true,
          createdAt: raw.createdAt,
        ));

        if (raw.response != null && raw.response!.isNotEmpty) {
          _messages.add(ChatMessage(
            id: raw.id,
            message: raw.response!,
            response: raw.response!,
            isUser: false,
            createdAt: raw.createdAt,
          ));
        }
      }

      emit(ChatStateLoaded(messages: _messages));
    } on Object catch (e, s) {
      if (kDebugMode) print('$e, $s');
      emit(ChatStateError(messages: _messages));
    }
  }


  final ChatMessagesRepository _chatMessagesRepository;
  final List<ChatMessage> _messages = [];
}

sealed class ChatEvent {
  const ChatEvent();
}

class ChatEventSend extends ChatEvent {
  const ChatEventSend({required this.message});

  final String message;
}

class ChatEventRetrieve extends ChatEvent {
  const ChatEventRetrieve();
}

sealed class ChatState {
  const ChatState({required this.messages});

  final List<ChatMessage> messages;
}

class ChatStateInitial extends ChatState {
  const ChatStateInitial({required super.messages});
}

class ChatStateLoaded extends ChatState {
  const ChatStateLoaded({required super.messages});
}

class ChatStateLoading extends ChatState {
  const ChatStateLoading({required super.messages});
}

class ChatStateError extends ChatState {
  const ChatStateError({required super.messages});
}
