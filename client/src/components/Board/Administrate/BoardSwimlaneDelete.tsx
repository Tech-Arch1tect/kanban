import { useDeleteSwimlane } from "../../../hooks/boards/administration/useDeleteSwimlane";
import { ModelsSwimlane } from "../../../typescript-fetch-client";

export const BoardSwimlaneDelete = ({
  swimlane,
}: {
  swimlane: ModelsSwimlane;
}) => {
  const { deleteSwimlane, error } = useDeleteSwimlane();

  const handleDelete = () => {
    deleteSwimlane(swimlane);
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
          Error deleting swimlane {error.message}
        </div>
      )}
    </div>
  );
};
