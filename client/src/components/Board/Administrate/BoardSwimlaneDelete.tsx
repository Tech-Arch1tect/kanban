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
    <div>
      <button
        className="bg-red-500 text-white px-2 py-1 rounded-md"
        onClick={handleDelete}
      >
        Delete
      </button>
      {error && <div>Error deleting swimlane {error.message}</div>}
    </div>
  );
};
