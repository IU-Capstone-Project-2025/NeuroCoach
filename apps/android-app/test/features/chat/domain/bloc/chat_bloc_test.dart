import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'package:android_app/features/chat/domain/bloc/chat_bloc.dart';
import 'package:android_app/features/chat/domain/entities/chat_message_entity.dart';
import 'package:android_app/features/chat/domain/repositories/chat_messages_repository.dart';

import 'chat_bloc_test.mocks.dart';

@GenerateMocks([
  ChatMessagesRepository,
])
void main() {
  group('ChatBloc', () {
    late MockChatMessagesRepository mockChatMessagesRepository;
    late ChatBloc bloc;

    setUp(() {
      mockChatMessagesRepository = MockChatMessagesRepository();
      bloc = ChatBloc(chatMessagesRepository: mockChatMessagesRepository);
    });

    final now = DateTime.parse('2024-01-01T12:00:00Z');
    final userMessage = ChatMessage(
      id: '1',
      message: 'Hello',
      response: null,
      isUser: true,
      createdAt: now,
    );
    final botMessage = ChatMessage(
      id: '2',
      message: 'Hello',
      response: 'Hi there!',
      isUser: false,
      createdAt: now,
    );

    test('emits [Loaded (user), Loaded (bot)] on send success', () async {
      when(mockChatMessagesRepository.sendMessage('Hello')).thenAnswer((_) async => 'Hi there!');

      final expected = [
        isA<ChatStateLoaded>().having((s) => s.messages.length, 'messages.length', 1),
        isA<ChatStateLoaded>().having((s) => s.messages.length, 'messages.length', 2),
      ];

      expectLater(
        bloc.stream,
        emitsInOrder(expected),
      );

      bloc.add(const ChatEventSend(message: 'Hello'));
    });

    test('emits [Loaded (user), Error (bot error)] on send failure', () async {
      when(mockChatMessagesRepository.sendMessage('Hello')).thenThrow(Exception('fail'));

      final expected = [
        isA<ChatStateLoaded>().having((s) => s.messages.length, 'messages.length', 1),
        isA<ChatStateError>().having((s) => s.messages.length, 'messages.length', 2),
      ];

      expectLater(
        bloc.stream,
        emitsInOrder(expected),
      );

      bloc.add(const ChatEventSend(message: 'Hello'));
    });

    test('emits [Loading, Loaded] on retrieve success', () async {
      final rawMessages = [
        ChatMessage(
          id: '1',
          message: 'Hi',
          response: 'Hello!',
          isUser: true,
          createdAt: now,
        ),
        ChatMessage(
          id: '2',
          message: 'How are you?',
          response: '',
          isUser: true,
          createdAt: now,
        ),
      ];
      when(mockChatMessagesRepository.fetchMessages()).thenAnswer((_) async => rawMessages);

      final expected = [
        isA<ChatStateLoading>(),
        isA<ChatStateLoaded>().having((s) => s.messages.length, 'messages.length', 3),
      ];

      expectLater(
        bloc.stream,
        emitsInOrder(expected),
      );

      bloc.add(const ChatEventRetrieve());
    });

    test('emits [Loading, Error] on retrieve failure', () async {
      when(mockChatMessagesRepository.fetchMessages()).thenThrow(Exception('fail'));

      final expected = [
        isA<ChatStateLoading>(),
        isA<ChatStateError>(),
      ];

      expectLater(
        bloc.stream,
        emitsInOrder(expected),
      );

      bloc.add(const ChatEventRetrieve());
    });
  });
}
