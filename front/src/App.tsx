import { Outlet } from 'react-router-dom';
import { Sidebar } from '@/components/layout/sidebar';
import { HeaderComponent } from './components/layout/header';

function App() {
  return (
    <div className='grid min-h-screen w-full lg:grid-cols-[280px_1fr]'>
      <Sidebar />
      <div className='flex flex-col'>
        <HeaderComponent />
        {/* Main content */}
        <div className='w-full min-h-[calc(100vh-var(--header)-40px)] bg-gray-100 dark:bg-gray-950 px-[24px] py-[20px]'>
          <Outlet />
        </div>
      </div>
    </div>
  );
}

export default App;
