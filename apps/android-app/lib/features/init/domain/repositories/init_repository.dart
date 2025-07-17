abstract class InitRepository {

  Future<String> getJWTToken();
  Future<bool> removeJWTToken();
  Future<String> getRefreshToken();
  Future<bool> removeRefreshToken();
}