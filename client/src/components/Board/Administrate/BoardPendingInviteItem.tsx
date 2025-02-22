import { useDeletePendingInvite } from "../../../hooks/boards/useDeletePendingInvite";
import { ModelsBoardInvite } from "../../../typescript-fetch-client/models/ModelsBoardInvite";

export const BoardPendingInviteItem = ({
  invite,
}: {
  invite: ModelsBoardInvite;
}) => {
  const { mutate, isPending } = useDeletePendingInvite({
    id: invite.id as number,
  });

  const handleDelete = () => {
    mutate();
  };

  return (
    <li className="border border-gray-200 dark:border-gray-700 p-4 rounded shadow-sm bg-gray-50 dark:bg-gray-700 hover:bg-gray-100 dark:hover:bg-gray-600 transition-colors">
      <div className="text-gray-700 dark:text-gray-300">
        <span className="font-medium">Email:</span> {invite.email}
      </div>
      <div className="text-gray-700 dark:text-gray-300">
        <span className="font-medium">Sent At:</span>{" "}
        {new Date(invite.createdAt || "").toLocaleString()}
      </div>
      <div className="text-gray-700 dark:text-gray-300">
        <span className="font-medium">Role:</span> {invite.roleName}
      </div>
      <button
        onClick={handleDelete}
        disabled={isPending}
        className="mt-2 px-3 py-1 bg-red-500 hover:bg-red-600 text-white rounded transition-colors"
      >
        {isPending ? "Deleting..." : "Delete Invite"}
      </button>
    </li>
  );
};
