import 'package:android_app/app/data/services/network_service.dart';
import 'package:android_app/features/chat/domain/entities/chat_message_entity.dart';
import 'package:android_app/features/chat/domain/repositories/chat_messages_repository.dart';
import 'package:flutter/foundation.dart';

class ChatMessagesNetworkRepository extends ChatMessagesRepository {
  @override
  Future<List<ChatMessage>> fetchMessages() async {
    final response = await NetworkService().request(
      method: 'GET',
      path: '/api/chat/history',
    );

    if (response.statusCode == 200) {
      if (response.data['messages'] is List) {
        if (kDebugMode) print(response.data['messages']);
        return (response.data['messages'] as List)
            .map((e) => ChatMessage.fromJson(e))
            .toList();
      }
    }
    return <ChatMessage>[];
  }

  @override
  Future<String> sendMessage(String message) async {
    final response = await NetworkService().request(
      method: 'POST',
      path: '/api/chat',
      body: {'message': message},
    );

    if (response.statusCode == 200) {
      return response.data['response'];
    }

    return '';
  }
}
