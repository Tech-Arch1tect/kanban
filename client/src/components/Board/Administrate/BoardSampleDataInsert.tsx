import { useState } from "react";
import { useInsertSampleData } from "../../../hooks/boards/administration/useInsertSampleData";

interface BoardSampleDataInsertProps {
  boardId: string;
}

export function BoardSampleDataInsert({ boardId }: BoardSampleDataInsertProps) {
  const [numTasks, setNumTasks] = useState(100);
  const [numComments, setNumComments] = useState(1000);
  const { mutate, isError, isSuccess, error, data, isPending } =
    useInsertSampleData();

  const handleInsert = () => {
    mutate({ boardId: Number(boardId), numTasks, numComments });
  };

  return (
    <div className="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md space-y-4">
      <h2 className="text-2xl font-semibold text-gray-800 dark:text-gray-200">
        Sample Data
      </h2>

      {isError && (
        <div className="text-red-600 dark:text-red-400">
          Error: {String(error)}
        </div>
      )}

      {isSuccess && (
        <div className="text-green-600 dark:text-green-400">
          Sample data inserted successfully!
        </div>
      )}

      <div className="flex flex-col gap-2">
        <label
          htmlFor="numTasks"
          className="text-gray-700 dark:text-gray-300 font-medium"
        >
          Number of Tasks:
        </label>
        <input
          type="number"
          id="numTasks"
          value={numTasks}
          onChange={(e) => setNumTasks(Number(e.target.value))}
          className="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
        />
      </div>

      <div className="flex flex-col gap-2">
        <label
          htmlFor="numComments"
          className="text-gray-700 dark:text-gray-300 font-medium"
        >
          Number of Comments:
        </label>
        <input
          type="number"
          id="numComments"
          value={numComments}
          onChange={(e) => setNumComments(Number(e.target.value))}
          className="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
        />
      </div>

      <button
        className="px-4 py-2 bg-blue-600 dark:bg-blue-700 text-white rounded-md hover:bg-blue-700 dark:hover:bg-blue-800 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 disabled:opacity-50 transition-colors"
        onClick={handleInsert}
        disabled={isPending}
      >
        {isPending ? "Inserting..." : "Insert Sample Data"}
      </button>
    </div>
  );
}
