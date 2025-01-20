import { createFileRoute } from "@tanstack/react-router";
import { useBoards } from "../../hooks/boards/useBoards";
import { useState } from "react";
import { useCreateBoard } from "../../hooks/boards/useCreateBoard";

export const Route = createFileRoute("/admin/boards")({
  component: () => <AdminBoards />,
});

const AdminBoards = () => {
  const { boards, error, isLoading } = useBoards();
  const {
    mutate: createBoard,
    isPending: isCreating,
    isError: isCreateError,
  } = useCreateBoard();
  const [newBoardName, setNewBoardName] = useState("");
  const [newBoardSlug, setNewBoardSlug] = useState("");

  const handleCreateBoard = () => {
    createBoard({
      request: {
        name: newBoardName,
        slug: newBoardSlug,
        swimlanes: ["Default"],
        columns: ["Backlog", "Todo", "In Progress", "Done"],
      },
    });
  };

  if (isLoading)
    return (
      <div className="flex justify-center items-center h-full">
        <div className="text-gray-500 dark:text-gray-400 text-lg">
          Loading...
        </div>
      </div>
    );

  if (error)
    return (
      <div className="flex justify-center items-center h-full">
        <div className="text-red-500 dark:text-red-400 text-lg">
          Error: {error.message}
        </div>
      </div>
    );

  return (
    <div className="max-w-3xl mx-auto p-6 bg-white dark:bg-gray-800 shadow-md rounded-md">
      <h1 className="text-2xl font-semibold mb-4 text-gray-800 dark:text-gray-200">
        Admin - Boards
      </h1>

      <div className="flex flex-col sm:flex-row items-start sm:items-center mb-6 space-y-4 sm:space-y-0 sm:space-x-4">
        <input
          value={newBoardName}
          onChange={(e) => setNewBoardName(e.target.value)}
          placeholder="Board name"
          className="w-full sm:w-auto flex-1 px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
        />
        <input
          value={newBoardSlug}
          onChange={(e) => setNewBoardSlug(e.target.value)}
          placeholder="Board slug"
          className="w-full sm:w-auto flex-1 px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
        />
        <button
          onClick={handleCreateBoard}
          disabled={isCreating}
          className={`px-4 py-2 bg-blue-600 dark:bg-blue-700 text-white rounded-md hover:bg-blue-700 dark:hover:bg-blue-800 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 ${
            isCreating ? "opacity-50 cursor-not-allowed" : ""
          }`}
        >
          {isCreating ? "Creating..." : "Create Board"}
        </button>
      </div>

      {isCreateError && (
        <div className="mb-4 text-red-500 dark:text-red-400">
          Error creating board. Please try again.
        </div>
      )}

      <ul className="space-y-2">
        {boards && boards.boards?.length && boards.boards.length > 0 ? (
          boards.boards.map((board) => (
            <li
              key={board.id}
              className="p-4 bg-gray-50 dark:bg-gray-700 border border-gray-200 dark:border-gray-600 rounded-md flex justify-between items-center"
            >
              <span className="text-gray-700 dark:text-gray-200">
                {board.name}
              </span>
            </li>
          ))
        ) : (
          <li className="text-gray-500 dark:text-gray-400">
            No boards available.
          </li>
        )}
      </ul>
    </div>
  );
};
