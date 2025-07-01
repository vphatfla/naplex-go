// import { useNavigate } from 'react-router';
// import Button from '../../components/common/Button/Button';
import Layout from '../../components/Layout/Layout';
import { useAuth } from '../../context/AuthContext';

const Question = () => {
  const { user } = useAuth();
//   const navigate = useNavigate()

//   const handleDailyPopQuiz = () => {
//     navigate("/daily-quiz");
//   }

//   const handleCompQuestion = () => {
//     navigate("/question");
//   }
  
  return (
    <Layout>
      <div className="min-h-[calc(100vh-52px)] flex flex-row items-center justify-center">
        <div className="w-1/2 text-center animate-fade-in-up">
          <h1 className="text-4xl md:text-5xl font-semibold text-apple-gray-600 dark:text-apple-gray-50 mb-4">
            Comprehensive Question for { user?.name}
          </h1>
          <p className="text-lg text-apple-gray-400 dark:text-apple-gray-300">
            Under Dev
          </p>
        </div>
        {/* <div className="w-1/2 text-center flex flex-col gap-4 animate-fade-in-up justify-center">
          <Button className='w-100' onClick={handleDailyPopQuiz}>Daily Pop Quiz</Button>
          <Button className='w-100' onClick={handleCompQuestion}>Comprehensive Questions</Button>
        </div> */}
      </div>
    </Layout>
  );
};

export default Question;