import { useState, useEffect } from 'react';

const DarkModeToggle = () => {
    const [darkMode, setDarkMode] = useState(false);

    useEffect(() => {
        const isDarkMode = localStorage.getItem('darkMode') === 'true';
        setDarkMode(isDarkMode);
    }, []);

    useEffect(() => {
        document.documentElement.classList.toggle('dark', darkMode);
        localStorage.setItem('darkMode', String(darkMode));
    }, [darkMode]);

    const toggleDarkMode = () => {
        setDarkMode((prevMode) => !prevMode);
    };

    return (
    <button
        onClick={toggleDarkMode}
        className="px-4 py-2 rounded-md bg-gray-800 text-white dark:bg-gray-200 dark:text-gray-800"
    >
        {darkMode ? 'Light Mode' : 'Dark Mode'}
    </button>
    );
};

export default DarkModeToggle;