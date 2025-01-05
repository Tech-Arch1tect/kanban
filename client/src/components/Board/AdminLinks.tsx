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
    <div className="flex gap-2">
      <Link
        className="p-2 bg-blue-500 text-white rounded-md"
        params={{ boardId: board.id!.toString() }}
        to="/boards/administrate/$boardId"
      >
        Administrate
      </Link>
    </div>
  );
}
