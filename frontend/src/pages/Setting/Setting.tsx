import Layout from '../../components/Layout/Layout';

const Settings: React.FC = () => {
  return (
    <Layout>
      <div className="min-h-[calc(100vh-52px)] px-6 py-12">
        <div className="max-w-[800px] mx-auto">
          <h1 className="text-3xl font-semibold text-apple-gray-600 dark:text-apple-gray-50 mb-8 animate-fade-in-up">
            Settings
          </h1>
          
          <div className="bg-apple-gray-50 dark:bg-apple-gray-100 rounded-apple-lg p-8 animate-fade-in-up" style={{ animationDelay: '0.1s' }}>
            <div className="text-center text-apple-gray-400 dark:text-apple-gray-300">
              <svg className="w-16 h-16 mx-auto mb-4 text-apple-gray-300 dark:text-apple-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
              <p className="text-sm">Settings options coming soon...</p>
            </div>
          </div>
        </div>
      </div>
    </Layout>
  );
};

export default Settings;