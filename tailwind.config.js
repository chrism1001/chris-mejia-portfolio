/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./internal/**/*.{go,js,templ,html}"],
  theme: {
    extend: {},
    fontFamily: {
      'roboto': ['Roboto', 'sans-serif'],
    },
  },
  plugins: [],
}

