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
      className="space-y-4 p-6 border border-gray-200 dark:border-gray-600 rounded-md bg-gray-50 dark:bg-gray-800"
    >
      <div>
        <label
          htmlFor="columnName"
          className="block mb-1 font-medium text-gray-700 dark:text-gray-300"
        >
          Column Name
        </label>
        <input
          id="columnName"
          type="text"
          value={columnName}
          onChange={(e) => setColumnName(e.target.value)}
          className="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
          placeholder="Enter column name"
        />
      </div>

      <button
        type="submit"
        className="inline-flex items-center px-4 py-2 text-white bg-blue-600 dark:bg-blue-700 rounded-md hover:bg-blue-700 dark:hover:bg-blue-800 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 dark:focus:ring-blue-600 transition-colors"
      >
        Create Column
      </button>

      {message && (
        <p className="text-green-600 dark:text-green-400">{message}</p>
      )}

      {error && (
        <p className="text-red-600 dark:text-red-400">{error.message}</p>
      )}
    </form>
  );
};
