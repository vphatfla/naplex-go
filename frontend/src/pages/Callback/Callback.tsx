import { useEffect } from 'react';
import { useNavigate } from 'react-router';
import { useAuth } from '../../context/AuthContext';

const Callback: React.FC = () => {
  const navigate = useNavigate();
  const { checkAuth } = useAuth();

  useEffect(() => {
    const handleCallback = async () => {
      // The backend handles the OAuth callback and sets the session cookie
      // We just need to check if the user is authenticated
      await checkAuth();
      
      // Check if there's an error in the URL (OAuth failure)
      const urlParams = new URLSearchParams(window.location.search);
      const error = urlParams.get('error');
      
      if (error) {
        // Redirect to landing page with error
        navigate('/', { state: { error: 'Authentication failed. Please try again.' } });
      } else {
        // Successful authentication, redirect to home
        navigate('/home');
      }
    };

    handleCallback();
  }, [navigate, checkAuth]);

  return (
    <div className="min-h-screen flex items-center justify-center bg-white dark:bg-black">
      <div className="text-center">
        <div className="w-12 h-12 border-4 border-apple-gray-200 dark:border-apple-gray-500 rounded-full animate-spin border-t-apple-blue mx-auto mb-4"></div>
        <p className="text-apple-gray-400">Completing authentication...</p>
      </div>
    </div>
  );
};

export default Callback;