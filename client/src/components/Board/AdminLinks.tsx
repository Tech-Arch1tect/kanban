import { Link } from "@tanstack/react-router";
import { useAuth } from "../../hooks/auth/useAuth";
import { useUserProfile } from "../../hooks/profile/useUserProfile";
import { ModelsBoard } from "../../typescript-fetch-client";

export default function AdminLinks({ board }: { board: ModelsBoard }) {
  const { profile } = useUserProfile();
  const { isAdmin } = useAuth(profile);

  // todo: check if user is admin of the board
  if (!isAdmin) return null;

  return (
    <div className="flex gap-3">
      <Link
        className="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition-colors"
        params={{ boardId: board.id!.toString() }}
        to="/boards/administrate/$boardId"
      >
        Administrate
      </Link>
    </div>
  );
}
