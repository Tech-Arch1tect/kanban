import React from "react";
import { ModelsBoard } from "../../../typescript-fetch-client";
import { useBoardDelete } from "../../../hooks/boards/administration/useBoardDelete";

interface BoardDeleteProps {
  board: ModelsBoard;
}

const BoardDelete: React.FC<BoardDeleteProps> = ({ board }) => {
  const { deleteBoard, isDeleting, isError, error, isSuccess } =
    useBoardDelete();

  const handleDelete = () => {
    deleteBoard(board.id!);
  };

  return (
    <div className="space-y-2">
      <button
        className="bg-red-600 text-white px-4 py-2 rounded-md hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500"
        onClick={handleDelete}
        disabled={isDeleting}
      >
        {isDeleting ? "Deleting..." : "Delete Board"}
      </button>
      {isError && (
        <div className="text-red-600">
          Error: {error?.message || "Something went wrong."}
        </div>
      )}
      {isSuccess && (
        <div className="text-green-600">Board deleted successfully!</div>
      )}
    </div>
  );
};

export default BoardDelete;
