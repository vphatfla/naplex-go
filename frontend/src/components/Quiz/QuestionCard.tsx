import React, { useState, useEffect } from 'react';
import type { Question } from '../../types';
import { questionService } from '../../service';

interface QuestionCardProps {
  question: Question;
  onComplete?: (passed: boolean) => void;
  showNavigation?: boolean;
  questionNumber?: number;
  totalQuestions?: number;
}

const QuestionCard: React.FC<QuestionCardProps> = ({
  question,
  onComplete,
  showNavigation = false,
  questionNumber,
  totalQuestions,
}) => {
  const [selectedAnswer, setSelectedAnswer] = useState<string>('');
  const [showResult, setShowResult] = useState(false);
  const [isCorrect, setIsCorrect] = useState(false);
  const [attempts, setAttempts] = useState(question.attempts || 0);
  const [isSaved, setIsSaved] = useState(question.saved || false);
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    setSelectedAnswer('');
    setShowResult(false);
    setIsCorrect(false);
    setAttempts(question.attempts || 0);
    setIsSaved(question.saved || false);
  }, [question]);

  const handleSubmit = async () => {
    if (!selectedAnswer) return;

    const correct = selectedAnswer === question.correct_answer;
    setIsCorrect(correct);
    setShowResult(true);
    setAttempts(prev => prev + 1);

    try {
      setIsLoading(true);
      await questionService.updateQuestion({
        question_id: question.question_id,
        status: correct ? 'PASSED' : 'FAILED',
        attempts: attempts + 1,
        saved: isSaved,
        hidden: isSaved, // temp set, need to implement
      });
      
      if (onComplete) {
        onComplete(correct);
      }
    } catch (error) {
      console.error('Failed to update question status:', error);
    } finally {
      setIsLoading(false);
    }
  };

  const handleSaveToggle = async () => {
    try {
      setIsLoading(true);
      const newSavedStatus = !isSaved;
      await questionService.updateQuestion({
        question_id: question.question_id,
        status: selectedAnswer === question.correct_answer ? 'PASSED' : 'FAILED', // need to implement here
        attempts: attempts + 1,
        saved: newSavedStatus,
      });
      setIsSaved(newSavedStatus);
    } catch (error) {
      console.error('Failed to update saved status:', error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="max-w-4xl mx-auto p-6 rounded-2xl shadow-sm">
      {/* Header */}
      <div className="flex justify-between items-start mb-6">
        <div className="flex-1">
          {showNavigation && questionNumber && totalQuestions && (
            <div className="text-sm text-apple-gray-400 dark:text-apple-gray-300 mb-2">
              Question {questionNumber} of {totalQuestions}
            </div>
          )}
          <h2 className="text-xl font-semibold text-apple-gray-600 dark:text-apple-gray-50">
            {question.title}
          </h2>
        </div>
        <button
          onClick={handleSaveToggle}
          disabled={isLoading}
          className={`ml-4 p-2 rounded-lg transition-colors ${
            isSaved
              ? 'text-yellow-500 hover:text-yellow-600'
              : 'text-apple-gray-400 hover:text-apple-gray-600 dark:hover:text-apple-gray-200'
          }`}
          aria-label={isSaved ? 'Unsave question' : 'Save question'}
        >
          <svg
            className="w-6 h-6"
            fill={isSaved ? 'currentColor' : 'none'}
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={2}
              d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z"
            />
          </svg>
        </button>
      </div>

      {/* Question */}
      <div className="mb-6">
        <p className="text-lg text-apple-gray-600 dark:text-apple-gray-200 leading-relaxed">
          {question.question}
        </p>
      </div>

      {/* Multiple Choice Options */}
      <div className="space-y-3 mb-6">
        {question.multiple_choices.map((choice, index) => (
          <label
            key={index}
            className={`flex items-center p-4 rounded-xl cursor-pointer transition-all ${
              showResult
                ? choice === question.correct_answer
                  ? 'bg-green-50 dark:bg-green-900/20 border-2 border-green-500'
                  : choice === selectedAnswer && !isCorrect
                  ? 'bg-red-50 dark:bg-red-900/20 border-2 border-red-500'
                  : 'bg-apple-gray-50 dark:bg-apple-gray-700 border-2 border-transparent'
                : selectedAnswer === choice
                ? 'bg-blue-50 dark:bg-blue-900/20 border-2 border-blue-500'
                : 'bg-apple-gray-50 dark:bg-apple-gray-700 border-2 border-transparent hover:bg-apple-gray-100 dark:hover:bg-apple-gray-600'
            }`}
          >
            <input
              type="radio"
              name="answer"
              value={choice}
              checked={selectedAnswer === choice}
              onChange={(e) => setSelectedAnswer(e.target.value)}
              disabled={showResult}
              className="sr-only"
            />
            <div className="flex items-center flex-1">
              <div
                className={`w-5 h-5 rounded-full border-2 mr-3 flex items-center justify-center ${
                  selectedAnswer === choice
                    ? 'border-blue-500 bg-blue-500'
                    : 'border-apple-gray-300 dark:border-apple-gray-500'
                }`}
              >
                {selectedAnswer === choice && (
                  <div className="w-2 h-2 rounded-full bg-white" />
                )}
              </div>
              <span className="text-apple-gray-600 dark:text-apple-gray-100">
                {choice}
              </span>
            </div>
            {showResult && choice === question.correct_answer && (
              <svg
                className="w-5 h-5 text-green-500"
                fill="currentColor"
                viewBox="0 0 20 20"
              >
                <path
                  fillRule="evenodd"
                  d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                  clipRule="evenodd"
                />
              </svg>
            )}
            {showResult && choice === selectedAnswer && !isCorrect && (
              <svg
                className="w-5 h-5 text-red-500"
                fill="currentColor"
                viewBox="0 0 20 20"
              >
                <path
                  fillRule="evenodd"
                  d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
                  clipRule="evenodd"
                />
              </svg>
            )}
          </label>
        ))}
      </div>

      {/* Submit Button */}
      {!showResult && (
        <button
          onClick={handleSubmit}
          disabled={!selectedAnswer || isLoading}
          className={`w-full py-3 px-6 rounded-xl font-medium transition-all ${
            selectedAnswer && !isLoading
              ? 'bg-blue-500 text-white hover:bg-blue-600'
              : 'bg-apple-gray-200 dark:bg-apple-gray-600 text-apple-gray-400 dark:text-apple-gray-400 cursor-not-allowed'
          }`}
        >
          {isLoading ? 'Submitting...' : 'Submit Answer'}
        </button>
      )}

      {/* Result and Explanation */}
      {showResult && (
        <div className="space-y-4">
          <div
            className={`p-4 rounded-xl ${
              isCorrect
                ? 'bg-green-50 dark:bg-green-900/20 text-green-800 dark:text-green-200'
                : 'bg-red-50 dark:bg-red-900/20 text-red-800 dark:text-red-200'
            }`}
          >
            <p className="font-medium">
              {isCorrect ? '✅ Correct!' : '❌ Incorrect'}
            </p>
            {!isCorrect && (
              <p className="mt-1">
                The correct answer is: <strong>{question.correct_answer}</strong>
              </p>
            )}
          </div>

          {question.explanation && (
            <div className="p-4 bg-blue-50 dark:bg-blue-900/20 rounded-xl">
              <h3 className="font-medium text-blue-800 dark:text-blue-200 mb-2">
                Explanation
              </h3>
              <p className="text-blue-700 dark:text-blue-300">
                {question.explanation}
              </p>
            </div>
          )} 
        </div>
      )}

      {/* Stats */}
      <div className="mt-6 pt-6 border-t border-apple-gray-200 dark:border-apple-gray-600 flex items-center justify-between text-sm text-apple-gray-400 dark:text-apple-gray-300">
        <span>Attempts: {attempts}</span>
        {question.keywords && question.keywords.length > 0 && (
          <div className="flex gap-2">
            {question.keywords.map((keyword, index) => (
              <span
                key={index}
                className="px-2 py-1 bg-apple-gray-100 dark:bg-apple-gray-700 rounded-md"
              >
                {keyword}
              </span>
            ))}
          </div>
        )}
      </div>
    </div>
  );
};

export default QuestionCard;