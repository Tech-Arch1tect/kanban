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

  if (isLoading) return <div>Loading...</div>;

  if (error) return <div>Error loading board</div>;

  const board = data?.board;

  return (
    <div className="p-4">
      <h1 className="text-2xl font-bold mb-4">{board?.name}</h1>
      <AdminLinks board={board!} />

      <div className="mb-4">
        <label className="mr-2 font-semibold">Filter tasks by query:</label>
        <input
          type="text"
          value={taskQuery}
          onChange={(e) => setTaskQuery(e.target.value)}
          className="border p-1"
          placeholder='e.g. "status:open"'
        />
        {tasksLoading && <div>Loading tasks...</div>}
        {tasksError && <div>Error loading tasks</div>}
        {tasks && (
          <div>
            Found <strong>{tasks?.tasks?.length ?? 0}</strong> tasks for query
            <strong> "{taskQuery}"</strong>
          </div>
        )}
      </div>

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
  );
}
