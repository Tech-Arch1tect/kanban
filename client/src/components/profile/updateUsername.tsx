import { useState } from "react";
import { useUpdateUsername } from "../../hooks/profile/useUpdateUsername.js";
import { ModelsUser } from "../../typescript-fetch-client/index.js";

const UpdateUsername = ({ user }: { user: ModelsUser }) => {
  const [username, setUsername] = useState(user.username);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");

  const updateUsernameMutation = useUpdateUsername();

  const handleUpdateUsername = () => {
    setError("");
    setSuccess("");

    if (!username) {
      setError("Username is required.");
      return;
    }

    updateUsernameMutation.mutate(
      { username: username },
      {
        onSuccess: () => {
          setSuccess("Username updated successfully!");
        },
        onError: () => {
          setError("An error occurred while updating the username.");
        },
      }
    );
  };

  return (
    <div className="max-w-md mx-auto p-6 bg-white shadow-lg rounded-lg">
      <h2 className="text-2xl font-bold mb-6">Update Username</h2>
      {error && <div className="text-red-500 mb-4">{error}</div>}
      {success && <div className="text-green-500 mb-4">{success}</div>}
      <div className="mb-4">
        <label className="block text-gray-700 text-sm font-bold mb-2">
          Username:
          <input
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          />
        </label>
      </div>
      <div className="mb-6">
        <button
          onClick={handleUpdateUsername}
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
        >
          Update Username
        </button>
      </div>
    </div>
  );
};

export default UpdateUsername;
