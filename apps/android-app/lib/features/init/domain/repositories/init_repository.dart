abstract class InitRepository {

  Future<String> getJWTToken();
  Future<bool> removeJWTToken();
}