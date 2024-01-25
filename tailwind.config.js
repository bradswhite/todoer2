/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [ "./src/**/*.templ"],
  darkMode: 'class',
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: [
      'winter', 'night'
    ]
  },
}

