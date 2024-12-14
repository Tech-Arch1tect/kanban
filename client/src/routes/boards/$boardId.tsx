import { createFileRoute, useParams } from "@tanstack/react-router";
import { useBoardData } from "../../hooks/boards/useBoardData";
import { useEffect } from "react";

export const Route = createFileRoute("/boards/$boardId")({
  component: BoardView,
});

function BoardView() {
  const { boardId } = useParams({ from: "/boards/$boardId" });
  const { data, isLoading, error } = useBoardData(boardId);
  useEffect(() => {
    console.log(data);
  }, [boardId, data, isLoading, error]);
  return <div>Board ID: {boardId}</div>;
}
