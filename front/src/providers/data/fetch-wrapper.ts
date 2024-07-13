import { abortSignal } from '@/lib/utils';
import axios, { AxiosRequestConfig, AxiosResponse } from 'axios';

/**
 * Executes a custom fetch request with optional access token and content type headers.
 *
 * @param {string} url - The URL to fetch.
 * @param {RequestInit} options - The options for the fetch request.
 * @return {Promise<Response>} A promise that resolves to the fetch response.
 */
const customFetch = async (url: string, options: AxiosRequestConfig) => {
  // const accessToken = localStorage.getItem('token');
  const headers = options.headers as Record<string, string>;

  return await axios(url, {
    ...options,
    headers: {
      ...headers,
      // Authorization: headers?.Authorization,
      'Content-Type': 'application/json'
    },
    signal: options?.signal || abortSignal(5000)
  });
};

/**
 * A wrapper function for making HTTP requests using the fetch API.
 *
 * @param {string} url - The URL to make the request to.
 * @param {RequestInit} options - The options for the fetch request.
 * @return {Promise<T>} A promise that resolves to the response object.
 */
export const fetchWrapper = async <T = AxiosResponse>(
  url: string,
  options: AxiosRequestConfig
): Promise<T> => {
  const response = await customFetch(url, options);

  const body = await response.data;

  if (!response.data.success) {
    return Promise.reject(body);
  }

  return body as T;
};
