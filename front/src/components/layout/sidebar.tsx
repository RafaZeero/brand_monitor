import { Link, useLocation } from 'react-router-dom';
import { cn } from '@/lib/utils';
import { ActivityIcon, HomeIcon, WholeWord } from 'lucide-react';
import { ReactNode } from 'react';
import { PROJECT_FULL_NAME } from '@/config';

export const Sidebar = () => {
  return (
    <div className='hidden border-r bg-gray-100/40 lg:block dark:bg-gray-800/40'>
      <div className='flex h-(var(--header)) max-h-screen flex-col gap-2'>
        <div className='flex h-[60px] items-center border-b px-6'>
          <Link className='flex items-center gap-2 font-semibold' to='/'>
            <ActivityIcon className='h-6 w-6' />
            <span>{PROJECT_FULL_NAME}</span>
          </Link>
        </div>
        <div className='flex-1 overflow-auto py-2'>
          <nav className='grid items-start px-4 text-sm font-medium gap-[8px]'>
            <NavigationLink to='/'>
              <HomeIcon className='h-4 w-4' />
              Página principal
            </NavigationLink>
            <NavigationLink to='/register-term'>
              <WholeWord className='h-4 w-4' />
              Cadastrar Termo
            </NavigationLink>
          </nav>
        </div>
      </div>
    </div>
  );
};

const NavigationLink = ({ to, children }: { to: string; children: ReactNode }) => {
  const location = useLocation();

  const activeRoute = (routeName: string) => {
    return location.pathname === routeName
      ? // ? 'dark:border-l-[3px] dark:border-l-white bg-black dark:bg-transparent dark:text-gray-400 text-white hover:text-white'
        `dark:border-l-[5px] dark:border-l-white dark:bg-transparent dark:text-white 
        border-l-[5px] border-l-black text-black`
      : '';
  };

  return (
    <Link
      className={cn(
        'flex items-center gap-3 rounded-none px-3 py-2 text-gray-500 transition-all hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-50',
        activeRoute(to)
      )}
      to={to}
    >
      {children}
    </Link>
  );
};
