import { useEffect, useState } from "react";
import { useParams } from "@tanstack/react-router";
import BoardColumns from "./BoardColumns";
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
      <div className="flex justify-center items-center h-screen text-gray-600">
        Loading...
      </div>
    );

  if (error)
    return (
      <div className="flex justify-center items-center h-screen text-red-600">
        Error loading board
      </div>
    );

  const board = data?.board;

  return (
    <div className="p-6 bg-gray-50 min-h-screen">
      {/* Header Section */}
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-3xl font-bold text-gray-800">{board?.name}</h1>
        <AdminLinks board={board!} />
      </div>

      {/* Task Filter Section */}
      <div className="mb-8 bg-white p-6 rounded-lg shadow-sm">
        <label className="block text-sm font-medium text-gray-700 mb-2">
          Filter tasks by query:
        </label>
        <div className="flex items-center gap-4">
          <input
            type="text"
            value={taskQuery}
            onChange={(e) => setTaskQuery(e.target.value)}
            className="w-full p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder='e.g. "status:open"'
          />
          {tasksLoading && (
            <div className="text-gray-500">Loading tasks...</div>
          )}
          {tasksError && (
            <div className="text-red-500">Error loading tasks</div>
          )}
        </div>
        {tasks && (
          <div className="mt-4 text-sm text-gray-600">
            Found{" "}
            <strong className="text-blue-600">
              {tasks?.tasks?.length ?? 0}
            </strong>{" "}
            tasks for query
            <strong className="text-blue-600"> "{taskQuery}"</strong>
          </div>
        )}
      </div>

      {/* Board Columns and Swimlanes */}
      <div className="">
        <BoardColumns columns={board?.columns ?? []} />
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
