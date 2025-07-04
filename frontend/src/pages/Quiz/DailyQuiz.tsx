import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router';
import Layout from '../../components/Layout/Layout';
import QuestionCard from '../../components/Quiz/QuestionCard';
import { useAuth } from '../../context/AuthContext';
import { questionService } from '../../service';
import type { Question } from '../../types/question';

interface QuizResult {
  questionId: number;
  passed: boolean;
  attempts: number;
}

const DailyQuiz = () => {
  const { user } = useAuth();
  const navigate = useNavigate();
  const [questions, setQuestions] = useState<Question[]>([]);
  const [currentQuestionIndex, setCurrentQuestionIndex] = useState(0);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [results, setResults] = useState<QuizResult[]>([]);
  const [quizCompleted, setQuizCompleted] = useState(false);

  useEffect(() => {
    fetchDailyQuiz();
  }, []);

  const fetchDailyQuiz = async () => {
    try {
      setLoading(true);
      setError(null);
      setResults([]);
      setCurrentQuestionIndex(0);
      setQuizCompleted(false);
      
      const fetchedQuestions = await questionService.getDailyQuestions(10);
      if (fetchedQuestions.length > 0) {
        setQuestions(fetchedQuestions);
      } else {
        setError('No questions available for the daily quiz');
      }
    } catch (err) {
      setError('Failed to load quiz questions. Please try again later.');
      console.error('Error fetching quiz questions:', err);
    } finally {
      setLoading(false);
    }
  };

  const handleQuestionComplete = (passed: boolean) => {
    const currentQuestion = questions[currentQuestionIndex];
    const newResult: QuizResult = {
      questionId: currentQuestion.question_id,
      passed,
      attempts: (currentQuestion.attempts || 0) + 1,
    };
    
    setResults([...results, newResult]);

    // Auto-advance after a short delay
    setTimeout(() => {
      if (currentQuestionIndex < questions.length - 1) {
        setCurrentQuestionIndex(currentQuestionIndex + 1);
      } else {
        setQuizCompleted(true);
      }
    }, 2000);
  };

  const handleNavigateToQuestion = (index: number) => {
    setCurrentQuestionIndex(index);
  };

  const calculateScore = () => {
    const correctAnswers = results.filter(r => r.passed).length;
    const percentage = Math.round((correctAnswers / questions.length) * 100);
    return { correctAnswers, percentage };
  };

  if (loading) {
    return (
      <Layout>
        <div className="min-h-[calc(100vh-52px)] flex items-center justify-center">
          <div className="text-center">
            <div className="inline-flex items-center justify-center w-16 h-16 mb-4">
              <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500"></div>
            </div>
            <p className="text-apple-gray-400 dark:text-apple-gray-300">
              Loading your daily quiz...
            </p>
          </div>
        </div>
      </Layout>
    );
  }

  if (error) {
    return (
      <Layout>
        <div className="min-h-[calc(100vh-52px)] flex items-center justify-center">
          <div className="text-center max-w-md">
            <svg
              className="w-16 h-16 mx-auto mb-4 text-red-500"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
              />
            </svg>
            <h2 className="text-xl font-semibold text-apple-gray-600 dark:text-apple-gray-50 mb-2">
              Oops! Something went wrong
            </h2>
            <p className="text-apple-gray-400 dark:text-apple-gray-300 mb-6">
              {error}
            </p>
            <button
              onClick={fetchDailyQuiz}
              className="px-6 py-3 bg-blue-500 text-white rounded-xl font-medium hover:bg-blue-600 transition-colors"
            >
              Try Again
            </button>
          </div>
        </div>
      </Layout>
    );
  }

  if (quizCompleted) {
    const { correctAnswers, percentage } = calculateScore();
    
    return (
      <Layout>
        <div className="min-h-[calc(100vh-52px)] flex items-center justify-center py-8">
          <div className="max-w-2xl w-full mx-auto px-4 text-center animate-fade-in-up">
            <h1 className="text-3xl md:text-4xl font-semibold text-apple-gray-600 dark:text-apple-gray-50 mb-8">
              Quiz Complete!
            </h1>
            
            {/* Score Display */}
            <div className="bg-white dark:bg-apple-gray-800 rounded-2xl p-8 shadow-sm mb-8">
              <div className="text-6xl font-bold mb-4">
                <span className={percentage >= 70 ? 'text-green-500' : 'text-orange-500'}>
                  {percentage}%
                </span>
              </div>
              <p className="text-xl text-apple-gray-600 dark:text-apple-gray-200 mb-2">
                You got {correctAnswers} out of {questions.length} questions correct
              </p>
              <p className="text-apple-gray-400 dark:text-apple-gray-300">
                {percentage >= 70 
                  ? "Great job! Keep up the excellent work!" 
                  : "Keep practicing, you're improving!"}
              </p>
            </div>

            {/* Question Summary */}
            <div className="bg-apple-gray-50 dark:bg-apple-gray-800 rounded-2xl p-6 mb-8">
              <h3 className="text-lg font-medium text-apple-gray-600 dark:text-apple-gray-50 mb-4">
                Question Summary
              </h3>
              <div className="grid grid-cols-5 gap-2">
                {results.map((result, index) => (
                  <div
                    key={result.questionId}
                    className={`p-3 rounded-lg font-medium ${
                      result.passed 
                        ? 'bg-green-100 dark:bg-green-900/30 text-green-700 dark:text-green-300' 
                        : 'bg-red-100 dark:bg-red-900/30 text-red-700 dark:text-red-300'
                    }`}
                  >
                    Q{index + 1}
                  </div>
                ))}
              </div>
            </div>

            {/* Actions */}
            <div className="flex flex-col sm:flex-row gap-4 justify-center">
              <button
                onClick={fetchDailyQuiz}
                className="px-6 py-3 bg-blue-500 text-white rounded-xl font-medium hover:bg-blue-600 transition-colors"
              >
                Take Another Quiz
              </button>
              <button
                onClick={() => navigate('/home')}
                className="px-6 py-3 bg-apple-gray-200 dark:bg-apple-gray-700 text-apple-gray-600 dark:text-apple-gray-100 rounded-xl font-medium hover:bg-apple-gray-300 dark:hover:bg-apple-gray-600 transition-colors"
              >
                Back to Home
              </button>
            </div>
          </div>
        </div>
      </Layout>
    );
  }

  return (
    <Layout>
      <div className="min-h-[calc(100vh-52px)] py-8 flex items-center justify-center">
        <div className="max-w-4xl mx-auto px-4">
          {/* Header */}
          <div className="text-center mb-8 animate-fade-in-up">
            <h1 className="text-3xl md:text-4xl font-semibold text-apple-gray-600 dark:text-apple-gray-50 mb-2">
              Daily Quiz
            </h1>
            <p className="text-lg text-apple-gray-400 dark:text-apple-gray-300">
              {user?.first_name ? `Let's go, ${user.first_name}!` : "Let's go!"} 
              Complete all 10 questions to test your knowledge.
            </p>
          </div>

          {/* Progress Bar */}
          <div className="mb-8 animate-fade-in-up animation-delay-100">
            <div className="flex justify-between items-center mb-2">
              <span className="text-sm text-apple-gray-400 dark:text-apple-gray-300">
                Progress
              </span>
              <span className="text-sm text-apple-gray-400 dark:text-apple-gray-300">
                {results.length} of {questions.length} completed
              </span>
            </div>
            <div className="w-full bg-apple-gray-200 dark:bg-apple-gray-700 rounded-full h-2">
              <div
                className="bg-blue-500 h-2 rounded-full transition-all duration-300"
                style={{ width: `${(results.length / questions.length) * 100}%` }}
              />
            </div>
          </div>

          {/* Question Navigation Pills */}
          <div className="mb-8 animate-fade-in-up animation-delay-200">
            <div className="flex flex-wrap gap-2 justify-center">
              {questions.map((_, index) => {
                const isCompleted = index < results.length;
                const isCurrent = index === currentQuestionIndex;
                const result = results.find((r, i) => i === index);
                
                return (
                  <button
                    key={index}
                    onClick={() => isCompleted && handleNavigateToQuestion(index)}
                    disabled={!isCompleted || isCurrent}
                    className={`w-10 h-10 rounded-lg font-medium transition-all ${
                      isCurrent
                        ? 'bg-blue-500 text-white scale-110'
                        : isCompleted
                        ? result?.passed
                          ? 'bg-green-100 dark:bg-green-900/30 text-green-700 dark:text-green-300 hover:scale-105 cursor-pointer'
                          : 'bg-red-100 dark:bg-red-900/30 text-red-700 dark:text-red-300 hover:scale-105 cursor-pointer'
                        : 'bg-apple-gray-100 dark:bg-apple-gray-700 text-apple-gray-400 dark:text-apple-gray-400'
                    }`}
                  >
                    {index + 1}
                  </button>
                );
              })}
            </div>
          </div>

          {/* Current Question */}
          {questions[currentQuestionIndex] && (
            <div className="animate-fade-in-up animation-delay-300">
              <QuestionCard
                question={questions[currentQuestionIndex]}
                onComplete={handleQuestionComplete}
                showNavigation={true}
                questionNumber={currentQuestionIndex + 1}
                totalQuestions={questions.length}
              />
            </div>
          )}

          {/* Manual Navigation (for answered questions) */}
          {results.length > 0 && !quizCompleted && (
            <div className="mt-8 flex justify-between items-center animate-fade-in">
              <button
                onClick={() => setCurrentQuestionIndex(Math.max(0, currentQuestionIndex - 1))}
                disabled={currentQuestionIndex === 0}
                className={`flex items-center gap-2 px-4 py-2 rounded-lg font-medium transition-colors ${
                  currentQuestionIndex === 0
                    ? 'bg-apple-gray-100 dark:bg-apple-gray-700 text-apple-gray-300 dark:text-apple-gray-500 cursor-not-allowed'
                    : 'bg-apple-gray-200 dark:bg-apple-gray-700 text-apple-gray-600 dark:text-apple-gray-100 hover:bg-apple-gray-300 dark:hover:bg-apple-gray-600'
                }`}
              >
                <svg
                  className="w-5 h-5"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M15 19l-7-7 7-7"
                  />
                </svg>
                Previous
              </button>
              
              {currentQuestionIndex < results.length - 1 && (
                <button
                  onClick={() => setCurrentQuestionIndex(Math.min(questions.length - 1, currentQuestionIndex + 1))}
                  className="flex items-center gap-2 px-4 py-2 bg-apple-gray-200 dark:bg-apple-gray-700 text-apple-gray-600 dark:text-apple-gray-100 rounded-lg font-medium hover:bg-apple-gray-300 dark:hover:bg-apple-gray-600 transition-colors"
                >
                  Next
                  <svg
                    className="w-5 h-5"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M9 5l7 7-7 7"
                    />
                  </svg>
                </button>
              )}
            </div>
          )}
        </div>
      </div>
    </Layout>
  );
};

export { DailyQuiz };