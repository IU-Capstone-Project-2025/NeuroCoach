import 'package:android_app/constants/variables.dart';
import 'package:dio/dio.dart';
import 'package:dio_smart_retry/dio_smart_retry.dart';
import 'package:flutter/foundation.dart';

class NetworkService {
  NetworkService._internal();

  static final NetworkService _instance = NetworkService._internal();

  factory NetworkService() => _instance;

  final DioClient _dioClient = DioClient();

  Future<Response<dynamic>> request({
    required String method,
    required String path,
    Map<String, dynamic>? queryParams,
    dynamic body,
    Map<String, dynamic>? headers,
  }) async {
    final options = Options(method: method, headers: headers);

    try {
      return await _dioClient.dio.request(
        path,
        data: body,
        queryParameters: queryParams,
        options: options,
      );
    } on DioException catch (e, s) {
      if (kDebugMode) print ('$e, $s');
      rethrow;
    }
  }

  void setToken(String token) => _dioClient.setToken(token);
  void removeToken() => _dioClient.removeToken();
}

class DioClient {
  final Dio _dio;
  String? _token;

  DioClient._internal()
    : _dio = Dio(
        BaseOptions(
          baseUrl: 'http://fuckyou.com',
          connectTimeout: const Duration(seconds: 10),
          receiveTimeout: const Duration(seconds: 10),
          contentType: 'application/json',
        ),
      ) {
    _dio.interceptors.add(
      RetryInterceptor(
        dio: _dio,
        logPrint: print,
        retries: 3,
        retryDelays: const [
          Duration(seconds: 1),
          Duration(seconds: 2),
          Duration(seconds: 3),
        ],
        retryEvaluator: (error, _) =>
            error.type != DioExceptionType.cancel &&
            error.type != DioExceptionType.badResponse &&
            error.type != DioExceptionType.connectionTimeout,
      ),
    );

    _dio.interceptors.add(
      InterceptorsWrapper(
        onRequest: (options, handler) {
          if (_token != null) {
            options.headers['Authorization'] = 'Bearer $_token';
          }
          return handler.next(options);
        },
      ),
    );
  }

  static final DioClient _instance = DioClient._internal();

  factory DioClient() => _instance;

  Dio get dio => _dio;

  void setToken(String token) {
    _token = token;
  }

  void removeToken() {
    _token = null;
  }
}
