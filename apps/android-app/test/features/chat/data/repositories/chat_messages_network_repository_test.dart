import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'package:android_app/features/chat/data/repositories/chat_messages_network_repository.dart';
import 'package:android_app/app/data/services/network_service.dart';
import 'package:dio/dio.dart';
import 'package:android_app/features/chat/domain/entities/chat_message_entity.dart';

import 'chat_messages_network_repository_test.mocks.dart';

@GenerateMocks([NetworkService])
void main() {
  group('ChatMessagesNetworkRepository', () {
    late MockNetworkService mockNetworkService;
    late ChatMessagesNetworkRepository repository;

    setUp(() {
      mockNetworkService = MockNetworkService();
      repository = ChatMessagesNetworkRepository(networkService: mockNetworkService);
    });

    test('fetchMessages returns parsed messages on 200', () async {
      final now = DateTime.parse('2024-01-01T12:00:00Z');
      final response = Response(
        requestOptions: RequestOptions(path: ''),
        statusCode: 200,
        data: {
          'messages': [
            {
              'id': '1',
              'message': 'Hi',
              'response': 'Hello!',
              'is_user': true,
              'created_at': now.toIso8601String(),
            },
          ],
        },
      );
      when(mockNetworkService.request(method: 'GET', path: '/api/chat/history')).thenAnswer((_) async => response);

      final result = await repository.fetchMessages();
      expect(result, isA<List<ChatMessage>>());
      expect(result.length, 1);
      expect(result.first.message, 'Hi');
    });

    test('fetchMessages returns empty list on 200 with no messages', () async {
      final response = Response(
        requestOptions: RequestOptions(path: ''),
        statusCode: 200,
        data: {'messages': []},
      );
      when(mockNetworkService.request(method: 'GET', path: '/api/chat/history')).thenAnswer((_) async => response);

      final result = await repository.fetchMessages();
      expect(result, isEmpty);
    });

    test('fetchMessages returns empty list on non-200', () async {
      final response = Response(
        requestOptions: RequestOptions(path: ''),
        statusCode: 500,
        data: {},
      );
      when(mockNetworkService.request(method: 'GET', path: '/api/chat/history')).thenAnswer((_) async => response);

      final result = await repository.fetchMessages();
      expect(result, isEmpty);
    });

    test('sendMessage returns response string on 200', () async {
      final response = Response(
        requestOptions: RequestOptions(path: ''),
        statusCode: 200,
        data: {'response': 'Hi there!'},
      );
      when(mockNetworkService.request(
        method: 'POST',
        path: '/api/chat',
        body: {'message': 'Hello'},
      )).thenAnswer((_) async => response);

      final result = await repository.sendMessage('Hello');
      expect(result, 'Hi there!');
    });

    test('sendMessage returns empty string on non-200', () async {
      final response = Response(
        requestOptions: RequestOptions(path: ''),
        statusCode: 500,
        data: {},
      );
      when(mockNetworkService.request(
        method: 'POST',
        path: '/api/chat',
        body: {'message': 'Hello'},
      )).thenAnswer((_) async => response);

      final result = await repository.sendMessage('Hello');
      expect(result, '');
    });
  });
}
