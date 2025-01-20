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
        className="bg-red-500 dark:bg-red-600 text-white px-3 py-1 rounded-md hover:bg-red-600 dark:hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 dark:focus:ring-red-600 transition-colors"
        onClick={handleDelete}
      >
        Delete
      </button>

      {error && (
        <div className="text-red-600 dark:text-red-400 mt-1">
          Error deleting column: {error.message}
        </div>
      )}
    </div>
  );
};
