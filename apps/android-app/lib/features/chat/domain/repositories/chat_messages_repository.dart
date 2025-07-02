import '../entities/chat_message_entity.dart';

abstract class ChatMessagesRepository {
  Future<String> sendMessage(String message);
  Future<List<ChatMessage>> fetchMessages();
}
