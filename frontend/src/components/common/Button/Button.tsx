import type { ButtonHTMLAttributes, ReactNode } from "react";

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
    variant?: 'primary' | 'secondary' | 'ghost';
    size?: 'sm' | 'md' | 'lg';
    children: ReactNode;
    className?: string;
}

const Button: React.FC<ButtonProps> = ({
    variant = 'secondary',
    size = 'md',
    children,
    className = '',
    ...props
}) => {
    const baseStyles = 'inline-flex items-center justify-center font-normal cursor-pointer transition-all duration-200 hover:-translate-y-[1px] hover:shadow-lg active:translate-y-0 active:shadow-md disabled:opacity-50 disabled:cursor-not-allowed';

    const variants = {
        primary: 'bg-apple-blue hover:bg-apple-blue-dark text-white',
        secondary: 'bg-white dark:bg-black border border-apple-gray-200 dark:border-apple-gray-500 text-apple-gray-600 dark:text-apple-gray-50 hover:bg-apple-gray-50 dark:hover:bg-apple-gray-600',
        ghost: 'bg-transparent hover:bg-apple-gray-50 dark:hover:bg-apple-gray-600 text-apple-gray-600 dark:text-apple-gray-50'
      };
      
      const sizes = {
        sm: 'h-10 px-4 text-sm rounded-apple-sm gap-2',
        md: 'h-14 px-6 text-[17px] rounded-apple-sm gap-3',
        lg: 'h-16 px-8 text-lg rounded-apple gap-4'
      };
    
    return (
        <button
          className={`${baseStyles} ${variants[variant]} ${sizes[size]} ${className}`}
          {...props}
        >
          {children}
        </button>
      );
}

export default Button;