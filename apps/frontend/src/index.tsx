import { StrictMode, useState } from "react";
import { createRoot } from "react-dom/client";
import { Landing } from "./screens/Landing";
import { Frame } from "./screens/Frame";

const App = () => {
  const [currentView, setCurrentView] = useState<'landing' | 'training'>('landing');

  const handleNavigateToTraining = () => {
    setCurrentView('training');
  };

  const handleNavigateToLanding = () => {
    setCurrentView('landing');
  };

  if (currentView === 'training') {
    return <Frame onNavigateToLanding={handleNavigateToLanding} />;
  }

  return <Landing onNavigateToTraining={handleNavigateToTraining} />;
};

createRoot(document.getElementById("app") as HTMLElement).render(
  <StrictMode>
    <App />
  </StrictMode>,
);