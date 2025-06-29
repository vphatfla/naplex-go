import { type ReactNode } from 'react';
import Header from '../Header/Header';

interface LayoutProps {
  children: ReactNode;
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <div className="min-h-screen bg-gray-200 dark:bg-black">
      <Header />
      <main className="pt-[52px]">
        {children}
      </main>
    </div>
  );
};

export default Layout;