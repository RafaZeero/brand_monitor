import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';

import { Button } from '@/components/ui/button';
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage
} from '@/components/ui/form';
import axios from 'axios';
import { useState } from 'react';
import { Input } from '@/components/ui/input';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';

const formSchema = z.object({
  email: z.string().min(1, { message: 'Email é obrigatório' }).email()
});

export const FetchTermsPage = () => {
  const [terms, setTerms] = useState<
    {
      id: string;
      user_email: string;
      term: string;
      competitors: string[];
      created_at: Date;
    }[]
  >([]);
  const [isLoading, setIsLoading] = useState(false);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: ''
    }
  });

  const onSubmit = async ({ email }: z.infer<typeof formSchema>) => {
    console.log({ email });
    setIsLoading(true);
    try {
      const url = new URL('http://localhost:3333/v1/search');
      url.searchParams.append('email', email);
      const res = await axios.get(url.toString());
      console.log({ res: res.data });
      setTerms(res.data.data);
    } catch (error) {
      console.log(error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div>
      <h1 className='text-3xl mb-4 underline underline-offset-8'>Buscar termos</h1>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className='space-y-8 w-[500px]'>
          <FormField
            control={form.control}
            name='email'
            render={({ field }) => (
              <FormItem>
                <FormLabel>Email</FormLabel>
                <FormControl>
                  <Input
                    placeholder='Digite seu melhor email'
                    {...field}
                    className='p-[16px] h-[54px]'
                  />
                </FormControl>
                <FormDescription>
                  Email utilizado para enviar as informações da pesquisa.
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
          <div className='flex flex-col gap-8 w-fit'>
            <Button type='submit' disabled={isLoading}>
              Buscar
            </Button>
          </div>
        </form>
      </Form>
      <div className='flex flex-wrap gap-8 mt-4'>
        {terms.length > 0 &&
          terms.map(term => (
            <Card key={term.id} className='w-full sm:w-[350px]'>
              <CardHeader>
                <CardTitle>{term.term}</CardTitle>
              </CardHeader>
              <CardContent>
                {term.competitors.map((competitor, idx) => (
                  <p key={`${term.id}-${idx}`}>{competitor}</p>
                ))}
              </CardContent>
            </Card>
          ))}
      </div>
    </div>
  );
};
