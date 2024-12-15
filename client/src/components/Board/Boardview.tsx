import { useState } from "react";
import { useParams } from "@tanstack/react-router";
import { useBoardData } from "../../hooks/boards/useBoardData";
import BoardColumns from "./BoardColumns";
import { ModelsSwimlane } from "../../typescript-fetch-client";
import BoardSwimlanes from "./BoardSwimlanes";

export default function BoardView() {
  const { boardId } = useParams({ from: "/boards/$boardId" });
  const { data, isLoading, error } = useBoardData(boardId);

  const [taskQuery, setTaskQuery] = useState("status:open");

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error loading board</div>;

  const board = data?.board;

  return (
    <div className="p-4">
      <h1 className="text-2xl font-bold mb-4">{board?.name}</h1>

      <div className="mb-4">
        <label className="mr-2 font-semibold">Filter tasks by query:</label>
        <input
          type="text"
          value={taskQuery}
          onChange={(e) => setTaskQuery(e.target.value)}
          className="border p-1"
          placeholder='e.g. "status:open"'
        />
      </div>

      <BoardColumns columns={board?.columns ?? []} />
      {board?.swimlanes?.map((swimlane: ModelsSwimlane) => (
        <BoardSwimlanes
          key={swimlane.id}
          swimlane={swimlane}
          columns={board?.columns ?? []}
        />
      ))}
    </div>
  );
}
