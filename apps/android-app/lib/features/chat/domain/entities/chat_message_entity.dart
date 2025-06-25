class ChatMessage {
  final int id;
  final String message;
  final String? response;
  final bool isUser;
  final DateTime createdAt;

  ChatMessage({
    required this.id,
    required this.message,
    required this.response,
    required this.isUser,
    required this.createdAt,
  });

  factory ChatMessage.fromJson(Map<String, dynamic> json) {
    return ChatMessage(
      id: json['id'],
      message: json['message'],
      response: json['response'],
      isUser: json['is_user'],
      createdAt: DateTime.parse(json['created_at']),
    );
  }
}
