/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: ['class'],
  content: [
    './pages/**/*.{ts,tsx}',
    './components/**/*.{ts,tsx}',
    './app/**/*.{ts,tsx}',
    './src/**/*.{ts,tsx}'
  ],
  prefix: '',
  theme: {
    container: {
      center: true,
      padding: '2rem',
      screens: {
        '2xl': '1400px'
      }
    },
    extend: {
      colors: {
        transparent: 'transparent',
        current: 'currentColor',

        /* Ultramarine Blue */
        'ultramarine-blue-3355ff': 'var(--ultramarine-blue-3355ff)',
        'ultramarine-blue-4462ff': 'var(--ultramarine-blue-4462ff)',
        'ultramarine-blue-5c77ff': 'var(--ultramarine-blue-5c77ff)',
        'ultramarine-blue-8599ff': 'var(--ultramarine-blue-8599ff)',
        'ultramarine-blue-adbbff': 'var(--ultramarine-blue-adbbff)',

        /* Platinum White */
        'platinum-white-c9c9c9': 'var(--platinum-white-c9c9c9)',
        'platinum-white-e1e1e1': 'var(--platinum-white-e1e1e1)',
        'platinum-white-f0f0f0': 'var(--platinum-white-f0f0f0)',
        'platinum-white-f5f5f5': 'var(--platinum-white-f5f5f5)',
        'platinum-white-fafafa': 'var(--platinum-white-fafafa)',

        /* Eeire Black */
        'eeire-black-161517': 'var(--eeire-black-161517)',
        'eeire-black-1c1c1e': 'var(--eeire-black-1c1c1e)',
        'eeire-black-28282a': 'var(--eeire-black-28282a)',
        'eeire-black-464649': 'var(--eeire-black-464649)',
        'eeire-black-5a5a5e': 'var(--eeire-black-5a5a5e)',

        /* Black Mamba */
        'black-mamba-100f11': 'var(--black-mamba-100f11)',

        /* Palet Violet */
        'palet-violet-b899ff': 'var(--palet-violet-b899ff)',
        'palet-violet-c6adff': 'var(--palet-violet-c6adff)',
        'palet-violet-d4c2ff': 'var(--palet-violet-d4c2ff)',
        'palet-violet-e2d6ff': 'var(--palet-violet-e2d6ff)',
        'palet-violet-f1ebff': 'var(--palet-violet-f1ebff)',

        /* Fresh Air */
        'fresh-air-99d5ff': 'var(--fresh-air-99d5ff)',
        'fresh-air-adddff': 'var(--fresh-air-adddff)',
        'fresh-air-c2e5ff': 'var(--fresh-air-c2e5ff)',
        'fresh-air-d6edff': 'var(--fresh-air-d6edff)',
        'fresh-air-ebf6ff': 'var(--fresh-air-ebf6ff)',

        /* Royal Orange */
        'royal-orange-ff8b1f': 'var(--royal-orange-ff8b1f)',
        'royal-orange-ff9f47': 'var(--royal-orange-ff9f47)',
        'royal-orange-ffab5c': 'var(--royal-orange-ffab5c)',
        'royal-orange-ffc085': 'var(--royal-orange-ffc085)',
        'royal-orange-ffd5ad': 'var(--royal-orange-ffd5ad)',

        /* Mid Aquamarine */
        'mid-aquamarine-1fff80': 'var(--mid-aquamarine-1fff80)',
        'mid-aquamarine-70ffae': 'var(--mid-aquamarine-70ffae)',
        'mid-aquamarine-85ffba': 'var(--mid-aquamarine-85ffba)',
        'mid-aquamarine-adffd1': 'var(--mid-aquamarine-adffd1)',
        'mid-aquamarine-d6ffe8': 'var(--mid-aquamarine-d6ffe8)',

        /* Red */
        'red-f34236': 'var(--red-f34236)',
        'red-ff4a3e': 'var(--red-ff4a3e)',
        'red-ff5e53': 'var(--red-ff5e53)',
        'red-ff776e': 'var(--red-ff776e)',
        'red-ff9993': 'var(--red-ff9993)',

        border: 'hsl(var(--border))',
        input: 'hsl(var(--input))',
        ring: 'hsl(var(--ring))',
        background: 'hsl(var(--background))',
        foreground: 'hsl(var(--foreground))',
        primary: {
          DEFAULT: 'hsl(var(--primary))',
          foreground: 'hsl(var(--primary-foreground))'
        },
        secondary: {
          DEFAULT: 'hsl(var(--secondary))',
          foreground: 'hsl(var(--secondary-foreground))'
        },
        destructive: {
          DEFAULT: 'hsl(var(--destructive))',
          foreground: 'hsl(var(--destructive-foreground))'
        },
        muted: {
          DEFAULT: 'hsl(var(--muted))',
          foreground: 'hsl(var(--muted-foreground))'
        },
        accent: {
          DEFAULT: 'hsl(var(--accent))',
          foreground: 'hsl(var(--accent-foreground))'
        },
        popover: {
          DEFAULT: 'hsl(var(--popover))',
          foreground: 'hsl(var(--popover-foreground))'
        },
        card: {
          DEFAULT: 'hsl(var(--card))',
          foreground: 'hsl(var(--card-foreground))'
        }
      },
      borderRadius: {
        lg: 'var(--radius)',
        md: 'calc(var(--radius) - 2px)',
        sm: 'calc(var(--radius) - 4px)'
      },
      keyframes: {
        'accordion-down': {
          from: { height: '0' },
          to: { height: 'var(--radix-accordion-content-height)' }
        },
        'accordion-up': {
          from: { height: 'var(--radix-accordion-content-height)' },
          to: { height: '0' }
        }
      },
      animation: {
        'accordion-down': 'accordion-down 0.2s ease-out',
        'accordion-up': 'accordion-up 0.2s ease-out'
      }
    }
  },
  plugins: [require('tailwindcss-animate')]
};
