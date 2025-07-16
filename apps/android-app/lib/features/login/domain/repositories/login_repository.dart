abstract class LoginRepository {
  Future<List<String>> login(String email, String password, bool rememberMe);
  Future<List<String>> signUp(String email, String password, int height, int weight, int age, String goal, List<String> healthIssues, String timeframe, String fitnessLevel, int availableMinutes, bool rememberMe);
}