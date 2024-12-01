import React, { useState } from "react";
import { authApi } from "../lib/api";

const ForgotPassword = () => {
  const [email, setEmail] = useState("");
  const [code, setCode] = useState("");
  const [password, setPassword] = useState("");
  const [stage, setStage] = useState(1); // 1: Email input, 2: Code & Password input
  const [error, setError] = useState("");
  const [success, setSuccess] = useState(false);

  const submitStageOne = async (email: string) => {
    try {
      const response = await authApi.apiV1AuthPasswordResetPost({
        passwordReset: {
          email: email,
        },
      });
      setStage(2);
      setError("");
    } catch (error) {
      setError("Failed to send reset code. Please try again.");
    }
  };

  const submitStageTwo = async (
    email: string,
    code: string,
    password: string
  ) => {
    try {
      const response = await authApi.apiV1AuthResetPasswordPost({
        resetPassword: {
          email: email,
          code: code,
          password: password,
        },
      });
      setSuccess(true);
      setError("");
    } catch (error) {
      setError("Failed to reset password. Please try again.");
    }
  };

  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    if (stage === 1) {
      submitStageOne(email);
    } else if (stage === 2) {
      submitStageTwo(email, code, password);
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
