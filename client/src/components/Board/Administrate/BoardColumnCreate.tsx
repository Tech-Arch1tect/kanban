import { useState } from "react";
import { useCreateColumn } from "../../../hooks/boards/administration/useCreateColumn";

export const BoardColumnCreate = ({ boardId }: { boardId: string }) => {
  const { mutate, error } = useCreateColumn();
  const [columnName, setColumnName] = useState("");
  const [message, setMessage] = useState("");

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (!columnName) {
      setMessage("Column name is required");
      return;
    }
    mutate({ name: columnName, boardId: Number(boardId) });
    setMessage("Column created");
    setColumnName("");
  };

  return (
    <form
      onSubmit={handleSubmit}
      className="space-y-4 p-4 border border-gray-200 rounded-md bg-gray-50"
    >
      <div>
        <label
          htmlFor="columnName"
          className="block mb-1 font-medium text-gray-700"
        >
          Column Name
        </label>
        <input
          id="columnName"
          type="text"
          value={columnName}
          onChange={(e) => setColumnName(e.target.value)}
          className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          placeholder="Enter column name"
        />
      </div>
      <button
        type="submit"
        className="inline-flex items-center px-4 py-2 text-white bg-blue-600 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
      >
        Create Column
      </button>
      {message && <p className="text-green-600">{message}</p>}
      {error && <p className="text-red-600">{error.message}</p>}
    </form>
  );
};
