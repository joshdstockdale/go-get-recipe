/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/**/*.templ"],
  theme: {
    fontFamily: {
      sans: '"SourceSansPro-SemiBold", "Source Sans Pro SemiBold", "Source Sans Pro", sans-serif;',
    },
    borderRadius: {
      none: "0",
      sm: "0.5rem",
      DEFAULT: "1rem",
      md: "1rem",
      lg: "2rem",
      full: "9999px",
    },
    extend: {
      animation: {
        "scale-up": "scale-up 4s linear",
      },
      keyframes: {
        "scale-up": {
          "0%": { transform: "scaleY(0)" },
          "15%": { transform: "scaleY(1)" },
        },
      },
      flexBasis: {
        "1/3": "calc(33.3% - 16px)",
        "2/3": "calc(66.6% - 16px)",
      },
      colors: {
        hunter: {
          DEFAULT: "var(--hunter-color)",
          50: "var(--hunter-color-50)",
          100: "var(--hunter-color-100)",
          200: "var(--hunter-color-200)",
          300: "var(--hunter-color-300)",
          400: "var(--hunter-color-400)",
          500: "var(--hunter-color-500)",
          600: "var(--hunter-color-600)",
          700: "var(--hunter-color-700)",
          800: "var(--hunter-color-800)",
          900: "var(--hunter-color-900)",
        },
      },
    },
  },
  plugins: [],
};
