import React, { useState } from "react";
import { authApi } from "../lib/api";
import { useNavigate } from "@tanstack/react-router";

const AuthForm = ({ mode: initialMode }: { mode: "login" | "register" }) => {
  const [mode, setMode] = useState(initialMode);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [totpCode, setTotpCode] = useState("");
  const [error, setError] = useState("");
  const [totpRequired, setTotpRequired] = useState(false);
  const navigate = useNavigate();

  const handleTotp = async () => {
    try {
      const totp = await authApi.apiV1AuthTotpConfirmPost({
        request: {
          code: totpCode,
        },
      });
      if (totp.message === "totp_confirmed") {
        navigate({ to: "/" });
      } else {
        setError("Invalid TOTP code");
      }
    } catch (error) {
      setError("Invalid TOTP code");
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      if (mode === "login") {
        const auth = await authApi.apiV1AuthLoginPost({
          user: {
            email,
            password,
          },
        });
        if (auth.message === "totp_required") {
          setTotpRequired(true);
        } else {
          navigate({ to: "/" });
        }
      } else if (mode === "register") {
        const auth = await authApi.apiV1AuthRegisterPost({
          user: {
            email,
            password,
          },
        });
        navigate({ to: "/login" });
      }
    } catch (error) {
      setError(
        mode === "login" ? "Invalid email or password" : "Registration failed"
      );
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50">
      <form
        onSubmit={handleSubmit}
        className="bg-white p-8 rounded shadow-md w-96"
      >
        <div className="mb-4 text-center">
          <button
            type="button"
            onClick={() => setMode(mode === "login" ? "register" : "login")}
            className="text-indigo-600 hover:underline"
          >
            {mode === "login"
              ? "Not got an account yet? Register"
              : "Got an account already? Login"}
          </button>
        </div>
        <div className="mb-4">
          <label htmlFor="email" className="block text-gray-700">
            Email
          </label>
          <input
            type="email"
            id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="mt-1 block w-full border-gray-300 rounded shadow-sm focus:border-indigo-500 focus:ring focus:ring-indigo-500"
            required
          />
        </div>
        <div className="mb-4">
          <label htmlFor="password" className="block text-gray-700">
            Password
          </label>
          <input
            type="password"
            id="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="mt-1 block w-full border-gray-300 rounded shadow-sm focus:border-indigo-500 focus:ring focus:ring-indigo-500"
            required
          />
        </div>
        {mode === "login" && totpRequired && (
          <div className="mb-4">
            <label htmlFor="totpCode" className="block text-gray-700">
              Please enter your 2fa code from your authenticator app
            </label>
            <input
              type="text"
              id="totpCode"
              value={totpCode}
              onChange={(e) => setTotpCode(e.target.value)}
              className="mt-1 block w-full border-gray-300 rounded shadow-sm focus:border-indigo-500 focus:ring focus:ring-indigo-500"
            />
          </div>
        )}
        {error && <p className="text-red-500 mb-4 text-center">{error}</p>}
        {mode === "login" && totpRequired ? (
          <button
            type="button"
            onClick={handleTotp}
            className="w-full bg-indigo-600 text-white py-2 px-4 rounded hover:bg-indigo-700 transition duration-300"
          >
            Confirm TOTP
          </button>
        ) : (
          <button
            type="submit"
            className="w-full bg-indigo-600 text-white py-2 px-4 rounded hover:bg-indigo-700 transition duration-300"
          >
            {mode === "login" ? "Login" : "Register"}
          </button>
        )}
      </form>
    </div>
  );
};

export default AuthForm;
