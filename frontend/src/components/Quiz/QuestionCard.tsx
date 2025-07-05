import React, { useState, useEffect } from 'react';
import type { Question } from '../../types';
import { questionService } from '../../service';
import Button from '../common/Button/Button';

interface QuestionCardProps {
  question: Question;
  onComplete?: () => void; // mark the question as done/complete/attempted, does not require the user did correctly
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
  const [isLoading, setIsLoading] = useState(false);
  const [isSaved, setIsSaved] = useState(question.saved || false);
  const [isHidden, setIsHidden] = useState(question.hidden || false);

  useEffect(() => {
    setSelectedAnswer('');
    setShowResult(false);
    setIsCorrect(false);
    setAttempts(question.attempts || 0);
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
        saved: question.saved || false,
        hidden: question.saved || false,
      });
      
      if (onComplete) {
        onComplete();
      }
    } catch (error) {
      console.error('Failed to update question status:', error);
    } finally {
      setIsLoading(false);
    }
  };

  const handleSaved = async () => {
    const newSaved = !isSaved;
    const correct = selectedAnswer === question.correct_answer;
    try {
      setIsLoading(true);
      await questionService.updateQuestion({
        question_id: question.question_id,
        status: correct ? 'PASSED' : 'FAILED',
        saved: newSaved
      });
          setIsSaved(newSaved);
    } catch (error) {
      console.error('Failed to update question status:', error);
    } finally {
      setIsLoading(false);
    }
  };

  const handleHidden = async () => {
    const newHidden = !isHidden;

    const correct = selectedAnswer === question.correct_answer;
    try {
      setIsLoading(true);
      await questionService.updateQuestion({
        question_id: question.question_id,
        status: correct ? 'PASSED' : 'FAILED',
        hidden: newHidden
      });
      setIsHidden(newHidden);
    } catch (error) {
      console.error('Failed to update question status:', error);
    } finally {
      setIsLoading(false);
    }
  };
  const getChoiceStyle = (choice: string) => {
    if (!showResult) {
      return
       'bg-blue-50 dark:bg-blue-900 border-blue-500 dark:border-blue-400'
        //: 'bg-white dark:bg-apple-gray-800 border-apple-gray-200 dark:border-apple-gray-700 hover:border-apple-gray-300 dark:hover:border-apple-gray-600';
    }

    if (choice === question.correct_answer) {
      return 'bg-green-50 dark:bg-green-900/10 border-green-500 dark:border-green-400';
    }
    
    if (choice === selectedAnswer && !isCorrect) {
      return 'bg-red-50 dark:bg-red-900/10 border-red-500 dark:border-red-400';
    }

    return 'bg-white dark:bg-green border-apple-gray-200 dark:border-apple-gray-700 dark:text-black opacity-60';
  };

  const getChoiceIcon = (choice: string) => {
    if (!showResult) return null;

    if (choice === question.correct_answer) {
      return (
        <svg className="w-5 h-5 text-green-500" fill="currentColor" viewBox="0 0 20 20">
          <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clipRule="evenodd" />
        </svg>
      );
    }

    if (choice === selectedAnswer && !isCorrect) {
      return (
        <svg className="w-5 h-5 text-red-500" fill="currentColor" viewBox="0 0 20 20">
          <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clipRule="evenodd" />
        </svg>
      );
    }

    return null;
  };

  return (
    <div className="max-w-3xl mx-auto flex flex-col gap-10">
      {/* Progress indicator - minimal and elegant */}
      {showNavigation && questionNumber && totalQuestions && (
        <div className="mb-8 text-center">
          <div className="inline-flex items-center space-x-2">
            <span className="text-sm font-medium text-apple-gray-500 dark:text-apple-gray-400">
              {questionNumber}
            </span>
            <span className="text-sm text-apple-gray-300 dark:text-apple-gray-600">/</span>
            <span className="text-sm text-apple-gray-400 dark:text-apple-gray-500">
              {totalQuestions}
            </span>
          </div>
        </div>
      )}

      {/* Question */}
      <div className="mb-10">
        <h2 className="text-2xl font-medium text-apple-gray-900 dark:text-white leading-relaxed">
          {question.question}
        </h2>
      </div>

      {/* Answer choices - rectangular cards */}
      <div className="space-y-3 mb-10 flex flex-col gap-5">
        {question.multiple_choices.map((choice, index) => (
          <label
            key={index}
            className={`
              block relative p-5 rounded-md border-2 cursor-pointer
              transition-all duration-200 transform
              ${getChoiceStyle(choice)}
              ${!showResult && selectedAnswer === choice ? 'scale-[1.02] shadow-lg' : ''}
              ${!showResult && selectedAnswer === choice ? 'bg-blue-400  dark:text-white' : ''}
              ${!showResult ? 'hover:scale-[1.01] hover:shadow-md' : ''}
            `}
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
            
            <div className="flex items-center justify-between">
              <span className={`
                text-base font-medium
                ${showResult && choice === question.correct_answer ? 'text-green-700 dark:text-green-300' : ''}
                ${showResult && choice === selectedAnswer && !isCorrect ? 'text-red-700 dark:text-red-300' : ''}
                ${!showResult || (choice !== question.correct_answer && choice !== selectedAnswer) ? 'text-apple-gray-700 dark:text-apple-gray-200' : ''}
              `}>
                {choice}
              </span>
              {getChoiceIcon(choice)}
            </div>

          </label>
        ))}
      </div>

      {/* Submit button - only show when not showing results */}
      {!showResult && (
        <Button
          onClick={handleSubmit}
          disabled={!selectedAnswer || isLoading}
          className={`
            w-full py-4 px-6 font-medium text-base
            transition-all duration-200 transform
            ${!selectedAnswer || isLoading
              ? 'bg-apple-gray-100 dark:bg-apple-gray-800 text-apple-gray-400 cursor-not-allowed'
              : ''
            }
          `}
        >
          {isLoading ? 'Checking...' : 'Submit Answer'}
        </Button>
      )}

      {/* Result feedback - simplified and elegant */}
      {showResult && (
        <div className="space-y-4 animate-fade-in flex flex-col gap-5">
          {/* Result message */}
          <div className={`
            p-6 rounded-2xl text-center
            ${isCorrect 
              ? 'bg-green-50 dark:bg-green-900/10' 
              : 'bg-red-50 dark:bg-red-900/10'
            }
          `}>
            <p className={`
              text-lg font-medium
              ${isCorrect 
                ? 'text-green-700 dark:text-green-300' 
                : 'text-red-700 dark:text-red-300'
              }
            `}>
              {isCorrect ? 'Correct! Well done.' : `Incorrect. The answer is ${question.correct_answer}.`}
            </p>
            <div className="p-6 bg-apple-gray-50 dark:bg-apple-gray-800/50 rounded-2xl">
              <p className="text-apple-gray-600 dark:text-apple-gray-300 leading-relaxed">
                {question.explanation}
              </p>
            </div>
          </div>
          <div className="flex items-center justify-center gap-3 p-2">
            {/* Save Icon */}
            <button onClick={handleSaved} 
                  className={`p-2 rounded-lg transition-colors ${
                  isSaved 
                    ? 'text-yellow-500 hover:text-yellow-600 bg-yellow-50 dark:bg-yellow-500/10' 
                    : 'text-apple-gray-400 hover:text-apple-gray-600 dark:text-apple-gray-300 dark:hover:text-apple-gray-100 hover:bg-apple-gray-100 dark:hover:bg-apple-gray-700'
                } ${isLoading ? 'opacity-50 cursor-not-allowed' : ''}`}
                aria-label={isSaved ? 'Unsave question' : 'Save question'}
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
                  d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z"
                />
              </svg>
            </button>

            {/* Hidden/Eye Icon */}
            <button onClick={handleHidden}
                className={`p-2 rounded-lg transition-colors ${
                  isHidden 
                    ? 'text-yellow-500 hover:text-yellow-600 bg-yellow-50 dark:bg-yellow-500/10' 
                    : 'text-apple-gray-400 hover:text-apple-gray-600 dark:text-apple-gray-300 dark:hover:text-apple-gray-100 hover:bg-apple-gray-100 dark:hover:bg-apple-gray-700'
                } ${isLoading ? 'opacity-50 cursor-not-allowed' : ''}`}
                aria-label={isHidden ? 'Un-hide question' : 'Hide question'}
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
                  d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21"
                />
              </svg>
            </button>
          </div>
        </div>
        
      )}
    </div>
  );
};

export default QuestionCard;