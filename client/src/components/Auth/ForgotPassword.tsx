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
    <div className="min-h-screen flex items-center justify-center bg-gray-50">
      <form
        onSubmit={handleSubmit}
        className="bg-white p-8 rounded shadow-md w-96"
      >
        {success ? (
          <div>
            <h2 className="text-2xl font-bold mb-4 text-center">
              Password Reset Successful
            </h2>
            <p className="text-gray-700 text-center">
              Your password has been successfully reset.
            </p>
            <p className="text-center pt-4">
              <Link to="/login">Login</Link>
            </p>
          </div>
        ) : (
          <>
            <h2 className="text-2xl font-bold mb-4 text-center">
              Forgot Password
            </h2>
            {stage === 1 ? (
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
            ) : (
              <>
                <div className="mb-4">
                  <label htmlFor="code" className="block text-gray-700">
                    Code
                  </label>
                  <input
                    type="text"
                    id="code"
                    value={code}
                    onChange={(e) => setCode(e.target.value)}
                    className="mt-1 block w-full border-gray-300 rounded shadow-sm focus:border-indigo-500 focus:ring focus:ring-indigo-500"
                    required
                  />
                </div>
                <div className="mb-4">
                  <label htmlFor="password" className="block text-gray-700">
                    New Password
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
              </>
            )}
            {error && <p className="text-red-500 mb-4 text-center">{error}</p>}
            <button
              type="submit"
              className="w-full bg-indigo-600 text-white py-2 px-4 rounded hover:bg-indigo-700 transition duration-300"
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
