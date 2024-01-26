/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [ "./src/**/*.templ", "./dist/404.html"],
  darkMode: 'class',
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: [
      'night', 'winter'
    ]
  },
}

