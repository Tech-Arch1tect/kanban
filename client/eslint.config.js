import globals from "globals";
import pluginJs from "@eslint/js";
import tseslint from "typescript-eslint";
import pluginReact from "eslint-plugin-react";

/** @type {import('eslint').Linter.Config[]} */
export default [
  {
    settings: {
      react: {
        version: "detect",
      },
    },
  },
  { files: ["**/*.{js,mjs,cjs,ts,jsx,tsx}"] },
  { languageOptions: { globals: globals.browser } },
  pluginJs.configs.recommended,
  ...tseslint.configs.recommended,
  pluginReact.configs.flat.recommended,
  { ignores: ["dist", "src/typescript-fetch-client/*"] },
  { rules: { "react/react-in-jsx-scope": "off" } },
  {
    files: ["**/*.{ts,tsx}"],
    rules: {
      "react/prop-types": "off",
    },
  },
];
