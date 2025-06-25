abstract class LoginRepository {
  Future<String> login(String email, String password, bool rememberMe);
  Future<String> signUp(String email, String password, int height, int weight, int age, String goal, List<String> healthIssues, String timeframe, String fitnessLevel, int availableMinutes, bool rememberMe);
}