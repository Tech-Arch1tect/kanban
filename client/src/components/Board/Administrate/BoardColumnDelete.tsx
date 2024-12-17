import { useDeleteColumn } from "../../../hooks/boards/administration/useDeleteColumn";
import { ModelsColumn } from "../../../typescript-fetch-client";

export const BoardColumnDelete = ({ column }: { column: ModelsColumn }) => {
  const { deleteColumn, error } = useDeleteColumn();

  const handleDelete = () => {
    deleteColumn(column);
  };

  return (
    <div>
      <button
        className="bg-red-500 text-white px-2 py-1 rounded-md"
        onClick={handleDelete}
      >
        Delete
      </button>
      {error && <div>Error deleting column {error.message}</div>}
    </div>
  );
};
