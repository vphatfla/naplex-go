import { Route, Routes } from 'react-router'
import { AuthProvider } from './context/AuthContext'
import ProtectedRoute from './components/ProtectedRoute'
import Landing from './pages/Landing/Landing'
import Callback from './pages/Callback/Callback'
import Home from './pages/Home/Home'
import Profile from './pages/Profile/Profile'
import Settings from './pages/Setting/Setting'
import DailyQuiz from './pages/DailyQuiz'
import Question from './pages/Question'

function App() {
  return (
    <AuthProvider>
      <Routes>
        <Route path="/" element={<Landing />} />
        <Route path="/callback" element={<Callback />} />
        <Route
          path="/home"
          element={
            <ProtectedRoute>
              <Home />
            </ProtectedRoute>
          }
        />
        <Route
          path="/profile"
          element={
            <ProtectedRoute>
              <Profile />
            </ProtectedRoute>
          }
        />
        <Route
          path="/settings"
          element={
            <ProtectedRoute>
              <Settings />
            </ProtectedRoute>
          }
        />
        <Route
          path="/daily-quiz"
          element={
            <ProtectedRoute>
              <DailyQuiz />
            </ProtectedRoute>
          }
        />
        <Route
          path="/question"
          element={
            <ProtectedRoute>
              <Question />
            </ProtectedRoute>
          }
        />
      </Routes>
    </AuthProvider>
  );
}

export default App