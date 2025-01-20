import { useState } from "react";
import { useChangePassword } from "../../hooks/useChangePassword";

const ChangePassword = () => {
  const [currentPassword, setCurrentPassword] = useState("");
  const [newPassword, setNewPassword] = useState("");
  const [confirmNewPassword, setConfirmNewPassword] = useState("");
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");

  const changePasswordMutation = useChangePassword();

  const handleChangePassword = () => {
    setError("");
    setSuccess("");

    if (!currentPassword || !newPassword || !confirmNewPassword) {
      setError("All fields are required.");
      return;
    }

    if (newPassword !== confirmNewPassword) {
      setError("New password and confirm password do not match.");
      return;
    }

    changePasswordMutation.mutate(
      { currentPassword, newPassword },
      {
        onSuccess: () => {
          setSuccess("Password changed successfully!");
          setCurrentPassword("");
          setNewPassword("");
          setConfirmNewPassword("");
        },
        onError: () => {
          setError("An error occurred while changing the password.");
        },
      }
    );
  };

  return (
    <div className="max-w-md mx-auto p-6 bg-white dark:bg-gray-800 shadow-lg rounded-lg">
      <h2 className="text-2xl font-bold mb-6 text-gray-900 dark:text-gray-200">
        Change Password
      </h2>
      {error && (
        <div className="text-red-500 dark:text-red-400 mb-4">{error}</div>
      )}
      {success && (
        <div className="text-green-500 dark:text-green-400 mb-4">{success}</div>
      )}
      <div className="mb-4">
        <label className="block text-gray-700 dark:text-gray-200 text-sm font-bold mb-2">
          Current Password:
          <input
            type="password"
            value={currentPassword}
            onChange={(e) => setCurrentPassword(e.target.value)}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 dark:text-gray-200 leading-tight focus:outline-none focus:shadow-outline bg-white dark:bg-gray-700 border-gray-300 dark:border-gray-600"
          />
        </label>
      </div>
      <div className="mb-4">
        <label className="block text-gray-700 dark:text-gray-200 text-sm font-bold mb-2">
          New Password:
          <input
            type="password"
            value={newPassword}
            onChange={(e) => setNewPassword(e.target.value)}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 dark:text-gray-200 leading-tight focus:outline-none focus:shadow-outline bg-white dark:bg-gray-700 border-gray-300 dark:border-gray-600"
          />
        </label>
      </div>
      <div className="mb-6">
        <label className="block text-gray-700 dark:text-gray-200 text-sm font-bold mb-2">
          Confirm New Password:
          <input
            type="password"
            value={confirmNewPassword}
            onChange={(e) => setConfirmNewPassword(e.target.value)}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 dark:text-gray-200 leading-tight focus:outline-none focus:shadow-outline bg-white dark:bg-gray-700 border-gray-300 dark:border-gray-600"
          />
        </label>
      </div>
      <button
        onClick={handleChangePassword}
        className="bg-blue-500 hover:bg-blue-700 dark:bg-blue-600 dark:hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline transition duration-300"
      >
        Change Password
      </button>
    </div>
  );
};

export default ChangePassword;
