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
import { Input } from '@/components/ui/input';
import { useState } from 'react';
import { InputTags } from '@/components/ui/input-tag';
import axios from 'axios';

const formSchema = z.object({
  email: z.string().min(1, { message: 'Email é obrigatório' }).email('Email inválido'),
  terms: z.array(z.string())
});

export const RegisterTermPage = () => {
  const [isLoading, setIsLoading] = useState(false);
  const [values, setValues] = useState<string[]>([]);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: '',
      terms: []
    }
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    setIsLoading(true);
    try {
      const res = (
        await axios.post('http://localhost:3333/v1/search', {
          email: values.email,
          items: values.terms
        })
      ).data;

      if (res.success) {
        console.log(res);
      }
    } catch (error) {
      console.log(error);
    } finally {
      setIsLoading(false);
    }
  }
  return (
    <div>
      <h1 className='text-3xl mb-4 underline underline-offset-8'>Verificar termos</h1>
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
                <FormDescription>Email que você vai receber as informações.</FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name='terms'
            render={({ field }) => (
              <FormItem>
                <FormLabel>Termos</FormLabel>
                <FormControl>
                  <InputTags
                    {...field}
                    className='dark:bg-background'
                    placeholder='Digite um termo para ser verificado'
                    value={values}
                    onChange={newTags => {
                      setValues(newTags);
                      form.setValue('terms', newTags as string[]);
                    }}
                  />
                </FormControl>
                <FormDescription>
                  Termos separados por &#39;ENTER&#39; ou &#39;vírugla&#39;.
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />

          <div className='flex flex-col gap-8 w-fit'>
            <Button type='submit' disabled={isLoading}>
              Verificar
            </Button>
          </div>
        </form>
      </Form>
    </div>
  );
};
