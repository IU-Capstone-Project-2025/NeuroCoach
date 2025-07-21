import React from 'react';
import { ArrowRight, Smartphone, Zap, Target, BarChart3, Users, Star, Download } from 'lucide-react';
import { Button } from '../../components/ui/button';
import { Card, CardContent } from '../../components/ui/card';

interface LandingProps {
  onNavigateToTraining: () => void;
}

export const Landing = ({ onNavigateToTraining }: LandingProps): JSX.Element => {
  const features = [
    {
      icon: <Target className="w-8 h-8 text-[#666eff]" />,
      title: "Personalized Workouts",
      description: "AI-powered training plans tailored to your fitness level and goals"
    },
    {
      icon: <BarChart3 className="w-8 h-8 text-[#666eff]" />,
      title: "Progress Tracking",
      description: "Monitor your achievements, calories burned, and workout completion"
    },
    {
      icon: <Zap className="w-8 h-8 text-[#666eff]" />,
      title: "Smart Coach",
      description: "Get real-time guidance and motivation from your virtual fitness coach"
    },
    {
      icon: <Users className="w-8 h-8 text-[#666eff]" />,
      title: "Community Support",
      description: "Connect with other fitness enthusiasts and share your journey"
    }
  ];

  const testimonials = [
    {
      name: "Sarah Johnson",
      rating: 5,
      comment: "This app completely transformed my fitness routine. The AI coach is incredibly motivating!"
    },
    {
      name: "Mike Chen",
      rating: 5,
      comment: "Love the personalized workouts and progress tracking. Best fitness app I've used."
    },
    {
      name: "Emma Davis",
      rating: 5,
      comment: "The technique explanations are so detailed. I finally learned proper form!"
    }
  ];

  const handleDownload = () => {
    // In a real app, this would trigger the actual download
    alert('Download will start soon! Available on App Store and Google Play.');
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-[#1f1f1f] via-[#2a2a2a] to-[#1f1f1f]">
      {/* Navigation */}
      <nav className="fixed top-0 w-full z-50 bg-[#1f1f1f]/80 backdrop-blur-md border-b border-[#5b5b5b]/30">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center">
              <div className="w-8 h-8 bg-[#666eff] rounded-lg flex items-center justify-center mr-3">
                <Zap className="w-5 h-5 text-white" />
              </div>
              <span className="font-['AR_One_Sans',Helvetica] text-white text-xl font-bold">
                NeuroCoach
              </span>
            </div>
            <Button
              onClick={onNavigateToTraining}
              className="bg-[#666eff] hover:bg-[#5555ee] text-white px-6 py-2 rounded-full"
            >
              Try Now
            </Button>
          </div>
        </div>
      </nav>

      {/* Hero Section */}
      <section className="pt-24 pb-16 px-4 sm:px-6 lg:px-8">
        <div className="max-w-7xl mx-auto">
          <div className="text-center mb-16">
            <h1 className="font-['AR_One_Sans',Helvetica] text-4xl md:text-6xl lg:text-7xl font-bold text-white mb-6">
              Your Personal
              <span className="text-[#666eff] block">AI Fitness Coach</span>
            </h1>
            <p className="font-['AR_One_Sans',Helvetica] text-xl md:text-2xl text-gray-300 mb-8 max-w-3xl mx-auto">
              Transform your fitness journey with personalized workouts, real-time coaching,
              and comprehensive progress tracking. All powered by advanced AI technology.
            </p>
            <div className="flex flex-col sm:flex-row gap-4 justify-center">
              <Button
                onClick={onNavigateToTraining}
                className="bg-[#666eff] hover:bg-[#5555ee] text-white px-8 py-4 text-lg rounded-full flex items-center"
              >
                Get Started Free
                <ArrowRight className="ml-2 w-5 h-5" />
              </Button>
              <Button
                variant="outline"
                className="border-[#666eff] text-[#666eff] hover:bg-[#666eff] hover:text-white px-8 py-4 text-lg rounded-full"
              >
                Watch Demo
              </Button>
            </div>
          </div>

          {/* App Preview */}
          <div className="relative max-w-4xl mx-auto">
            <div className="bg-gradient-to-r from-[#666eff]/20 to-[#9499ff]/20 rounded-3xl p-8 backdrop-blur-sm border border-[#666eff]/30">
              <div className="bg-[#1f1f1f] rounded-2xl p-6 shadow-2xl">
                <div className="flex items-center mb-4">
                  <div className="flex space-x-2">
                    <div className="w-3 h-3 bg-red-500 rounded-full"></div>
                    <div className="w-3 h-3 bg-yellow-500 rounded-full"></div>
                    <div className="w-3 h-3 bg-green-500 rounded-full"></div>
                  </div>
                  <span className="ml-4 text-gray-400 text-sm">FitCoach AI - Training Interface</span>
                </div>
                <div className="aspect-video bg-gradient-to-br from-[#2a2a2a] to-[#1f1f1f] rounded-lg flex items-center justify-center">
                  <div className="text-center">
                    <Smartphone className="w-16 h-16 text-[#666eff] mx-auto mb-4" />
                    <p className="text-white text-lg">Interactive Training Interface</p>
                    <p className="text-gray-400">Personalized workouts with AI coaching</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-20 px-4 sm:px-6 lg:px-8">
        <div className="max-w-7xl mx-auto">
          <div className="text-center mb-16">
            <h2 className="font-['AR_One_Sans',Helvetica] text-3xl md:text-5xl font-bold text-white mb-6">
              Why Choose FitCoach AI?
            </h2>
            <p className="font-['AR_One_Sans',Helvetica] text-xl text-gray-300 max-w-3xl mx-auto">
              Experience the future of fitness with cutting-edge AI technology designed to maximize your results
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
            {features.map((feature, index) => (
              <Card key={index} className="bg-[#2a2a2a]/50 border-[#5b5b5b]/30 backdrop-blur-sm hover:bg-[#2a2a2a]/70 transition-all duration-300">
                <CardContent className="p-6 text-center">
                  <div className="mb-4 flex justify-center">
                    {feature.icon}
                  </div>
                  <h3 className="font-['AR_One_Sans',Helvetica] text-xl font-bold text-white mb-3">
                    {feature.title}
                  </h3>
                  <p className="text-gray-300">
                    {feature.description}
                  </p>
                </CardContent>
              </Card>
            ))}
          </div>
        </div>
      </section>

      {/* How It Works Section */}
      <section className="py-20 px-4 sm:px-6 lg:px-8 bg-[#2a2a2a]/30">
        <div className="max-w-7xl mx-auto">
          <div className="text-center mb-16">
            <h2 className="font-['AR_One_Sans',Helvetica] text-3xl md:text-5xl font-bold text-white mb-6">
              How It Works
            </h2>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            <div className="text-center">
              <div className="w-16 h-16 bg-[#666eff] rounded-full flex items-center justify-center mx-auto mb-6">
                <span className="text-white text-2xl font-bold">1</span>
              </div>
              <h3 className="font-['AR_One_Sans',Helvetica] text-xl font-bold text-white mb-4">
                Set Your Goals
              </h3>
              <p className="text-gray-300">
                Tell us about your fitness level, goals, and preferences to create your personalized plan
              </p>
            </div>

            <div className="text-center">
              <div className="w-16 h-16 bg-[#666eff] rounded-full flex items-center justify-center mx-auto mb-6">
                <span className="text-white text-2xl font-bold">2</span>
              </div>
              <h3 className="font-['AR_One_Sans',Helvetica] text-xl font-bold text-white mb-4">
                Train with AI Coach
              </h3>
              <p className="text-gray-300">
                Follow guided workouts with real-time feedback and technique corrections from your AI coach
              </p>
            </div>

            <div className="text-center">
              <div className="w-16 h-16 bg-[#666eff] rounded-full flex items-center justify-center mx-auto mb-6">
                <span className="text-white text-2xl font-bold">3</span>
              </div>
              <h3 className="font-['AR_One_Sans',Helvetica] text-xl font-bold text-white mb-4">
                Track Progress
              </h3>
              <p className="text-gray-300">
                Monitor your achievements, calories burned, and unlock new levels as you progress
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* Testimonials Section */}
      <section className="py-20 px-4 sm:px-6 lg:px-8">
        <div className="max-w-7xl mx-auto">
          <div className="text-center mb-16">
            <h2 className="font-['AR_One_Sans',Helvetica] text-3xl md:text-5xl font-bold text-white mb-6">
              What Users Say
            </h2>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            {testimonials.map((testimonial, index) => (
              <Card key={index} className="bg-[#2a2a2a]/50 border-[#5b5b5b]/30 backdrop-blur-sm">
                <CardContent className="p-6">
                  <div className="flex mb-4">
                    {[...Array(testimonial.rating)].map((_, i) => (
                      <Star key={i} className="w-5 h-5 text-yellow-400 fill-current" />
                    ))}
                  </div>
                  <p className="text-gray-300 mb-4 italic">
                    "{testimonial.comment}"
                  </p>
                  <p className="text-white font-semibold">
                    - {testimonial.name}
                  </p>
                </CardContent>
              </Card>
            ))}
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="py-20 px-4 sm:px-6 lg:px-8 bg-gradient-to-r from-[#666eff]/20 to-[#9499ff]/20">
        <div className="max-w-4xl mx-auto text-center">
          <h2 className="font-['AR_One_Sans',Helvetica] text-3xl md:text-5xl font-bold text-white mb-6">
            Ready to Transform Your Fitness?
          </h2>
          <p className="font-['AR_One_Sans',Helvetica] text-xl text-gray-300 mb-8">
            Join thousands of users who have already achieved their fitness goals with FitCoach AI
          </p>

          <div className="flex flex-col sm:flex-row gap-4 justify-center mb-12">
            <Button
              onClick={handleDownload}
              className="bg-[#666eff] hover:bg-[#5555ee] text-white px-8 py-4 text-lg rounded-full flex items-center justify-center"
            >
              <Download className="mr-2 w-5 h-5" />
              Download App
            </Button>
            <Button
              onClick={onNavigateToTraining}
              variant="outline"
              className="border-[#666eff] text-[#666eff] hover:bg-[#666eff] hover:text-white px-8 py-4 text-lg rounded-full"
            >
              Try Web Version
            </Button>
          </div>

          <div className="flex justify-center space-x-8">
            <div className="text-center">
              <div className="text-3xl font-bold text-[#666eff]">10K+</div>
              <div className="text-gray-300">Active Users</div>
            </div>
            <div className="text-center">
              <div className="text-3xl font-bold text-[#666eff]">4.9★</div>
              <div className="text-gray-300">App Rating</div>
            </div>
            <div className="text-center">
              <div className="text-3xl font-bold text-[#666eff]">1M+</div>
              <div className="text-gray-300">Workouts Completed</div>
            </div>
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="py-12 px-4 sm:px-6 lg:px-8 border-t border-[#5b5b5b]/30">
        <div className="max-w-7xl mx-auto">
          <div className="flex flex-col md:flex-row justify-between items-center">
            <div className="flex items-center mb-4 md:mb-0">
              <div className="w-8 h-8 bg-[#666eff] rounded-lg flex items-center justify-center mr-3">
                <Zap className="w-5 h-5 text-white" />
              </div>
              <span className="font-['AR_One_Sans',Helvetica] text-white text-xl font-bold">
                FitCoach AI
              </span>
            </div>
            <div className="text-gray-400 text-sm">
              © 2025 FitCoach AI. All rights reserved.
            </div>
          </div>
        </div>
      </footer>
    </div>
  );
};