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
    <div className="bg-white p-4 rounded shadow space-y-4">
      <h2 className="text-2xl font-semibold">Sample Data</h2>
      {isError && <div className="text-red-600">Error: {String(error)}</div>}
      {isSuccess && (
        <div className="text-green-600">Sample data inserted successfully!</div>
      )}
      <div className="flex flex-col gap-2">
        <label htmlFor="numTasks">Number of Tasks:</label>
        <input
          type="number"
          id="numTasks"
          value={numTasks}
          onChange={(e) => setNumTasks(Number(e.target.value))}
        />
      </div>
      <div className="flex flex-col gap-2">
        <label htmlFor="numComments">Number of Comments:</label>
        <input
          type="number"
          id="numComments"
          value={numComments}
          onChange={(e) => setNumComments(Number(e.target.value))}
        />
      </div>
      <button
        className="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700 disabled:opacity-50"
        onClick={handleInsert}
        disabled={isPending}
      >
        {isPending ? "Inserting..." : "Insert Sample Data"}
      </button>
    </div>
  );
}
