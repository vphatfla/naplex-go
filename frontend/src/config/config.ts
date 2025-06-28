export const config = {
    API_URL: import.meta.env.VITE_API_URL || 'http://localhost:8080',
    FRONTEND_URL: import.meta.env.VITE_FRONTEND_URL || 'http://localhost:5173'
};

export const endpoints = {
    auth: {
        googleLogin: `${config.API_URL}/auth/google/login`,
        logout: `${config.API_URL}/auth/logout`,
    },
    user: {
        profile: `${config.API_URL}/user/profile`,
    },
    question: {
        base: `${config.API_URL}/question`,
        passed: `${config.API_URL}/question/passed`,
        failed: `${config.API_URL}/question/failed`,
        daily: `${config.API_URL}/question/daily`,
    },
};