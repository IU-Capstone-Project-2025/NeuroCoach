abstract class RememberMeRepository {
  Future<void> rememberUser({required String jwtToken, required String email});
}