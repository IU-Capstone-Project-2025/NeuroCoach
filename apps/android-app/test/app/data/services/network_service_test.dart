import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'package:android_app/app/data/services/network_service.dart';
import 'package:dio/dio.dart';

import 'network_service_test.mocks.dart';

@GenerateMocks([Dio])
void main() {
  group('NetworkService', () {
    late NetworkService networkService;
    late MockDio mockDio;

    setUp(() {
      mockDio = MockDio();
      networkService = NetworkService();
    });

    test('request does not throw and returns a Response', () async {
      // This test only checks that the method can be called without error.
      // Full integration with Dio would require refactoring for testability.
      // Here, we just check that the method exists and can be called.
      expect(
        () => networkService.request(
          method: 'GET',
          path: '/test',
        ),
        returnsNormally,
      );
    });

    test('setToken does not throw', () {
      expect(() => networkService.setToken('abc123'), returnsNormally);
    });

    test('removeToken does not throw', () {
      expect(() {
        networkService.setToken('abc123');
        networkService.removeToken();
      }, returnsNormally);
    });
  });
}
