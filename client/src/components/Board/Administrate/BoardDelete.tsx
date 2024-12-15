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
    <div>
      <button
        className="bg-red-500 text-white px-4 py-2 rounded-md"
        onClick={handleDelete}
        disabled={isDeleting}
      >
        {isDeleting ? "Deleting..." : "Delete Board"}
      </button>

      {isError && (
        <div style={{ color: "red" }}>
          Error: {error?.message || "Something went wrong."}
        </div>
      )}
      {isSuccess && (
        <div style={{ color: "green" }}>Board deleted successfully!</div>
      )}
    </div>
  );
};

export default BoardDelete;
