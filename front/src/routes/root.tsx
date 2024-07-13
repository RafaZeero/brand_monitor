import App from '@/App';
import { HomePage } from '@/pages/home';
import { NotFound } from '@/pages/not_found';
import { RegisterTermPage } from '@/pages/register_term';
import { createBrowserRouter } from 'react-router-dom';

export const routes = createBrowserRouter([
  {
    path: '/',
    element: <App />,
    children: [
      {
        path: '',
        element: <HomePage />
      },
      {
        path: 'register-term',
        element: <RegisterTermPage />
      }
    ]
  },
  // Fallback to home if no route is matched
  {
    path: '*',
    element: <NotFound />
  }
]);
