import { CircleCheck, CircleOff } from 'lucide-react';
import { useEffect, useState } from 'react';
import { fetchWrapper } from '@/providers/data/fetch-wrapper';

export const HomePage = () => {
  const [status, setStatus] = useState<boolean>(false);
  const [isLoading, setIsLoading] = useState<boolean>(true);

  const showStatus = () => {
    return status ? (
      <CircleCheck className='h-4 w-4' color='green' />
    ) : (
      <CircleOff className='h-4 w-4' color='red' />
    );
  };

  const handleStatus = (newStatus: boolean) => {
    setStatus(newStatus);
  };

  useEffect(() => {
    const getStatusApi = async () => {
      setIsLoading(true);
      try {
        const res = await fetchWrapper<{ success: boolean; status: string }>(
          'http://localhost:3333/v1/healthz',
          {
            method: 'GET'
          }
        );

        if (res.success) {
          handleStatus(true);
        }
      } catch (error) {
        console.log(error);
        handleStatus(false);
      } finally {
        setIsLoading(false);
      }
    };

    getStatusApi();
  }, []);

  return (
    <div>
      <p className='inline-flex text-center align-middle items-center gap-[8px]'>
        Api status: {isLoading ? '' : showStatus()}
      </p>
    </div>
  );
};
