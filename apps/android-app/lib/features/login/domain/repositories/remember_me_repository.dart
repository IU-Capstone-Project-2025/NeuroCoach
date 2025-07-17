abstract class RememberMeRepository {
  Future<void> rememberUser({required String jwtToken, required String refreshToken, required String email});
}