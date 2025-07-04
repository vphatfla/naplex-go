import type { Question, QuestionUpdate } from "../types";

const API_BASE_URL = 'http://localhost:8080';

const questionService = {
    async getDailyQuestions(numQuestions: number): Promise<Question[]> {
        const response = await fetch(
            `${API_BASE_URL}/question/daily?num_question=${numQuestions}`,
            {
                credentials: 'include',
            });

        if (!response.ok) {
            throw new Error('Failed to fetch daily questions');
        }
        
        return response.json();
    },

    async getQuestion(questionId: number): Promise<Question> {
        const response = await fetch(
            `${API_BASE_URL}/question/?question_id=${questionId}`,
            {
                credentials: 'include',
            });

        if (!response.ok) {
        throw new Error('Failed to fetch question');
        }

        return response.json();
    },

    async updateQuestion(update: QuestionUpdate): Promise<Question> {
        const response = await fetch(`${API_BASE_URL}/question/`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: 'include',
            body: JSON.stringify(update),
        });

        if (!response.ok) {
            throw new Error('Failed to update question');
        }

        return response.json();
    },
}

export {questionService};