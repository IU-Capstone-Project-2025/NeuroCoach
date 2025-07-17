abstract class HealthRepository {
  Future<bool> checkToken(String token);
  Future<List<String>> refresh(String refreshToken);
}