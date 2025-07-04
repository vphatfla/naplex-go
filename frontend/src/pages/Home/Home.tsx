import { Link } from '../../components/common/Anchor';
import Layout from '../../components/Layout/Layout';
import { useAuth } from '../../context/AuthContext';

const Home = () => {
  const { user } = useAuth();
  
  return (
    <Layout>
      <div className="min-h-[calc(100vh-52px)] flex flex-row items-center justify-center">
        <div className="w-1/2 text-center animate-fade-in-up">
          <h1 className="text-4xl md:text-5xl font-semibold text-apple-gray-600 dark:text-apple-gray-50 mb-4">
            Welcome back, {user?.first_name || user?.name}
          </h1>
          <p className="text-lg text-apple-gray-400 dark:text-apple-gray-300">
            Your NAPLEX preparation journey starts here.
          </p>
        </div>
        <div className="w-1/2 text-center flex flex-col gap-4 animate-fade-in-up items-center justify-center">
            <Link href='daily-question' className='w-100' variant='primary' size='md'>Daily Question</Link>
            <Link href='daily-quiz' className='w-100' variant='primary' size='md'>Daily 10 Question Quiz</Link>
            <Link href='random-quiz' className='w-100' variant='primary' size='md'>Random Quiz</Link>
            <Link href='missed-question' className='w-100' variant='primary' size='md'>Review Your Missed Question</Link>
            <Link href='flagged-question' className='w-100' variant='primary' size='md'>Review Your Flagged Question</Link>
        </div>
      </div>
    </Layout>
  );
};

export default Home;