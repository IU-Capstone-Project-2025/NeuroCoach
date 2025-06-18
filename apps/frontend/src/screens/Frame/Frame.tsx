import { HomeIcon, SendIcon, MenuIcon } from "lucide-react";
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

export const Frame = (): JSX.Element => {
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

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
                  className="relative bg-[#1f1f1fc9] rounded-[15px] shadow-[inset_-2px_2px_7px_1px_#00000054] backdrop-blur-[17.95px] backdrop-brightness-[100%] p-4 md:p-6 aspect-square flex flex-col items-center justify-center"
                >
                  <div className="absolute inset-4 bg-neutral-900 rounded-[15px] shadow-[inset_-2px_2px_7px_1px_#00000054] blur-[9.55px]" />

                  <Card
                    className="relative w-16 h-16 md:w-20 md:h-20 lg:w-24 lg:h-24 rounded-[15px] border-[0.5px] border-solid border-[#f5f5f563] shadow-[inset_-2px_2px_7px_1px_#00000054]"
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