import { HomeIcon, SendIcon, MenuIcon, ChevronDownIcon, ChevronUpIcon, ArrowLeftIcon, UserIcon } from "lucide-react";
import React, { useState } from "react";
import { Avatar } from "../../components/ui/avatar";
import { Badge } from "../../components/ui/badge";
import { Button } from "../../components/ui/button";
import { Card, CardContent } from "../../components/ui/card";
import { Input } from "../../components/ui/input";

// Data for numbered steps
const steps = [
  { id: 1, color: "#666eff", textColor: "text-white" },
  { id: 2, color: "#666eff", textColor: "text-white" },
  { id: 3, color: "#FFFFFF", textColor: "text-[#7a81ff]" },
  { id: 4, color: "#363636", textColor: "text-black" },
];

// Workout data for each step
const workoutPlans = {
  1: {
    title: "WORKOUT",
    subtitle: "Training using the linear split methodology for the pectoral muscles.",
    explanation: "Explanation of technique",
    exercises: [
      {
        id: 1,
        name: "Bench press",
        isExpanded: false,
        details: {
          repetitions: "8-12",
          sets: "4",
          timePerSet: "30-40 seconds",
          restBetweenSets: "60-90 seconds",
          technique: {
            startingPosition: "Lie flat on the bench with your feet firmly planted on the ground. Grip the barbell with hands slightly wider than shoulder-width apart.",
            execution: "Lower the bar to your chest in a controlled manner, then press it back up to the starting position.",
            tips: "Keep your core engaged and maintain a slight arch in your back."
          }
        }
      },
      {
        id: 2,
        name: "Dumbbell press at an angle of 45°",
        isExpanded: false,
        details: {
          repetitions: "10-15",
          sets: "3",
          timePerSet: "35-45 seconds",
          restBetweenSets: "45-60 seconds",
          technique: {
            startingPosition: "Set the bench to a 45-degree incline. Hold dumbbells at shoulder level with palms facing forward.",
            execution: "Press the dumbbells up and slightly inward until your arms are fully extended, then lower back to starting position.",
            tips: "Focus on squeezing your chest muscles at the top of the movement."
          }
        }
      },
      {
        id: 3,
        name: "Push-ups with wide arm placement",
        isExpanded: true,
        details: {
          repetitions: "12-15",
          sets: "3",
          timePerSet: "40-50 seconds",
          restBetweenSets: "60-90 seconds",
          technique: {
            startingPosition: "Get into a standard push-up position, but place your hands wider than shoulder-width apart (about 1.5-2 times shoulder width). Fingers slightly turned outward, body in a straight line from head to heels.",
            loweringPhase: "Inhale as you slowly lower your body by bending your elbows. Elbows should go out to the sides at about a 45-70° angle from your torso, not straight out. Lower your chest below elbow level, almost touching the floor. Keep your back and neck straight.",
            tips: "Engage your core and glutes to stabilize your body."
          }
        }
      }
    ]
  },
  2: {
    title: "WORKOUT",
    subtitle: "Training using the linear split methodology for the back muscles.",
    explanation: "Explanation of technique",
    exercises: [
      {
        id: 1,
        name: "Pull-ups",
        isExpanded: false,
        details: {
          repetitions: "6-10",
          sets: "4",
          timePerSet: "25-35 seconds",
          restBetweenSets: "90-120 seconds"
        }
      },
      {
        id: 2,
        name: "Bent-over rows",
        isExpanded: false,
        details: {
          repetitions: "8-12",
          sets: "3",
          timePerSet: "30-40 seconds",
          restBetweenSets: "60-90 seconds"
        }
      },
      {
        id: 3,
        name: "Lat pulldowns",
        isExpanded: false,
        details: {
          repetitions: "10-15",
          sets: "3",
          timePerSet: "35-45 seconds",
          restBetweenSets: "45-60 seconds"
        }
      }
    ]
  },
  3: {
    title: "WORKOUT",
    subtitle: "Training using the linear split methodology for the leg muscles.",
    explanation: "Explanation of technique",
    exercises: [
      {
        id: 1,
        name: "Squats",
        isExpanded: false,
        details: {
          repetitions: "10-15",
          sets: "4",
          timePerSet: "40-50 seconds",
          restBetweenSets: "90-120 seconds"
        }
      },
      {
        id: 2,
        name: "Lunges",
        isExpanded: false,
        details: {
          repetitions: "12-16",
          sets: "3",
          timePerSet: "45-55 seconds",
          restBetweenSets: "60-90 seconds"
        }
      },
      {
        id: 3,
        name: "Calf raises",
        isExpanded: false,
        details: {
          repetitions: "15-20",
          sets: "3",
          timePerSet: "30-40 seconds",
          restBetweenSets: "45-60 seconds"
        }
      }
    ]
  },
  4: {
    title: "WORKOUT",
    subtitle: "Training using the linear split methodology for the shoulder muscles.",
    explanation: "Explanation of technique",
    exercises: [
      {
        id: 1,
        name: "Overhead press",
        isExpanded: false,
        details: {
          repetitions: "8-12",
          sets: "4",
          timePerSet: "30-40 seconds",
          restBetweenSets: "60-90 seconds"
        }
      },
      {
        id: 2,
        name: "Lateral raises",
        isExpanded: false,
        details: {
          repetitions: "12-15",
          sets: "3",
          timePerSet: "35-45 seconds",
          restBetweenSets: "45-60 seconds"
        }
      },
      {
        id: 3,
        name: "Rear delt flyes",
        isExpanded: false,
        details: {
          repetitions: "12-15",
          sets: "3",
          timePerSet: "30-40 seconds",
          restBetweenSets: "45-60 seconds"
        }
      }
    ]
  }
};

// Chat data
const chatMessages = [
  {
    type: "question",
    content:
      "How much water should you drink daily while exercising to maintain balance and peak performance?",
  },
  {
    type: "answer",
    content: `During exercise, it is recommended to drink 500-1000 ml of water per hour, depending on the intensity of the workout and sweating.

General recommendations:

Before exercise: 400-600 ml 2-3 hours before.

During: 150-350 ml every 15-20 minutes.

After: 500-700 ml for every 0.5 kg of weight lost.`,
  },
];

// Profile data
const profileData = {
  nickname: "FitnessPro",
  height: "175 cm",
  weight: "70 kg",
  trainingLevel: "amateur",
  achievements: [
    { level: 1, completed: true, progress: 100 },
    { level: 2, completed: true, progress: 100 },
    { level: 3, completed: false, progress: 80, score: true },
    { level: 4, completed: false, progress: 0 },
    { level: 5, completed: false, progress: 0 },
    { level: 6, completed: false, progress: 0 },
    { level: 7, completed: false, progress: 0 },
  ],
  stats: {
    passedPercentage: 32,
    completedExercises: 17,
    totalExercises: 157,
    caloriesPerDay: [
      { day: "1 DAY", calories: 450 },
      { day: "2 DAY", calories: 520 },
      { day: "3 DAY", calories: 480 },
      { day: "4 DAY", calories: 390 },
      { day: "5 DAY", calories: 550 },
      { day: "6 DAY", calories: 470 },
      { day: "7 DAY", calories: 600 },
    ]
  }
};

export const Frame = (): JSX.Element => {
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);
  const [selectedWorkout, setSelectedWorkout] = useState<number | null>(null);
  const [expandedExercises, setExpandedExercises] = useState<{ [key: number]: boolean }>({ 3: true });
  const [showProfile, setShowProfile] = useState(false);

  const handleStepClick = (stepId: number) => {
    setSelectedWorkout(stepId);
    setExpandedExercises({});
  };

  const handleBackClick = () => {
    setSelectedWorkout(null);
  };

  const handleProfileClick = () => {
    setShowProfile(!showProfile);
  };

  const handleProfileBackClick = () => {
    setShowProfile(false);
  };

  const toggleExercise = (exerciseId: number) => {
    setExpandedExercises(prev => ({
      ...prev,
      [exerciseId]: !prev[exerciseId]
    }));
  };

  // Profile View
  if (showProfile) {
    return (
      <div className="min-h-screen bg-[#1f1f1f] relative overflow-hidden">
        {/* Background image - responsive */}
        <div
          className="absolute inset-0 bg-cover bg-center bg-no-repeat opacity-80"
          style={{
            backgroundImage: "url('/search-photoroom-1.png')"
          }}
        />

        {/* Top navigation bar - responsive */}
        <header className="relative z-20 w-full h-16 md:h-20 flex items-center justify-between px-4 md:px-8 lg:px-[54px]">
          {/* Left side - Home button */}
          <Button
            variant="outline"
            size="icon"
            className="w-8 h-8 md:w-[38px] md:h-[38px] rounded-full bg-[#d9d9d921] border-[0.5px] border-[#5b5b5b] shadow-[inset_-2px_2px_7px_1px_#00000054]"
          >
            <HomeIcon className="h-4 w-4 md:h-5 md:w-5" />
          </Button>

          {/* Center navigation - Training button */}
          <div className="flex items-center space-x-4">
            <Button className="h-[37px] px-6 bg-[#666eff] rounded-[20px] border-[0.5px] border-[#f5f5f563] shadow-[inset_-2px_2px_7px_1px_#00000054]">
              <span className="font-['AR_One_Sans',Helvetica] text-white text-lg md:text-xl">
                Training
              </span>
            </Button>
          </div>

          {/* Right side - User button */}
          <Button
            variant="outline"
            size="icon"
            onClick={handleProfileClick}
            className="w-8 h-8 md:w-[38px] md:h-[38px] rounded-full bg-[#d9d9d921] border-[0.5px] border-[#5b5b5b] shadow-[inset_-2px_2px_7px_1px_#00000054]"
          >
            <UserIcon className="h-4 w-4 md:h-5 md:w-5 text-white" />
          </Button>
        </header>

        <div className="w-full h-px bg-gradient-to-r from-transparent via-[#5b5b5b] to-transparent relative z-10" />

        {/* Profile Content */}
        <div className="relative z-10 p-4 md:p-6 lg:p-8">
          <div className="max-w-7xl mx-auto">
            {/* Header with back button */}
            <div className="flex items-center mb-6">
              <Button
                variant="ghost"
                size="icon"
                onClick={handleProfileBackClick}
                className="mr-4 text-white hover:bg-[#d9d9d921]"
              >
                <ArrowLeftIcon className="h-6 w-6" />
              </Button>
              <h1 className="font-['AR_One_Sans',Helvetica] font-bold text-white text-2xl md:text-4xl">
                YOUR ACHIEVES
              </h1>
            </div>

            <div className="flex flex-col lg:flex-row gap-6 lg:gap-8">
              {/* Left section - Achievements */}
              <div className="flex-1">
                <div className="space-y-4">
                  {profileData.achievements.map((achievement) => (
                    <Card
                      key={achievement.level}
                      className="bg-[#d9d9d98a] rounded-[20px] border-none overflow-hidden"
                    >
                      <CardContent className="p-4 md:p-6">
                        <div className="flex items-center justify-between">
                          <span className="font-['AR_One_Sans',Helvetica] font-bold text-white text-lg md:text-xl">
                            LEVEL {achievement.level}
                          </span>
                          {achievement.score ? (
                            <div className="flex items-center">
                              <span className="font-['AR_One_Sans',Helvetica] text-white text-sm mr-2">
                                Your score
                              </span>
                              <ChevronUpIcon className="h-5 w-5 text-white" />
                            </div>
                          ) : (
                            <ChevronDownIcon className="h-5 w-5 text-white" />
                          )}
                        </div>

                        {achievement.score && (
                          <div className="mt-4">
                            <div className="flex items-center justify-between mb-2">
                              <span className="font-['AR_One_Sans',Helvetica] text-white text-lg">
                                LEVEL 3
                              </span>
                              <span className="font-['AR_One_Sans',Helvetica] text-white text-lg font-bold">
                                {achievement.progress}%
                              </span>
                            </div>
                            <div className="w-full bg-[#5b5b5b] rounded-full h-3">
                              <div
                                className="bg-[#666eff] h-3 rounded-full transition-all duration-300"
                                style={{ width: `${achievement.progress}%` }}
                              />
                            </div>
                          </div>
                        )}
                      </CardContent>
                    </Card>
                  ))}
                </div>
              </div>

              {/* Right section - Profile Info */}
              <div className="lg:max-w-md xl:max-w-lg">
                <Card className="bg-[#d9d9d98a] rounded-[20px] border-none mb-6">
                  <CardContent className="p-4 md:p-6">
                    {/* Profile header */}
                    <div className="flex items-center mb-6">
                      <Avatar className="w-12 h-12 mr-4 bg-[#666eff]">
                        <UserIcon className="h-6 w-6 text-white" />
                      </Avatar>
                      <span className="font-['AR_One_Sans',Helvetica] text-white text-xl font-bold">
                        profile
                      </span>
                    </div>

                    {/* Profile details */}
                    <div className="space-y-4">
                      <div>
                        <span className="font-['AR_One_Sans',Helvetica] text-white text-sm opacity-80">
                          NICKNAME
                        </span>
                        <p className="font-['AR_One_Sans',Helvetica] text-white text-base">
                          {profileData.nickname}
                        </p>
                      </div>

                      <div>
                        <span className="font-['AR_One_Sans',Helvetica] text-white text-sm opacity-80">
                          HEIGHT
                        </span>
                        <p className="font-['AR_One_Sans',Helvetica] text-white text-base">
                          {profileData.height}
                        </p>
                      </div>

                      <div>
                        <span className="font-['AR_One_Sans',Helvetica] text-white text-sm opacity-80">
                          WEIGHT
                        </span>
                        <p className="font-['AR_One_Sans',Helvetica] text-white text-base">
                          {profileData.weight}
                        </p>
                      </div>

                      <div>
                        <span className="font-['AR_One_Sans',Helvetica] text-white text-sm opacity-80">
                          LEVEL OF TRAINING
                        </span>
                        <div className="flex items-center mt-1">
                          <div className="w-2 h-2 bg-[#666eff] rounded-full mr-2"></div>
                          <span className="font-['AR_One_Sans',Helvetica] text-white text-base">
                            {profileData.trainingLevel}
                          </span>
                        </div>
                      </div>
                    </div>
                  </CardContent>
                </Card>

                {/* Progress Stats */}
                <div className="space-y-4">
                  {/* Passed percentage */}
                  <Card className="bg-[#d9d9d98a] rounded-[20px] border-none">
                    <CardContent className="p-4">
                      <div className="text-center mb-2">
                        <span className="font-['AR_One_Sans',Helvetica] text-white text-sm">
                          PASSED {profileData.stats.passedPercentage}% OUT OF 100%
                        </span>
                      </div>
                      <div className="w-full bg-[#5b5b5b] rounded-full h-2">
                        <div
                          className="bg-[#666eff] h-2 rounded-full"
                          style={{ width: `${profileData.stats.passedPercentage}%` }}
                        />
                      </div>
                    </CardContent>
                  </Card>

                  {/* Completed exercises */}
                  <Card className="bg-[#d9d9d98a] rounded-[20px] border-none">
                    <CardContent className="p-4">
                      <div className="text-center mb-2">
                        <span className="font-['AR_One_Sans',Helvetica] text-white text-sm">
                          COMPLETED {profileData.stats.completedExercises} EXERCISES OUT OF {profileData.stats.totalExercises}
                        </span>
                      </div>
                      <div className="w-full bg-[#5b5b5b] rounded-full h-2">
                        <div
                          className="bg-[#666eff] h-2 rounded-full"
                          style={{ width: `${(profileData.stats.completedExercises / profileData.stats.totalExercises) * 100}%` }}
                        />
                      </div>
                    </CardContent>
                  </Card>

                  {/* Calories chart */}
                  <Card className="bg-[#d9d9d98a] rounded-[20px] border-none">
                    <CardContent className="p-4">
                      <div className="text-center mb-4">
                        <span className="font-['AR_One_Sans',Helvetica] text-white text-sm">
                          HOW MANY CALORIES BURNED PER DAY
                        </span>
                      </div>

                      <div className="flex items-end justify-between h-32 mb-2">
                        {profileData.stats.caloriesPerDay.map((day, index) => (
                          <div key={index} className="flex flex-col items-center flex-1">
                            <div
                              className="bg-[#666eff] w-4 rounded-t-sm mb-1"
                              style={{
                                height: `${(day.calories / 600) * 100}%`,
                                minHeight: '20px'
                              }}
                            />
                          </div>
                        ))}
                      </div>

                      <div className="flex justify-between text-xs">
                        {profileData.stats.caloriesPerDay.map((day, index) => (
                          <span key={index} className="font-['AR_One_Sans',Helvetica] text-white text-xs">
                            {day.day}
                          </span>
                        ))}
                      </div>

                      {/* Y-axis labels */}
                      <div className="absolute right-2 top-4 flex flex-col justify-between h-24 text-xs">
                        <span className="font-['AR_One_Sans',Helvetica] text-white text-xs">600</span>
                        <span className="font-['AR_One_Sans',Helvetica] text-white text-xs">300</span>
                        <span className="font-['AR_One_Sans',Helvetica] text-white text-xs">0</span>
                      </div>
                    </CardContent>
                  </Card>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    );
  }

  if (selectedWorkout) {
    const workout = workoutPlans[selectedWorkout as keyof typeof workoutPlans];

    return (
      <div className="min-h-screen bg-[#1f1f1f] relative overflow-hidden">
        {/* Background image - responsive */}
        <div
          className="absolute inset-0 bg-cover bg-center bg-no-repeat opacity-80"
          style={{
            backgroundImage: "url('/search-photoroom-1.png')"
          }}
        />

        {/* Top navigation bar - responsive */}
        <header className="relative z-20 w-full h-16 md:h-20 flex items-center justify-between px-4 md:px-8 lg:px-[54px]">
          {/* Left side - Home button */}
          <Button
            variant="outline"
            size="icon"
            className="w-8 h-8 md:w-[38px] md:h-[38px] rounded-full bg-[#d9d9d921] border-[0.5px] border-[#5b5b5b] shadow-[inset_-2px_2px_7px_1px_#00000054]"
          >
            <HomeIcon className="h-4 w-4 md:h-5 md:w-5" />
          </Button>

          {/* Center navigation - Training button */}
          <div className="flex items-center space-x-4">
            <Button className="h-[37px] px-6 bg-[#666eff] rounded-[20px] border-[0.5px] border-[#f5f5f563] shadow-[inset_-2px_2px_7px_1px_#00000054]">
              <span className="font-['AR_One_Sans',Helvetica] text-white text-lg md:text-xl">
                Training
              </span>
            </Button>
          </div>

          {/* Right side - User button */}
          <Button
            variant="outline"
            size="icon"
            onClick={handleProfileClick}
            className="w-8 h-8 md:w-[38px] md:h-[38px] rounded-full bg-[#d9d9d921] border-[0.5px] border-[#5b5b5b] shadow-[inset_-2px_2px_7px_1px_#00000054]"
          >
            <img
              className="w-5 h-5 md:w-[25px] md:h-[25px] object-cover"
              alt="User icon"
              src="/icons8---------90-11.png"
            />
          </Button>
        </header>

        <div className="w-full h-px bg-gradient-to-r from-transparent via-[#5b5b5b] to-transparent relative z-10" />

        {/* Workout Plan Content */}
        <div className="relative z-10 p-4 md:p-6 lg:p-8">
          <div className="max-w-6xl mx-auto">
            {/* Header with back button */}
            <div className="flex items-center mb-6">
              <Button
                variant="ghost"
                size="icon"
                onClick={handleBackClick}
                className="mr-4 text-white hover:bg-[#d9d9d921]"
              >
                <ArrowLeftIcon className="h-6 w-6" />
              </Button>
              <h1 className="font-['AR_One_Sans',Helvetica] font-bold text-white text-2xl md:text-4xl">
                {workout.title}
              </h1>
            </div>

            {/* Workout description */}
            <div className="mb-6">
              <p className="font-['AR_One_Sans',Helvetica] text-white text-base md:text-lg mb-4">
                {workout.subtitle}
              </p>
              <p className="font-['AR_One_Sans',Helvetica] text-white text-sm md:text-base opacity-80">
                {workout.explanation}
              </p>
            </div>

            {/* Exercise list */}
            <div className="space-y-4">
              {workout.exercises.map((exercise) => (
                <Card
                  key={exercise.id}
                  className="bg-[#d9d9d98a] rounded-[20px] border-none overflow-hidden"
                >
                  <CardContent className="p-0">
                    {/* Exercise header */}
                    <div
                      className="flex items-center justify-between p-4 md:p-6 cursor-pointer hover:bg-[#d9d9d9aa] transition-colors"
                      onClick={() => toggleExercise(exercise.id)}
                    >
                      <div className="flex items-center">
                        <span className="font-['AR_One_Sans',Helvetica] font-bold text-white text-lg md:text-xl mr-4">
                          {exercise.id}.
                        </span>
                        <span className="font-['AR_One_Sans',Helvetica] text-white text-base md:text-lg">
                          {exercise.name}
                        </span>
                      </div>
                      {expandedExercises[exercise.id] ? (
                        <ChevronUpIcon className="h-5 w-5 text-white" />
                      ) : (
                        <ChevronDownIcon className="h-5 w-5 text-white" />
                      )}
                    </div>

                    {/* Exercise details */}
                    {expandedExercises[exercise.id] && (
                      <div className="px-4 md:px-6 pb-4 md:pb-6 bg-[#f0f0f0aa] text-black">
                        <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-6">
                          <div>
                            <p className="font-['AR_One_Sans',Helvetica] text-sm">
                              <strong>Repetitions:</strong> {exercise.details.repetitions}
                            </p>
                            <p className="font-['AR_One_Sans',Helvetica] text-sm">
                              <strong>Sets:</strong> {exercise.details.sets}
                            </p>
                          </div>
                          <div>
                            <p className="font-['AR_One_Sans',Helvetica] text-sm">
                              <strong>Time per set:</strong> {exercise.details.timePerSet}
                            </p>
                            <p className="font-['AR_One_Sans',Helvetica] text-sm">
                              <strong>Rest between sets:</strong> {exercise.details.restBetweenSets}
                            </p>
                          </div>
                        </div>

                        {exercise.details.technique && (
                          <div>
                            <h4 className="font-['AR_One_Sans',Helvetica] font-bold text-base mb-3">
                              Exercise Technique
                            </h4>

                            {exercise.details.technique.startingPosition && (
                              <div className="mb-4">
                                <p className="font-['AR_One_Sans',Helvetica] font-semibold text-sm mb-1">
                                  1. Starting Position:
                                </p>
                                <p className="font-['AR_One_Sans',Helvetica] text-sm leading-relaxed">
                                  — {exercise.details.technique.startingPosition}
                                </p>
                                {exercise.details.technique.execution && (
                                  <p className="font-['AR_One_Sans',Helvetica] text-sm leading-relaxed mt-2">
                                    — {exercise.details.technique.execution}
                                  </p>
                                )}
                                {exercise.details.technique.tips && (
                                  <p className="font-['AR_One_Sans',Helvetica] text-sm leading-relaxed mt-2">
                                    — {exercise.details.technique.tips}
                                  </p>
                                )}
                              </div>
                            )}

                            {exercise.details.technique.loweringPhase && (
                              <div className="mb-4">
                                <p className="font-['AR_One_Sans',Helvetica] font-semibold text-sm mb-1">
                                  2. Lowering Phase:
                                </p>
                                <p className="font-['AR_One_Sans',Helvetica] text-sm leading-relaxed">
                                  — {exercise.details.technique.loweringPhase}
                                </p>
                              </div>
                            )}
                          </div>
                        )}
                      </div>
                    )}
                  </CardContent>
                </Card>
              ))}
            </div>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-[#1f1f1f] relative overflow-hidden">
      {/* Background image - responsive */}
      <div
        className="absolute inset-0 bg-cover bg-center bg-no-repeat opacity-80"
        style={{
          backgroundImage: "url('/search-photoroom-1.png')"
        }}
      />

      {/* Top navigation bar - responsive */}
      <header className="relative z-20 w-full h-16 md:h-20 flex items-center justify-between px-4 md:px-8 lg:px-[54px]">
        {/* Left side - Home button */}
        <Button
          variant="outline"
          size="icon"
          className="w-8 h-8 md:w-[38px] md:h-[38px] rounded-full bg-[#d9d9d921] border-[0.5px] border-[#5b5b5b] shadow-[inset_-2px_2px_7px_1px_#00000054]"
        >
          <HomeIcon className="h-4 w-4 md:h-5 md:w-5" />
        </Button>

        {/* Center navigation - hidden on mobile, shown on tablet+ */}
        <div className="hidden md:flex space-x-4">
          <Button
            variant="outline"
            className="h-[37px] w-[81px] bg-[#d9d9d921] rounded-[20px] border-[0.5px] border-[#5b5b5b] shadow-[inset_-2px_2px_7px_1px_#00000054]"
          >
            <img
              className="w-[22px] h-[22px] object-cover"
              alt="HomeIcon icon"
              src="/icons8---------384-13.png"
            />
          </Button>

          <Button className="h-[37px] w-[145px] bg-[#666eff] rounded-[20px] border-[0.5px] border-[#f5f5f563] shadow-[inset_-2px_2px_7px_1px_#00000054]">
            <span className="font-['AR_One_Sans',Helvetica] text-white text-lg md:text-xl">
              Training
            </span>
          </Button>
        </div>

        {/* Right side - User button and mobile menu */}
        <div className="flex items-center space-x-2">
          <Button
            variant="outline"
            size="icon"
            onClick={handleProfileClick}
            className="w-8 h-8 md:w-[38px] md:h-[38px] rounded-full bg-[#d9d9d921] border-[0.5px] border-[#5b5b5b] shadow-[inset_-2px_2px_7px_1px_#00000054]"
          >
            <img
              className="w-5 h-5 md:w-[25px] md:h-[25px] object-cover"
              alt="User icon"
              src="/icons8---------90-11.png"
            />
          </Button>

          {/* Mobile menu button */}
          <Button
            variant="outline"
            size="icon"
            className="w-8 h-8 md:hidden rounded-full bg-[#d9d9d921] border-[0.5px] border-[#5b5b5b] shadow-[inset_-2px_2px_7px_1px_#00000054]"
            onClick={() => setIsMobileMenuOpen(!isMobileMenuOpen)}
          >
            <MenuIcon className="h-4 w-4" />
          </Button>
        </div>
      </header>

      {/* Mobile menu overlay */}
      {isMobileMenuOpen && (
        <div className="fixed inset-0 z-30 bg-black bg-opacity-50 md:hidden">
          <div className="absolute top-16 right-4 bg-[#1f1f1f] rounded-lg p-4 shadow-lg">
            <div className="flex flex-col space-y-2">
              <Button
                variant="outline"
                className="h-[37px] w-full bg-[#d9d9d921] rounded-[20px] border-[0.5px] border-[#5b5b5b] shadow-[inset_-2px_2px_7px_1px_#00000054]"
                onClick={() => setIsMobileMenuOpen(false)}
              >
                <img
                  className="w-[22px] h-[22px] object-cover mr-2"
                  alt="HomeIcon icon"
                  src="/icons8---------384-13.png"
                />
                <span className="text-white">Home</span>
              </Button>
              <Button
                className="h-[37px] w-full bg-[#666eff] rounded-[20px] border-[0.5px] border-[#f5f5f563] shadow-[inset_-2px_2px_7px_1px_#00000054]"
                onClick={() => setIsMobileMenuOpen(false)}
              >
                <span className="font-['AR_One_Sans',Helvetica] text-white text-lg">
                  Training
                </span>
              </Button>
            </div>
          </div>
        </div>
      )}

      <div className="w-full h-px bg-gradient-to-r from-transparent via-[#5b5b5b] to-transparent relative z-10" />

      {/* Main content area - responsive layout */}
      <div className="relative z-10 flex flex-col lg:flex-row min-h-[calc(100vh-4rem)] md:min-h-[calc(100vh-5rem)] p-4 md:p-6 lg:p-8 gap-4 lg:gap-8">

        {/* Left section - Steps and Coach */}
        <div className="flex-1 lg:max-w-md xl:max-w-lg">
          {/* Numbered steps section - responsive grid */}
          <section className="mb-6 lg:mb-8">
            <div className="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-2 gap-3 md:gap-4">
              {steps.map((step, index) => (
                <div
                  key={step.id}
                  className="relative bg-[#1f1f1fc9] rounded-[15px] shadow-[inset_-2px_2px_7px_1px_#00000054] backdrop-blur-[17.95px] backdrop-brightness-[100%] p-4 md:p-6 aspect-square flex flex-col items-center justify-center cursor-pointer hover:bg-[#2f2f2fc9] transition-colors"
                  onClick={() => handleStepClick(step.id)}
                >
                  <div className="absolute inset-4 bg-neutral-900 rounded-[15px] shadow-[inset_-2px_2px_7px_1px_#00000054] blur-[9.55px]" />

                  <Card
                    className="relative w-16 h-16 md:w-20 md:h-20 lg:w-24 lg:h-24 rounded-[15px] border-[0.5px] border-solid border-[#f5f5f563] shadow-[inset_-2px_2px_7px_1px_#00000054] hover:scale-105 transition-transform"
                    style={{ backgroundColor: step.color }}
                  >
                    <CardContent className="flex items-center justify-center h-full p-0">
                      <span
                        className={`font-['Zen_Dots',Helvetica] font-normal ${step.textColor} text-2xl md:text-3xl lg:text-4xl text-center`}
                      >
                        {step.id}
                      </span>
                    </CardContent>
                  </Card>
                </div>
              ))}
            </div>
          </section>

          {/* Coach avatar and speech bubble - responsive */}
          <div className="relative flex justify-center lg:justify-start">
            <div className="relative w-48 h-72 md:w-56 md:h-80 lg:w-[261px] lg:h-[405px]">
              <img
                className="w-full h-full object-cover"
                alt="Coach robot"
                src="/-------------6--photoroom-9.png"
              />

              <div className="absolute -top-2 -left-4 md:-left-6">
                <div className="relative">
                  <div className="bg-[#d9d9d9b8] rounded-[30px] px-4 py-2 md:px-6 md:py-3">
                    <span className="font-['AR_One_Sans',Helvetica] font-normal text-white text-sm md:text-lg lg:text-xl whitespace-nowrap">
                      Keep it up !
                    </span>
                  </div>
                  <div className="absolute top-full left-6 md:left-8">
                    <img
                      className="w-4 h-3 md:w-[21px] md:h-[18px]"
                      alt="Speech bubble pointer"
                      src="/polygon-11.svg"
                    />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        {/* Right section - Chat */}
        <div className="flex-1 lg:max-w-2xl xl:max-w-4xl">
          <Card className="h-full min-h-[500px] lg:min-h-[600px] xl:min-h-[700px] bg-[#d9d9d98a] rounded-[20px] md:rounded-[40px] lg:rounded-[77px] border-none">
            <CardContent className="p-0 h-full flex flex-col">
              {/* Chat header */}
              <div className="text-center pt-4 md:pt-6">
                <h2 className="font-['Anta',Helvetica] font-normal text-white text-xl md:text-2xl">
                  coach
                </h2>
              </div>

              {/* Chat messages */}
              <div className="flex-1 p-4 md:p-6 lg:p-8 overflow-y-auto flex flex-col space-y-4">
                {/* Question message */}
                <div className="flex justify-end">
                  <div className="flex items-start max-w-[85%] md:max-w-[70%]">
                    <Card className="bg-[#9499ff] rounded-[20px] md:rounded-[25px] border-[0.5px] border-solid border-[#f5f5f563] shadow-[inset_-2px_2px_7px_1px_#00000012] mr-2">
                      <CardContent className="p-3 md:p-4">
                        <p className="font-['AR_One_Sans',Helvetica] font-normal text-white text-sm md:text-base">
                          {chatMessages[0].content}
                        </p>
                      </CardContent>
                    </Card>
                    <Avatar className="w-8 h-8 md:w-[43px] md:h-[43px] bg-[#d9d9d9] rounded-full flex-shrink-0">
                      <img
                        className="w-5 h-5 md:w-[30px] md:h-[30px] object-cover"
                        alt="User icon"
                        src="/icons8---------90-11.png"
                      />
                    </Avatar>
                  </div>
                </div>

                {/* Answer message */}
                <div className="flex">
                  <div className="flex items-start max-w-[85%] md:max-w-[80%]">
                    <Avatar className="w-8 h-8 md:w-11 md:h-11 mr-2 flex-shrink-0">
                      <img
                        className="w-full h-full"
                        alt="Coach avatar"
                        src="/mask-group.png"
                      />
                    </Avatar>
                    <Card className="bg-[#303030] rounded-[20px] md:rounded-[30px] border-[0.5px] border-solid border-[#00000063] shadow-[inset_-2px_2px_7px_1px_#00000012]">
                      <CardContent className="p-4 md:p-6">
                        <p className="font-['AR_One_Sans',Helvetica] font-normal text-white text-sm md:text-base whitespace-pre-line">
                          {chatMessages[1].content}
                        </p>
                      </CardContent>
                    </Card>
                  </div>
                </div>
              </div>

              {/* Message input */}
              <div className="p-3 md:p-4 mt-auto">
                <Card className="w-full bg-[#d9d9d98a] rounded-[20px] md:rounded-[35px] border-none">
                  <CardContent className="p-3 md:p-4 flex items-center">
                    <Badge
                      className="bg-transparent flex items-center mr-2 border-none"
                      variant="outline"
                    >
                      <span className="font-['AR_One_Sans',Helvetica] font-normal text-white text-lg md:text-2xl">
                        Message
                      </span>
                      <div className="w-px h-6 md:h-[30px] bg-[#565656] ml-2"></div>
                    </Badge>
                    <Input
                      className="flex-1 bg-transparent border-none text-white focus-visible:ring-0 focus-visible:ring-offset-0 text-sm md:text-base"
                      placeholder="Type your message..."
                    />
                    <Button size="icon" className="w-8 h-8 md:w-10 md:h-[41px] ml-2 flex-shrink-0">
                      <SendIcon className="h-4 w-4 md:h-6 md:w-6" />
                    </Button>
                  </CardContent>
                </Card>
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
};