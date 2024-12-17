import { useParams } from "@tanstack/react-router";
import { useBoardData } from "../../../hooks/boards/useBoardData";
import BoardDelete from "./BoardDelete";
import { BoardSwimlaneCreate } from "./BoardSwimlaneCreate";
import { BoardSwimlaneList } from "./BoardSwimlaneList";

export default function BoardAdministrateIndex() {
  const { boardId } = useParams({ from: "/boards/administrate/$boardId" });

  const { data, isLoading, error } = useBoardData(boardId);

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error loading board</div>;

  const board = data?.board;

  if (!board) return <div>Board not found</div>;

  return (
    <div>
      <h1 className="text-2xl font-bold">Board Administration: {board.name}</h1>
      <BoardDelete board={board} />
      <BoardSwimlaneCreate boardId={boardId} />
      <BoardSwimlaneList board={board} />
    </div>
  );
}
