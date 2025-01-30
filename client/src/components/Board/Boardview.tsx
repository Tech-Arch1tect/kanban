import { useEffect, useState } from "react";
import { useParams } from "@tanstack/react-router";
import { ModelsSwimlane } from "../../typescript-fetch-client";
import BoardSwimlanes from "./BoardSwimlanes";
import { useGetTaskQuery } from "../../hooks/tasks/useTaskQuery";
import AdminLinks from "./AdminLinks";
import { useBoardDataBySlug } from "../../hooks/boards/useBoardDataBySlug";

export default function BoardView() {
  const { slug } = useParams({ from: "/boards/$slug" });
  const { data, isLoading, error } = useBoardDataBySlug(slug);

  const [taskQuery, setTaskQuery] = useState("status:open");

  const {
    data: tasks,
    isLoading: tasksLoading,
    error: tasksError,
  } = useGetTaskQuery(taskQuery, data?.board?.id ?? 0);

  if (isLoading)
    return (
      <div className="flex justify-center items-center h-screen text-gray-600 dark:text-gray-400">
        Loading...
      </div>
    );

  if (error)
    return (
      <div className="flex justify-center items-center h-screen text-red-600 dark:text-red-400">
        Error loading board
      </div>
    );

  const board = data?.board;

  return (
    <div className="p-6 bg-gray-50 dark:bg-gray-900 min-h-screen">
      {/* Header Section */}
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-3xl font-bold text-gray-800 dark:text-gray-200">
          {board?.name}
        </h1>
        <AdminLinks board={board!} />
      </div>

      {/* Task Filter Section */}
      <div className="mb-8 bg-white dark:bg-gray-800 p-6 rounded-lg shadow-sm dark:shadow-md">
        <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
          Filter tasks by query:
        </label>
        <div className="flex items-center gap-4">
          <input
            type="text"
            value={taskQuery}
            onChange={(e) => setTaskQuery(e.target.value)}
            className="w-full p-2 border border-gray-300 dark:border-gray-600 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 focus:border-transparent bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
            placeholder='e.g. "status:open"'
          />
          {tasksLoading && (
            <div className="text-gray-500 dark:text-gray-400">
              Loading tasks...
            </div>
          )}
          {tasksError && (
            <div className="text-red-500 dark:text-red-400">
              Error loading tasks
            </div>
          )}
        </div>
        {tasks && (
          <div className="mt-4 text-sm text-gray-600 dark:text-gray-400">
            Found{" "}
            <strong className="text-blue-600 dark:text-blue-400">
              {tasks?.tasks?.length ?? 0}
            </strong>{" "}
            tasks for query
            <strong className="text-blue-600 dark:text-blue-400">
              {" "}
              "{taskQuery}"
            </strong>
          </div>
        )}
      </div>

      {/* Board Columns and Swimlanes */}
      <div className="">
        {board?.swimlanes?.map((swimlane: ModelsSwimlane) => (
          <BoardSwimlanes
            key={swimlane.id}
            swimlane={swimlane}
            columns={board?.columns ?? []}
            tasks={tasks?.tasks ?? []}
          />
        ))}
      </div>
    </div>
  );
}
