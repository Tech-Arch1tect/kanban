import React, { useState, useEffect } from "react";
import { useSendResetCode } from "../../hooks/auth/useSendResetCode";
import { useResetPassword } from "../../hooks/auth/useResetPassword";
import { Link } from "@tanstack/react-router";

const ForgotPassword = ({ initialEmail = "", initialCode = "" }) => {
  const [email, setEmail] = useState(initialEmail);
  const [code, setCode] = useState(initialCode);
  const [password, setPassword] = useState("");
  const [stage, setStage] = useState(initialCode ? 2 : 1);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState(false);

  const sendResetCodeMutation = useSendResetCode();
  const resetPasswordMutation = useResetPassword();

  useEffect(() => {
    if (initialEmail) setEmail(initialEmail);
    if (initialCode) {
      setCode(initialCode);
      setStage(2);
    }
  }, [initialEmail, initialCode]);

  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    setError("");

    if (stage === 1) {
      sendResetCodeMutation.mutate(email, {
        onSuccess: () => {
          setStage(2);
        },
        onError: () => {
          setError("Failed to send reset code. Please try again.");
        },
      });
    } else if (stage === 2) {
      resetPasswordMutation.mutate(
        { email, code, password },
        {
          onSuccess: () => {
            setSuccess(true);
          },
          onError: () => {
            setError("Failed to reset password. Please try again.");
          },
        }
      );
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50 dark:bg-gray-900">
      <form
        onSubmit={handleSubmit}
        className="bg-white dark:bg-gray-800 p-8 rounded shadow-md w-96"
      >
        {success ? (
          <div>
            <h2 className="text-2xl font-bold mb-4 text-center text-gray-900 dark:text-gray-200">
              Password Reset Successful
            </h2>
            <p className="text-gray-700 dark:text-gray-300 text-center">
              Your password has been successfully reset.
            </p>
            <p className="text-center pt-4">
              <Link
                to="/login"
                className="text-indigo-600 dark:text-indigo-400 hover:underline"
              >
                Login
              </Link>
            </p>
          </div>
        ) : (
          <>
            <h2 className="text-2xl font-bold mb-4 text-center text-gray-900 dark:text-gray-200">
              Forgot Password
            </h2>
            {stage === 1 ? (
              <div className="mb-4">
                <label
                  htmlFor="email"
                  className="block text-gray-700 dark:text-gray-200"
                >
                  Email
                </label>
                <input
                  type="email"
                  id="email"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  className="mt-1 block w-full border-gray-300 dark:border-gray-600 rounded shadow-sm focus:border-indigo-500 focus:ring focus:ring-indigo-500 dark:focus:ring-indigo-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
                  required
                />
              </div>
            ) : (
              <>
                <div className="mb-4">
                  <label
                    htmlFor="code"
                    className="block text-gray-700 dark:text-gray-200"
                  >
                    Code
                  </label>
                  <input
                    type="text"
                    id="code"
                    value={code}
                    onChange={(e) => setCode(e.target.value)}
                    className="mt-1 block w-full border-gray-300 dark:border-gray-600 rounded shadow-sm focus:border-indigo-500 focus:ring focus:ring-indigo-500 dark:focus:ring-indigo-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
                    required
                  />
                </div>
                <div className="mb-4">
                  <label
                    htmlFor="password"
                    className="block text-gray-700 dark:text-gray-200"
                  >
                    New Password
                  </label>
                  <input
                    type="password"
                    id="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    className="mt-1 block w-full border-gray-300 dark:border-gray-600 rounded shadow-sm focus:border-indigo-500 focus:ring focus:ring-indigo-500 dark:focus:ring-indigo-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
                    required
                  />
                </div>
              </>
            )}
            {error && (
              <p className="text-red-500 dark:text-red-400 mb-4 text-center">
                {error}
              </p>
            )}
            <button
              type="submit"
              className="w-full bg-indigo-600 dark:bg-indigo-700 text-white py-2 px-4 rounded hover:bg-indigo-700 dark:hover:bg-indigo-800 transition duration-300"
            >
              {stage === 1 ? "Send Reset Code" : "Reset Password"}
            </button>
          </>
        )}
      </form>
    </div>
  );
};

export default ForgotPassword;
