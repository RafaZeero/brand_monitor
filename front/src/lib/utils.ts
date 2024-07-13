import { type ClassValue, clsx } from 'clsx';
import { twMerge } from 'tailwind-merge';

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export const abortSignal = (timeout: number) => {
  const controller = new AbortController();
  const signal = controller.signal;
  setTimeout(() => controller.abort(), timeout);
  return signal;
};
