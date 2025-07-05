/** @type {import('tailwindcss').Config} */
export default {
    content: [
      "./index.html",
      "./src/**/*.{js,ts,jsx,tsx}",
    ],
    darkMode: 'class',
    theme: {
      extend: {
        colors: {
          'apple-gray': {
            50: 'var(--apple-gray-50)',
            100: 'var(--apple-gray-100)',
            200: 'var(--apple-gray-200)',
            300: 'var(--apple-gray-300)',
            400: 'var(--apple-gray-400)',
            500: 'var(--apple-gray-500)',
            600: 'var(--apple-gray-600)',
          },
          'apple-blue': 'var(--apple-blue)',
          'apple-blue-dark': 'var(--apple-blue-dark)',
        },
        fontFamily: {
          sans: ['-apple-system', 'BlinkMacSystemFont', '"SF Pro Display"', '"SF Pro Text"', '"Helvetica Neue"', 'Helvetica', 'Arial', 'sans-serif'],
        },
        animation: {
          'fade-in-up': 'fade-in-up 0.6s ease-out',
          'float': 'float 20s ease-in-out infinite',
          'slide-down': 'slide-down 0.2s ease-out',
        },
        backdropBlur: {
          'apple': '20px',
        },
        borderRadius: {
          'apple-sm': '10px',
          'apple': '12px',
          'apple-lg': '16px',
        },
        boxShadow: {
          'apple': '0 4px 16px rgba(0, 0, 0, 0.08)',
          'apple-lg': '0 8px 32px rgba(0, 0, 0, 0.12)',
        },
      },
    },
    plugins: [],
  }