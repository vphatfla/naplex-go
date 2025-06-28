import { useState, useRef, useEffect } from 'react';
import { Link } from 'react-router';
import { useAuth } from '../../context/AuthContext';

const Header: React.FC = () => {
  const { user, logout } = useAuth();
  const [isDropdownOpen, setIsDropdownOpen] = useState(false);
  const dropdownRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (dropdownRef.current && !dropdownRef.current.contains(event.target as Node)) {
        setIsDropdownOpen(false);
      }
    };

    document.addEventListener('mousedown', handleClickOutside);
    return () => document.removeEventListener('mousedown', handleClickOutside);
  }, []);

  const handleLogout = async () => {
    setIsDropdownOpen(false);
    await logout();
  };

  return (
    <header className="fixed top-0 left-0 right-0 z-50 bg-white/80 dark:bg-black/80 backdrop-blur-apple border-b border-apple-gray-200 dark:border-apple-gray-500">
      <div className="max-w-[1200px] mx-auto px-6 h-[52px] flex items-center justify-between">
        {/* Logo */}
        <Link to="/home" className="flex items-center">
          <div className="w-8 h-8 bg-apple-blue rounded-apple-sm flex items-center justify-center text-white font-semibold text-sm hover:scale-105 transition-transform duration-200">
            M
          </div>
        </Link>

        {/* User Menu */}
        <div className="relative" ref={dropdownRef}>
          <button
            onClick={() => setIsDropdownOpen(!isDropdownOpen)}
            className="w-10 h-10 rounded-full overflow-hidden hover:ring-2 hover:ring-apple-gray-300 dark:hover:ring-apple-gray-400 transition-all duration-200 focus-visible:ring-2 focus-visible:ring-apple-blue"
            aria-label="User menu"
          >
            {user?.picture ? (
              <img
                src={user.picture}
                alt={user.name}
                className="w-full h-full object-cover"
              />
            ) : (
              <div className="w-full h-full bg-apple-gray-200 dark:bg-apple-gray-500 flex items-center justify-center">
                <svg className="w-6 h-6 text-apple-gray-400 dark:text-apple-gray-300" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                </svg>
              </div>
            )}
          </button>

          {/* Dropdown Menu */}
          {isDropdownOpen && (
            <div className="absolute right-0 mt-2 w-[240px] bg-white dark:bg-apple-gray-100 rounded-apple-lg shadow-apple-lg border border-apple-gray-200 dark:border-apple-gray-500 overflow-hidden animate-slide-down">
              <div className="px-4 py-3 border-b border-apple-gray-200 dark:border-apple-gray-500">
                <p className="text-sm font-medium text-apple-gray-600 dark:text-apple-gray-50 truncate">
                  {user?.name}
                </p>
                <p className="text-xs text-apple-gray-400 truncate">
                  {user?.email}
                </p>
              </div>
              
              <nav className="py-2">
                <Link
                  to="/profile"
                  onClick={() => setIsDropdownOpen(false)}
                  className="block px-4 py-2 text-sm text-apple-gray-600 dark:text-apple-gray-50 hover:bg-apple-gray-100 dark:hover:bg-apple-gray-200 transition-colors"
                >
                  Profile
                </Link>
                <Link
                  to="/settings"
                  onClick={() => setIsDropdownOpen(false)}
                  className="block px-4 py-2 text-sm text-apple-gray-600 dark:text-apple-gray-50 hover:bg-apple-gray-100 dark:hover:bg-apple-gray-200 transition-colors"
                >
                  Settings
                </Link>
                <hr className="my-2 border-apple-gray-200 dark:border-apple-gray-500" />
                <button
                  onClick={handleLogout}
                  className="block w-full text-left px-4 py-2 text-sm text-red-600 hover:bg-red-50 dark:hover:bg-red-900/20 transition-colors"
                >
                  Log out
                </button>
              </nav>
            </div>
          )}
        </div>
      </div>
    </header>
  );
};

export default Header;