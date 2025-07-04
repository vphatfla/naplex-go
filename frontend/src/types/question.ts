interface Question {
    question_id: number;
    title: string;
    question: string;
    multiple_choices: string[];
    correct_answer: string;
    explanation?: string;
    keywords?: string[];
    status?: 'PASSED' | 'FAILED' | 'NA';
    attempts?: number;
    saved?: boolean;
    hidden?: boolean;
}

interface QuestionUpdate {
  question_id: number;
  status?: 'PASSED' | 'FAILED' | 'NA';
  attempts?: number;
  saved?: boolean;
  hidden?: boolean;
}

export type {Question, QuestionUpdate};