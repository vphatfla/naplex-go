import type { AnchorHTMLAttributes, ReactNode } from "react";

interface LinkProps extends AnchorHTMLAttributes<HTMLAnchorElement> {
    variant?: 'primary' | 'secondary' | 'ghost';
    size?: 'sm' | 'md' | 'lg';
    children: ReactNode;
    className?: string;
    href: string;
}

const Link = ({
    variant = 'secondary',
    size = 'md',
    children,
    className = '',
    href,
    ...props
}: LinkProps) => {
    const baseStyles = 'inline-flex items-center justify-center font-normal cursor-pointer transition-all duration-200 hover:-translate-y-[1px] hover:shadow-lg active:translate-y-0 active:shadow-md disabled:opacity-50 disabled:cursor-not-allowed no-underline';

    const variants = {
        primary: 'bg-blue-500 hover:bg-blue-600 text-white',
        secondary: 'bg-white dark:bg-black border border-gray-200 dark:border-gray-500 text-gray-600 dark:text-gray-50 hover:bg-gray-50 dark:hover:bg-gray-600',
        ghost: 'bg-transparent hover:bg-gray-50 dark:hover:bg-gray-600 text-gray-600 dark:text-gray-50'
    };
      
    const sizes = {
        sm: 'h-10 px-4 text-sm rounded-md gap-2',
        md: 'h-14 px-6 text-[17px] rounded-md gap-3',
        lg: 'h-16 px-8 text-lg rounded-lg gap-4'
    };
    
    return (
        <a
            href={href}
            className={`${baseStyles} ${variants[variant]} ${sizes[size]} ${className}`}
            {...props}
        >
            {children}
        </a>
    );
}

export { Link };