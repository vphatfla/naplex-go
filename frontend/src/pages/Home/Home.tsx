import Layout from '../../components/Layout/Layout';
import { useAuth } from '../../context/AuthContext';

const Home: React.FC = () => {
  const { user } = useAuth();

  return (
    <Layout>
      <div className="min-h-[calc(100vh-52px)] flex items-center justify-center px-6">
        <div className="text-center animate-fade-in-up">
          <h1 className="text-4xl md:text-5xl font-semibold text-apple-gray-600 dark:text-apple-gray-50 mb-4">
            Welcome back, {user?.first_name || user?.name}
          </h1>
          <p className="text-lg text-apple-gray-400 dark:text-apple-gray-300">
            Your NAPLEX preparation journey starts here.
          </p>
        </div>
      </div>
    </Layout>
  );
};

export default Home;