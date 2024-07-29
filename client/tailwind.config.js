/** @type {import('tailwindcss').Config} */
export default {
    content: ['./index.html', './src/**/*.{js,ts,jsx,tsx,vue}'],
    theme: {
        extend: {
            fontFamily: {
                apple: ['SF Pro Display', 'Arial', 'sans-serif']
            },
            colors: {
                primary: '#F4BA30',
                F5F5F5: '#F5F5F5'
            },
            boxShadow: {
                'md': '0 10px 40px -15px rgba(0, 0, 0, 0.5)'
            }
        }
    },
    plugins: []
}
