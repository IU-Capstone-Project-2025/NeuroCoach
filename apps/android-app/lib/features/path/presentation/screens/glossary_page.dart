import 'package:android_app/constants/app_colors.dart';
import 'package:android_app/constants/app_text_styles.dart';
import 'package:android_app/features/path/domain/bloc/glossary_bloc.dart';
import 'package:auto_route/annotations.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

@RoutePage()
class GlossaryPage extends StatefulWidget {
  const GlossaryPage({super.key});

  @override
  State<StatefulWidget> createState() => _GlossaryPageState();
}

class _GlossaryPageState extends State<GlossaryPage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Glossary'),
        titleTextStyle: AppTextStyles.chatTitle,
        centerTitle: true,
      ),
      body: BlocBuilder<GlossaryBloc, GlossaryState>(
        builder: (context, state) {
          switch (state) {
            case GlossaryStateLoaded():
              if (state.glossary.isEmpty) {
                return Center(
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Text(
                        'There is nothing here...',
                        style: AppTextStyles.textButton,
                      ),
                      const SizedBox(height: 8),
                      TextButton(
                        onPressed: () =>
                            context.read<GlossaryBloc>().add(GlossaryEventLoad()),
                        child: Text('Refresh'),
                      ),
                    ],
                  ),
                );
              }
              return Column(
                children: [
                  Expanded(
                    child: ListView.builder(
                      itemCount: state.glossary.length,
                      padding: const EdgeInsets.all(16),
                      itemBuilder: (context, index) {
                        final exercise = state.glossary[index];
                        return Padding(
                          padding: const EdgeInsets.symmetric(vertical: 6.0),
                          child: Container(
                            decoration: BoxDecoration(
                              color: AppColors.grey.withAlpha(90),
                              borderRadius: BorderRadius.circular(16),
                            ),
                            child: Theme(
                              data: Theme.of(
                                context,
                              ).copyWith(dividerColor: Colors.transparent),
                              child: ExpansionTile(
                                title: Text(
                                  exercise.name,
                                  style: AppTextStyles.textButton.copyWith(
                                    color: AppColors.white,
                                  ),
                                ),
                                shape: RoundedRectangleBorder(
                                  borderRadius: BorderRadius.circular(16.0),
                                  side: const BorderSide(color: Colors.transparent),
                                ),
                                collapsedShape: RoundedRectangleBorder(
                                  borderRadius: BorderRadius.circular(16.0),
                                  side: const BorderSide(color: Colors.transparent),
                                ),
                                iconColor: AppColors.white,
                                collapsedIconColor: AppColors.white,
                                initiallyExpanded: false,
                                children: [
                                  Padding(
                                    padding: const EdgeInsets.all(12.0),
                                    child: Column(
                                      crossAxisAlignment: CrossAxisAlignment.start,
                                      children:
                                          [
                                                Text(
                                                  'Description: ${exercise.description}',
                                                  style: AppTextStyles.textButton,
                                                ),
                                                Image.network(
                                                  exercise.imageUrl,
                                                  scale: 0.5,
                                                ),
                                              ]
                                              .map(
                                                (e) => Padding(
                                                  padding:
                                                      const EdgeInsets.symmetric(
                                                        vertical: 2,
                                                      ),
                                                  child: e,
                                                ),
                                              )
                                              .toList(),
                                    ),
                                  ),
                                ],
                              ),
                            ),
                          ),
                        );
                      },
                    ),
                  ),
                ],
              );
            case GlossaryStateLoading():
              return Center(child: CircularProgressIndicator());
            default:
              return Center(
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Text(
                      'Something went wrong...',
                      style: AppTextStyles.textButton,
                    ),
                    const SizedBox(height: 8),
                    TextButton(
                      onPressed: () =>
                          context.read<GlossaryBloc>().add(GlossaryEventLoad()),
                      child: Text('Refresh'),
                    ),
                  ],
                ),
              );
          }
        },
      ),
    );
  }
}
