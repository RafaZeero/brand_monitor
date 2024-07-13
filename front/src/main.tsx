import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import { RouterProvider } from 'react-router-dom';
import { routes } from './routes/root';
import { ThemeProvider } from './components/theme-provider';
import { Toaster } from '@/components/ui/toaster';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <ThemeProvider defaultTheme='dark' storageKey='vite-ui-theme'>
      <RouterProvider router={routes} />
      <Toaster />
    </ThemeProvider>
  </React.StrictMode>
);
