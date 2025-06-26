import 'package:android_app/features/settings/domain/repositories/profile_repository.dart';
import 'package:android_app/features/settings/domain/user_profile_entity.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class ProfileBloc extends Bloc<ProfileEvent, ProfileState> {
  ProfileBloc({required ProfileRepository profileRepository})
      : _profileRepository = profileRepository,
        super(const ProfileStateInitial()) {
    on<ProfileEvent>(
            (event, emit) =>
        switch (event) {
          ProfileEventLoad() => _onLoad(event, emit),
        }
    );
  }

  Future<void> _onLoad(ProfileEventLoad event,
      Emitter<ProfileState> emit) async {
    emit(ProfileStateLoading());
    try {
      final profile = await _profileRepository.getUserProfile();
      if (profile != null) {
        emit(ProfileStateLoaded(userProfile: profile));
      } else {
        emit(ProfileStateError());
      }
    } on Object catch(e, s) {
      if (kDebugMode) print('$e, $s');
      emit(ProfileStateError());
    }
  }

  final ProfileRepository _profileRepository;
}

sealed class ProfileEvent {
  const ProfileEvent();
}

class ProfileEventLoad extends ProfileEvent {
  const ProfileEventLoad();
}

sealed class ProfileState {
  const ProfileState();
}

class ProfileStateInitial extends ProfileState {
  const ProfileStateInitial();
}

class ProfileStateLoading extends ProfileState {
  const ProfileStateLoading();
}

class ProfileStateLoaded extends ProfileState {
  const ProfileStateLoaded({required this.userProfile});

  final UserProfile userProfile;
}

class ProfileStateError extends ProfileState {
  const ProfileStateError();
}
