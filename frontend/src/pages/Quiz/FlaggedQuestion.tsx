import Layout from '../../components/Layout/Layout';
import { useAuth } from '../../context/AuthContext';

// questions that were flagged by the user
const FlaggedQuestion = () => {
  const { user } = useAuth();

  return (
    <Layout>
      <div className="min-h-[calc(100vh-52px)] flex flex-row items-center justify-center">
        <div className="w-1/2 text-center animate-fade-in-up">
          <h1 className="text-4xl md:text-5xl font-semibold text-apple-gray-600 dark:text-apple-gray-50 mb-4">
            Flagged Questions for { user?.name}
          </h1>
          <p className="text-lg text-apple-gray-400 dark:text-apple-gray-300">
            Under Dev
          </p>
        </div>
      </div>
    </Layout>
  );
};

export {FlaggedQuestion};