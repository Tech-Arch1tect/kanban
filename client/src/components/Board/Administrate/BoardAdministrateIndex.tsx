import { useParams } from "@tanstack/react-router";
import { useBoardData } from "../../../hooks/boards/useBoardData";
import BoardDelete from "./BoardDelete";
import { BoardSwimlaneCreate } from "./BoardSwimlaneCreate";
import { BoardSwimlaneList } from "./BoardSwimlaneList";
import { BoardColumnsList } from "./BoardColumnsList";
import { BoardColumnCreate } from "./BoardColumnCreate";
import { BoardSampleDataInsert } from "./BoardSampleDataInsert";
import { BoardAddOrInvite } from "./BoardAddOrInvite";
import { BoardPendingInvites } from "./BoardPendingInvites";
import { BoardUsers } from "./BoardUsers";

export default function BoardAdministrateIndex() {
  const { boardId } = useParams({ from: "/boards/administrate/$boardId" });
  const { data, isLoading, error } = useBoardData(boardId);

  if (isLoading)
    return (
      <div className="p-4 text-gray-600 dark:text-gray-400">Loading...</div>
    );
  if (error)
    return (
      <div className="p-4 text-red-600 dark:text-red-400">
        Error loading board
      </div>
    );

  const board = data?.board;
  if (!board)
    return (
      <div className="p-4 text-gray-600 dark:text-gray-400">
        Board not found
      </div>
    );

  return (
    <div className="p-4 bg-gray-50 dark:bg-gray-900 min-h-screen">
      <div className="max-w-4xl mx-auto p-4 space-y-8 ">
        <h1 className="text-3xl font-bold text-gray-900 dark:text-gray-100">
          {board.name} - Administration
        </h1>

        <div className="bg-white dark:bg-gray-800 p-4 rounded-lg shadow-sm dark:shadow-md space-y-4">
          <BoardSampleDataInsert boardId={boardId} />
        </div>

        <div className="bg-white dark:bg-gray-800 p-4 rounded-lg shadow-sm dark:shadow-md space-y-4">
          <BoardAddOrInvite boardId={Number(boardId)} />
          <BoardPendingInvites boardId={Number(boardId)} />
        </div>

        <div className="bg-white dark:bg-gray-800 p-4 rounded-lg shadow-sm dark:shadow-md space-y-4">
          <BoardUsers boardId={Number(boardId)} />
        </div>

        <div className="bg-white dark:bg-gray-800 p-4 rounded-lg shadow-sm dark:shadow-md space-y-4">
          <BoardDelete board={board} />
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
          <div className="bg-white dark:bg-gray-800 p-4 rounded-lg shadow-sm dark:shadow-md space-y-4">
            <h2 className="text-2xl font-semibold text-gray-900 dark:text-gray-100">
              Swimlanes
            </h2>
            <BoardSwimlaneCreate boardId={boardId} />
            <BoardSwimlaneList board={board} />
          </div>

          <div className="bg-white dark:bg-gray-800 p-4 rounded-lg shadow-sm dark:shadow-md space-y-4">
            <h2 className="text-2xl font-semibold text-gray-900 dark:text-gray-100">
              Columns
            </h2>
            <BoardColumnCreate boardId={boardId} />
            <BoardColumnsList board={board} />
          </div>
        </div>
      </div>
    </div>
  );
}
