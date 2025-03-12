import { Link } from "@tanstack/react-router";
import { ModelsBoard } from "../../typescript-fetch-client";
import { useCanAdministrate } from "../../hooks/boards/useCanAdministrate";

export default function AdminLinks({ board }: { board: ModelsBoard }) {
  const { canAdministrate } = useCanAdministrate(board.id!);

  if (!canAdministrate) return null;

  return (
    <div className="flex gap-3">
      <Link
        className="px-4 py-2 bg-blue-500 dark:bg-blue-600 text-white rounded-lg hover:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 focus:ring-offset-2 dark:focus:ring-offset-gray-800 transition-colors"
        params={{ boardId: board.id!.toString() }}
        to="/boards/administrate/$boardId"
      >
        Administrate
      </Link>
    </div>
  );
}
