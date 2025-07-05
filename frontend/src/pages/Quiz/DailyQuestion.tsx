import { useState, useEffect } from 'react';
import Layout from '../../components/Layout/Layout';
import QuestionCard from '../../components/Quiz/QuestionCard';
import { questionService } from '../../service';
import type { Question } from '../../types';
import Button from '../../components/common/Button/Button';
import { Link } from '../../components/common/Anchor';

const DailyQuestion = () => {
  const [question, setQuestion] = useState<Question | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [completed, setCompleted] = useState(false);

  useEffect(() => {
    fetchDailyQuestion();
  }, []);

  const fetchDailyQuestion = async () => {
    try {
      setLoading(true);
      setError(null);
      const questions = await questionService.getDailyQuestions(1);
      if (questions.length > 0) {
        setQuestion(questions[0]);
      } else {
        setError('No daily question available');
      }
    } catch (err) {
      setError('Failed to load daily question. Please try again later.');
      console.error('Error fetching daily question:', err);
    } finally {
      setLoading(false);
    }
  };

  const handleComplete = () => {
    setCompleted(true);
  };

  const handleNewQuestion = () => {
    setCompleted(false);
    fetchDailyQuestion();
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
              Loading question of the day...
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
              onClick={fetchDailyQuestion}
              className="px-6 py-3 bg-blue-500 text-white rounded-xl font-medium hover:bg-blue-600 transition-colors"
            >
              Try Again
            </button>
          </div>
        </div>
      </Layout>
    );
  }

  return (
    <Layout>
      <div className="min-h-[calc(100vh-52px)] py-8 flex items-center justify-center">
        <div className="max-w-4xl mx-auto px-4 flex flex-col gap-10">
          {/* Header */}
          <div className="text-center mb-8 animate-fade-in-up">
            <h1 className="text-3xl md:text-4xl font-semibold text-apple-gray-600 dark:text-apple-gray-50 mb-2">
              Your Daily Question
            </h1>
          </div>

          {/* Question Card */}
          {question && (
            <div className="animate-fade-in-up animation-delay-200">
              <QuestionCard 
                question={question} 
                onComplete={handleComplete}
              />
            </div>
          )}

          {/* Actions after completion */}
          {completed && (
            <div className="mt-8 text-center animate-fade-in flex flex-col gap-5">
              <p className="text-apple-gray-400 dark:text-apple-gray-300 mb-6">
                Great job! You've completed today's question.
              </p>
              <div className="flex flex-col sm:flex-row gap-4 justify-center">
                <Button
                  onClick={handleNewQuestion}
                  variant="secondary"
                >
                  Try Another Question
                </Button>
                <Link
                  href='/daily-quiz'
                  variant='primary'
                >
                  Take Full Quiz (10 Questions)
                </Link>
              </div>
            </div>
          )}

          {/* Daily Streak or Stats (Optional Enhancement) */}
        </div>
      </div>
    </Layout>
  );
};

export { DailyQuestion };