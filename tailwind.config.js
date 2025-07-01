/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.{templ}"],
  theme: {
    extend: {
      keyframes: {
        blink: {
          "0%, 100%": { opacity: "1" },
          "50%": { opacity: "0" },
        },
      },
      animation: {
        blink: "blink 1s step-start infinite",
      },
    },
  },
  plugins: [],
};
