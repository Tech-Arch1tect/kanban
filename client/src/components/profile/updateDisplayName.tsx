import { useState } from "react";
import { useUpdateDisplayName } from "../../hooks/profile/useUpdateDisplayName";
import { ModelsUser } from "../../typescript-fetch-client";

const UpdateDisplayName = ({ user }: { user: ModelsUser }) => {
  const [displayName, setDisplayName] = useState(user.displayName);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");

  const updateDisplayNameMutation = useUpdateDisplayName();

  const handleUpdateDisplayName = () => {
    setError("");
    setSuccess("");

    if (!displayName) {
      setError("Display name is required.");
      return;
    }

    updateDisplayNameMutation.mutate(
      { displayName },
      {
        onSuccess: () => {
          setSuccess("Display name updated successfully!");
        },
        onError: () => {
          setError("An error occurred while updating the display name.");
        },
      }
    );
  };

  return (
    <div className="max-w-md mx-auto p-6 bg-white shadow-lg rounded-lg">
      <h2 className="text-2xl font-bold mb-6">Update Display Name</h2>
      {error && <div className="text-red-500 mb-4">{error}</div>}
      {success && <div className="text-green-500 mb-4">{success}</div>}
      <div className="mb-4">
        <label className="block text-gray-700 text-sm font-bold mb-2">
          Display Name:
          <input
            type="text"
            value={displayName}
            onChange={(e) => setDisplayName(e.target.value)}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          />
        </label>
      </div>
      <div className="mb-6">
        <button
          onClick={handleUpdateDisplayName}
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
        >
          Update Display Name
        </button>
      </div>
    </div>
  );
};

export default UpdateDisplayName;
