import { useDeleteColumn } from "../../../hooks/boards/administration/useDeleteColumn";
import { ModelsColumn } from "../../../typescript-fetch-client";

export const BoardColumnDelete = ({ column }: { column: ModelsColumn }) => {
  const { deleteColumn, error } = useDeleteColumn();

  const handleDelete = () => {
    deleteColumn(column);
  };

  return (
    <div className="inline-block">
      <button
        className="bg-red-500 text-white px-3 py-1 rounded-md hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-red-500"
        onClick={handleDelete}
      >
        Delete
      </button>
      {error && (
        <div className="text-red-600 mt-1">
          Error deleting column {error.message}
        </div>
      )}
    </div>
  );
};
