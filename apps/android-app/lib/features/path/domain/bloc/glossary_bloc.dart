import 'package:android_app/features/path/domain/entities/workout_entity.dart';
import 'package:android_app/features/path/domain/repositories/glossary_repository.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class GlossaryBloc extends Bloc<GlossaryEvent, GlossaryState> {
  GlossaryBloc({required GlossaryRepository glossaryRepository})
    : _glossaryRepository = glossaryRepository,
      super(GlossaryStateInitial()) {
    on<GlossaryEvent>(
        (event, emit) => switch (event) {
          GlossaryEventLoad() => _onLoad(event, emit),
        }
    );
  }

  void _onLoad(GlossaryEventLoad event, Emitter<GlossaryState> emit) async {
    emit(GlossaryStateLoading());

    try {
      final exercises = await _glossaryRepository.getGlossary();
      emit(GlossaryStateLoaded(glossary: exercises));
    } on Exception catch (e, s) {
      if (kDebugMode) print('$e, $s');
      emit(GlossaryStateError(error: e));
    }
  }

  final GlossaryRepository _glossaryRepository;
}

sealed class GlossaryEvent {
  const GlossaryEvent();
}

class GlossaryEventLoad extends GlossaryEvent {
  const GlossaryEventLoad();
}

sealed class GlossaryState {
  const GlossaryState();
}

class GlossaryStateInitial extends GlossaryState {
  const GlossaryStateInitial();
}

class GlossaryStateLoading extends GlossaryState {
  const GlossaryStateLoading();
}

class GlossaryStateLoaded extends GlossaryState {
  final List<SimplifiedExerciseEntity> glossary;

  const GlossaryStateLoaded({required this.glossary});
}

class GlossaryStateError extends GlossaryState {
  final Exception error;

  const GlossaryStateError({required this.error});
}
