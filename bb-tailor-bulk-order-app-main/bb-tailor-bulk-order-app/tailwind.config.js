/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
        colors: {
            // Brooks Bingham Main Colors
            primaryBrand: '#FFFFFF',
            secondaryBrand: '#0E355B',
            success: '#117F42',
            error: '#E12C2C',

            // Text Colors
            primaryText: '#0E355B',
            secondaryText: '#4B5563',
            border: '#E8E8E8',

            // Button Colors
            buttonDefault: '#0E355B',
            buttonHover: '#154676',
            buttonText: '#D1D5DB',
        },
        fontFamily: {
            nunito: ['Nunito', 'sans-serif'],
        },
        fontSize: {
            heading: ['48px', '52px'],
            subheading: ['20px', '28px'],
            body1: ['24px', '32px'],
            body2: ['16px', '24px'],
            button: ['22px', '52px'],
            label: ['16px', '28px'],
        },
        letterSpacing: {
            normal: '0px',
        },
    },
  },
  plugins: [],
}
