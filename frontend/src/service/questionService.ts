import type { Question, QuestionUpdate } from "../types";
import { endpoints } from "../config/config";

const questionService = {
    async getDailyQuestions(numQuestions: number): Promise<Question[]> {
        const response = await fetch(
            `${endpoints.question.daily}?num_question=${numQuestions}`,
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
            `${endpoints.question.base}?question_id=${questionId}`,
            {
                credentials: 'include',
            });

        if (!response.ok) {
        throw new Error('Failed to fetch question');
        }

        return response.json();
    },

    async updateQuestion(update: QuestionUpdate): Promise<Question> {
        const response = await fetch(`${endpoints.question.base}`, {
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