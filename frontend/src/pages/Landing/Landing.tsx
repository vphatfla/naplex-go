import { useEffect } from 'react';
import { useNavigate, useLocation } from 'react-router';
import GoogleSignInButton from "../../components/auth/GoogleSignInButton";
import { useAuth } from '../../context/AuthContext';
import DarkModeToggle from '../../components/common/Button';

const Landing = () => {
    const { user, login } = useAuth();
    const navigate = useNavigate();
    const location = useLocation();
    const error = location.state?.error;

    useEffect(() => {
        // If user is already authenticated, redirect to home
        if (user) {
            navigate('/home');
        }
    }, [user, navigate]);

    const handleGoogleSignIn = () => {
        login();
    }

    return (
        <div className="min-h-screen bg-white dark:bg-black flex">
            {/* Header */}
            <header className="fixed top-0 left-0 right-0 h-[88px] bg-white/95 text-black dark:bg-black dark:text-white backdrop-blur-[20px] border-b border-[#E5E5E5] z-50">
                <div className="mx-auto px-5 md:px-[80px] h-full flex justify-between items-center">
                    <div className="text-xl font-semibold tracking-[-0.4px">
                        NAPLEX Go
                    </div>
                    <nav className="flex items-center gap-5 md:gap-8">
                        <a href="#about" className="text-base text-[#555555] tracking-[-0.24px] hover:text-black dark:text-white dark:hover:text-white -colors">
                            About
                        </a>
                        <span className="text-[#CCCCCC]">|</span>
                        <a href="#contact" className="text-base text-[#555555] tracking-[-0.24px] hover:text-black dark:text-white dark:hover:text-white transition-colors">
                            Contact Us
                        </a>
                        <DarkModeToggle></DarkModeToggle>
                    </nav>
                </div>
            </header>

            {/* Hero Section */}
            <main className="flex-1 flex items-center justify-center pt-[200px] pb-[200px] px-5">
                <div className="text-center max-w-[600px]">
                    <h1 className="text-4xl md:text-5xl font-bold tracking-[-1px] mb-4 leading-[1.2] text-black dark:text-white">
                        Naplex Go
                    </h1>
                    <p className="text-lg md:text-xl text-[#666666] tracking-[-0.4px] mb-10! leading-[1.4]">
                        Your journey to NAPLEX success starts here
                    </p>
                    
                    {error && (
                        <div className="mb-6 p-3 bg-red-50 border border-red-200 rounded-xl">
                            <p className="text-sm text-red-600 text-center">{error}</p>
                        </div>
                    )}
                    
                    <GoogleSignInButton onClick={handleGoogleSignIn} />
                </div>
            </main>

            {/* Footer */}
            <footer className="fixed bottom-10 left-0 right-0 justify-center items-center right-0bg-[#FAFAFA] py-10 text-center">
                <div className="text-base text-[#888888] font-medium tracking-[-0.24px] mb-2">
                    NAPLEX Go
                </div>
                <div className="text-sm text-[#888888] tracking-[-0.08px]">
                    Copyright © 2025 NAPLEX Go
                </div>
            </footer>
        </div>
    )
}

export default Landing;