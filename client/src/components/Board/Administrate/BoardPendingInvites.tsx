import React from "react";
import { useGetPendingInvites } from "../../../hooks/boards/useGetPendingInvites";
import { ModelsBoardInvite } from "../../../typescript-fetch-client/models/ModelsBoardInvite";

export const BoardPendingInvites = ({ boardId }: { boardId: number }) => {
  const { data, isLoading, error } = useGetPendingInvites({ id: boardId });

  if (isLoading) {
    return (
      <div className="bg-white dark:bg-gray-800 p-6 rounded shadow text-center text-gray-700 dark:text-gray-300">
        Loading pending invites...
      </div>
    );
  }

  if (error) {
    return (
      <div className="bg-red-100 dark:bg-red-900 p-6 rounded shadow text-red-700 dark:text-red-300">
        Error fetching pending invites: {error.message}
      </div>
    );
  }

  if (!data || !data.invites || data.invites.length === 0) {
    return (
      <div className="bg-yellow-100 dark:bg-yellow-900 p-6 rounded shadow text-yellow-700 dark:text-yellow-300">
        No pending invites for this board.
      </div>
    );
  }

  return (
    <div className="bg-white dark:bg-gray-800 p-6 rounded shadow space-y-4">
      <h3 className="text-2xl font-semibold text-gray-800 dark:text-gray-200">
        Pending Invites
      </h3>

      <ul className="space-y-2">
        {data.invites.map((invite: ModelsBoardInvite, index: number) => (
          <li
            key={index}
            className="border border-gray-200 dark:border-gray-700 p-4 rounded shadow-sm bg-gray-50 dark:bg-gray-700 hover:bg-gray-100 dark:hover:bg-gray-600 transition-colors"
          >
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
          </li>
        ))}
      </ul>
    </div>
  );
};
