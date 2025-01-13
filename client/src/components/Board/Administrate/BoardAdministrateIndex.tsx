import { useParams } from "@tanstack/react-router";
import { useBoardData } from "../../../hooks/boards/useBoardData";
import BoardDelete from "./BoardDelete";
import { BoardSwimlaneCreate } from "./BoardSwimlaneCreate";
import { BoardSwimlaneList } from "./BoardSwimlaneList";
import { BoardColumnsList } from "./BoardColumnsList";
import { BoardColumnCreate } from "./BoardColumnCreate";
import { BoardSampleDataInsert } from "./BoardSampleDataInsert";
import { BoardAddOrInvite } from "./BoardAddOrInvite";

export default function BoardAdministrateIndex() {
  const { boardId } = useParams({ from: "/boards/administrate/$boardId" });
  const { data, isLoading, error } = useBoardData(boardId);

  if (isLoading) return <div className="p-4">Loading...</div>;
  if (error) return <div className="p-4 text-red-600">Error loading board</div>;

  const board = data?.board;
  if (!board) return <div className="p-4">Board not found</div>;

  return (
    <div className="max-w-4xl mx-auto p-4 space-y-8">
      <h1 className="text-3xl font-bold">{board.name} - Administration</h1>
      <div className="bg-white p-4 rounded shadow space-y-4">
        <BoardSampleDataInsert boardId={boardId} />
      </div>
      <div className="bg-white p-4 rounded shadow space-y-4">
        <BoardAddOrInvite boardId={Number(boardId)} />
      </div>
      <div className="bg-white p-4 rounded shadow space-y-4">
        <BoardDelete board={board} />
      </div>

      <div className="grid grid-cols-2 gap-8">
        <div className="bg-white p-4 rounded shadow space-y-4">
          <h2 className="text-2xl font-semibold">Swimlanes</h2>
          <BoardSwimlaneCreate boardId={boardId} />
          <BoardSwimlaneList board={board} />
        </div>
        <div className="bg-white p-4 rounded shadow space-y-4">
          <h2 className="text-2xl font-semibold">Columns</h2>
          <BoardColumnCreate boardId={boardId} />
          <BoardColumnsList board={board} />
        </div>
      </div>
    </div>
  );
}
