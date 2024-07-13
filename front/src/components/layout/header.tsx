import { Link } from 'react-router-dom';
import { WalletIcon } from '../icons';
import { ModeToggle } from './mode-toggle';
import { UserAvatar } from './avatar';

export const HeaderComponent = () => {
  return (
    <header className='flex h-14 lg:h-[60px] items-center gap-4 border-b bg-gray-100/40 px-6 dark:bg-gray-800/40'>
      <Link className='lg:hidden' to='#'>
        <WalletIcon className='h-6 w-6' />
        <span className='sr-only'>Home</span>
      </Link>
      <div className='w-fit ml-auto flex items-center gap-[16px]'>
        <ModeToggle />
        <UserAvatar />
      </div>
    </header>
  );
};
