import { useGetUsersWithAccessToBoard } from "../../../hooks/boards/useGetUsersWithAccessToBoard";
import { useRemoveUser } from "../../../hooks/boards/administration/useRemoveUser";
import { useChangeUserRole } from "../../../hooks/boards/administration/useChangeUserRole";
import { BoardChangeBoardRoleRequestRoleEnum } from "../../../typescript-fetch-client";

export const BoardUsers = ({ boardId }: { boardId: number }) => {
  const { data, isLoading, error } = useGetUsersWithAccessToBoard({
    id: boardId,
  });
  const { mutate: removeUser } = useRemoveUser();
  const { mutate: changeUserRole } = useChangeUserRole();

  if (isLoading) {
    return (
      <div className="bg-white dark:bg-gray-800 p-6 rounded shadow text-center text-gray-700 dark:text-gray-300">
        Loading users...
      </div>
    );
  }
  if (error) {
    return (
      <div className="bg-red-100 dark:bg-red-900 p-6 rounded shadow text-red-700 dark:text-red-300">
        Error loading users: {error.message}
      </div>
    );
  }

  const users = data?.users || [];

  const handleRemoveUser = (userId: number) => {
    removeUser({ request: { boardId, userId } });
  };

  const handleChangeUserRole = (userId: number, role: string) => {
    changeUserRole({
      request: {
        boardId,
        userId,
        role: role as BoardChangeBoardRoleRequestRoleEnum,
      },
    });
  };

  return (
    <div className="bg-white dark:bg-gray-800 p-6 rounded shadow space-y-4">
      <h2 className="text-2xl font-semibold text-gray-800 dark:text-gray-200">
        Board Users
      </h2>
      <ul className="space-y-4">
        {users.map((user) => (
          <li
            key={user.id}
            className="border border-gray-200 dark:border-gray-700 p-4 rounded shadow-sm bg-gray-50 dark:bg-gray-700 hover:bg-gray-100 dark:hover:bg-gray-600 transition-colors"
          >
            <div className="flex justify-between items-center">
              <span className="text-gray-700 dark:text-gray-300">
                {user.username}
              </span>
              <div className="flex items-center space-x-4">
                <select
                  value={user.appRole}
                  onChange={(e) =>
                    handleChangeUserRole(user.id || 0, e.target.value)
                  }
                  disabled={user.role === "admin"}
                  className="bg-white dark:bg-gray-800 border border-gray-300 dark:border-gray-600 rounded px-2 py-1 text-gray-700 dark:text-gray-300"
                >
                  <option value="reader">Reader</option>
                  <option value="member">Member</option>
                  <option value="admin">Admin</option>
                </select>
                <button
                  disabled={user.role === "admin"}
                  onClick={() => handleRemoveUser(user.id || 0)}
                  className={`px-4 py-2 rounded ${
                    user.role === "admin"
                      ? "bg-gray-400 text-gray-200 cursor-not-allowed"
                      : "bg-red-600 text-white hover:bg-red-700"
                  }`}
                >
                  Remove
                </button>
              </div>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
};
