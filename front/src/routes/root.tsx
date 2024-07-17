import App from '@/App';
import { FetchTermsPage } from '@/pages/fetch_results';
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
      },
      {
        path: 'fetch-term',
        element: <FetchTermsPage />
      }
    ]
  },
  // Fallback to home if no route is matched
  {
    path: '*',
    element: <NotFound />
  }
]);
