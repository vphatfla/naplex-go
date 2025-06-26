import GoogleSignInButton from "../../components/auth/GoogleSignInButton";

const Landing = () => {
    const handleGoogleSignIn = () => {
        console.log("To be implemented handle google sign in")
    }

    return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-white to-apple-gray-50 dark:from-black dark:to-apple-gray-600 overflow-x-hidden relative">
        {/* Floating elements for visual interest */}
        <div className="fixed w-[300px] h-[300px] -top-[150px] -left-[150px] rounded-full bg-gradient-radial from-apple-blue/5 to-transparent animate-float pointer-events-none" />
        <div className="fixed w-[300px] h-[300px] -bottom-[150px] -right-[150px] rounded-full bg-gradient-radial from-apple-blue/5 to-transparent animate-float pointer-events-none [animation-delay:10s]" />
        
        <div className="w-full max-w-[400px] px-5 animate-fade-in-up">
          {/* Logo */}
          <div className="text-center mb-[60px]">
            <div className="inline-flex items-center justify-center w-20 h-20 bg-apple-blue rounded-apple text-white text-4xl font-semibold shadow-lg hover:scale-105 transition-transform duration-300 ease-out">
              M
            </div>
          </div>
          
          {/* Auth Card */}
          <div className="bg-white dark:bg-black rounded-apple p-12 shadow-xl backdrop-blur-apple border border-apple-gray-200 dark:border-apple-gray-500">
            <h1 className="text-[28px] font-semibold text-center mb-3 tracking-tight text-apple-gray-600 dark:text-apple-gray-50">
              Welcome!
            </h1>
            
            <GoogleSignInButton onClick={handleGoogleSignIn} />
            
            {/* Privacy Notice */}
            <div className="mt-10 pt-6 border-t border-apple-gray-200 dark:border-apple-gray-500 text-center text-xs text-apple-gray-400 leading-relaxed">
              By continuing, you agree to our{' '}
              <a href="#" className="hover:border-b hover:border-apple-gray-400 transition-all">
                Terms of Service
              </a> and{' '}
              <a href="#" className="hover:border-b hover:border-apple-gray-400 transition-all">
                Privacy Policy
              </a>
            </div>
          </div>
        </div>
    </div>
    )
}

export default Landing;