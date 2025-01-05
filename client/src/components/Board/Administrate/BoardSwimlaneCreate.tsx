import { useState } from "react";
import { useCreateSwimlane } from "../../../hooks/boards/administration/useCreateSwimlane";

export const BoardSwimlaneCreate = ({ boardId }: { boardId: string }) => {
  const { mutate, error } = useCreateSwimlane();
  const [swimlaneName, setSwimlaneName] = useState("");
  const [message, setMessage] = useState("");

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (!swimlaneName) {
      setMessage("Swimlane name is required");
      return;
    }
    mutate({ name: swimlaneName, boardId: Number(boardId) });
    setMessage("Swimlane created");
    setSwimlaneName("");
  };

  return (
    <form
      onSubmit={handleSubmit}
      className="space-y-4 p-4 border border-gray-200 rounded-md bg-gray-50"
    >
      <div>
        <label
          htmlFor="swimlaneName"
          className="block mb-1 font-medium text-gray-700"
        >
          Swimlane Name
        </label>
        <input
          id="swimlaneName"
          type="text"
          value={swimlaneName}
          onChange={(e) => setSwimlaneName(e.target.value)}
          className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          placeholder="Enter swimlane name"
        />
      </div>
      <button
        type="submit"
        className="inline-flex items-center px-4 py-2 text-white bg-blue-600 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
      >
        Create Swimlane
      </button>
      {message && <p className="text-green-600">{message}</p>}
      {error && <p className="text-red-600">{error.message}</p>}
    </form>
  );
};
