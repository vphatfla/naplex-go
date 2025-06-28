import Layout from '../../components/Layout/Layout';
import { useAuth } from '../../context/AuthContext';

const Profile: React.FC = () => {
  const { user } = useAuth();

  return (
    <Layout>
      <div className="min-h-[calc(100vh-52px)] px-6 py-12">
        <div className="max-w-[800px] mx-auto">
          <h1 className="text-3xl font-semibold text-apple-gray-600 dark:text-apple-gray-50 mb-8 animate-fade-in-up">
            Profile
          </h1>
          
          <div className="bg-apple-gray-50 dark:bg-apple-gray-100 rounded-apple-lg p-8 animate-fade-in-up" style={{ animationDelay: '0.1s' }}>
            <div className="flex items-center space-x-6 mb-8">
              {user?.picture ? (
                <img
                  src={user.picture}
                  alt={user.name}
                  className="w-24 h-24 rounded-full object-cover"
                />
              ) : (
                <div className="w-24 h-24 bg-apple-gray-200 dark:bg-apple-gray-500 rounded-full flex items-center justify-center">
                  <svg className="w-12 h-12 text-apple-gray-400 dark:text-apple-gray-300" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                  </svg>
                </div>
              )}
              <div>
                <h2 className="text-2xl font-medium text-apple-gray-600 dark:text-apple-gray-50">
                  {user?.name}
                </h2>
                <p className="text-apple-gray-400">
                  {user?.email}
                </p>
              </div>
            </div>
            
            <div className="text-center text-apple-gray-400 dark:text-apple-gray-300">
              <p className="text-sm">Profile customization coming soon...</p>
            </div>
          </div>
        </div>
      </div>
    </Layout>
  );
};

export default Profile;