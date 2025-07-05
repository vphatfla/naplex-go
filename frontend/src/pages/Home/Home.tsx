import { Link } from '../../components/common/Anchor';
import Layout from '../../components/Layout/Layout';
import { useAuth } from '../../context/AuthContext';

const Home = () => {
  const { user } = useAuth();
  
  return (
    <Layout>
      <div className="min-h-[calc(100vh-52px)] flex flex-col gap-8 md:flex-row md:gap-0 items-center justify-center px-4 py-8 md:py-0">
        <div className="w-full md:w-1/2 text-center animate-fade-in-up mb-8 md:mb-0">
          <h1 className="text-3xl sm:text-4xl md:text-5xl font-semibold text-apple-gray-600 dark:text-apple-gray-50 mb-4">
            Welcome back, {user?.first_name || user?.name}
          </h1>
          <p className="text-base sm:text-lg text-apple-gray-400 dark:text-apple-gray-300">
            Your NAPLEX preparation journey starts here.
          </p>
        </div>
        <div className="w-full md:w-1/2 text-center flex flex-col gap-4 animate-fade-in-up items-center justify-center px-4 md:px-8">
          <Link href='daily-question' className='w-full max-w-sm' variant='primary' size='md'>
            Daily Question
          </Link>
          <Link href='daily-quiz' className='w-full max-w-sm' variant='primary' size='md'>
            Daily 10 Question Quiz
          </Link>
          <Link href='random-quiz' className='w-full max-w-sm' variant='primary' size='md'>
            Random Quiz
          </Link>
          <Link href='missed-question' className='w-full max-w-sm' variant='primary' size='md'>
            Review Your Missed Questions
          </Link>
          <Link href='flagged-question' className='w-full max-w-sm' variant='primary' size='md'>
            Review Your Flagged Questions
          </Link>
        </div>
      </div>
    </Layout>
  );
};

export default Home;