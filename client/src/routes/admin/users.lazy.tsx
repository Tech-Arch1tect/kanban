import { createLazyFileRoute } from "@tanstack/react-router";
import { useEffect, useState } from "react";
import { ModelsRole } from "../../typescript-fetch-client";
import {
  AdminAdminListUsersResponse,
  AdminAdminUpdateUserRoleRequestRoleEnum,
} from "../../typescript-fetch-client/index.js";
import { adminApi } from "../../lib/api";

export const Route = createLazyFileRoute("/admin/users")({
  component: () => <AdminUsers />,
});

const AdminUsers = () => {
  const [usersData, setUsersData] = useState<AdminAdminListUsersResponse>({});
  const [currentPage, setCurrentPage] = useState<number>(1);
  const [editingUserId, setEditingUserId] = useState<number | undefined>(
    undefined
  );
  const [searchQuery, setSearchQuery] = useState<string>("");

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const data = await adminApi.apiV1AdminUsersGet({
          page: currentPage,
          pageSize: 10,
          search: searchQuery,
        });
        setUsersData(data);
        console.log(data);
      } catch (error) {
        console.error("Error fetching users:", error);
      }
    };

    fetchUsers();
  }, [currentPage, editingUserId, searchQuery]);

  const handlePageChange = (newPage: number) => {
    setCurrentPage(newPage);
  };

  const handleRoleChange = async (
    userId: any,
    newRole: keyof typeof ModelsRole
  ) => {
    try {
      await adminApi.apiV1AdminUsersIdRolePut({
        id: userId,
        user: {
          role: newRole as AdminAdminUpdateUserRoleRequestRoleEnum,
        },
      });
      setEditingUserId(undefined);
    } catch (error) {
      console.error("Error updating role:", error);
    }
  };

  return (
    <div className="p-4 bg-white dark:bg-gray-800">
      <div className="mb-4">
        <input
          type="text"
          placeholder="Search users..."
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          className="p-2 border border-gray-300 dark:border-gray-600 rounded bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
        />
      </div>
      {usersData.users ? (
        <>
          <table className="min-w-full table-auto border-collapse border border-gray-300 dark:border-gray-600">
            <thead>
              <tr className="bg-gray-100 dark:bg-gray-700">
                <th className="border p-2 text-gray-700 dark:text-gray-200">
                  ID
                </th>
                <th className="border p-2 text-gray-700 dark:text-gray-200">
                  Created At
                </th>
                <th className="border p-2 text-gray-700 dark:text-gray-200">
                  Updated At
                </th>
                <th className="border p-2 text-gray-700 dark:text-gray-200">
                  Email
                </th>
                <th className="border p-2 text-gray-700 dark:text-gray-200">
                  Role
                </th>
              </tr>
            </thead>
            <tbody>
              {usersData.users.map((user) => (
                <tr
                  key={user.id}
                  className="hover:bg-gray-100 dark:hover:bg-gray-700"
                >
                  <td className="border p-2 text-gray-700 dark:text-gray-200">
                    {user.id}
                  </td>
                  <td className="border p-2 text-gray-700 dark:text-gray-200">
                    {user.createdAt
                      ? new Date(user.createdAt).toLocaleString()
                      : "N/A"}
                  </td>
                  <td className="border p-2 text-gray-700 dark:text-gray-200">
                    {user.updatedAt
                      ? new Date(user.updatedAt).toLocaleString()
                      : "N/A"}
                  </td>
                  <td className="border p-2 text-gray-700 dark:text-gray-200">
                    {user.email}
                  </td>
                  <td className="border p-2 text-gray-700 dark:text-gray-200">
                    {editingUserId === user.id ? (
                      <select
                        value={user.role}
                        onChange={(e) => {
                          if (!user.id) {
                            console.error("User ID is required");
                            return;
                          }
                          console.log(e.target.value);
                          handleRoleChange(
                            user.id,
                            e.target.value as
                              | "RoleUser"
                              | "RoleAdmin"
                              | "RoleDisabled"
                          );
                        }}
                        onBlur={() => setEditingUserId(undefined)}
                        autoFocus
                        className="bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200 border border-gray-300 dark:border-gray-600 rounded"
                      >
                        <option value="admin">Admin</option>
                        <option value="user">User</option>
                        <option value="disabled">Disabled</option>
                      </select>
                    ) : (
                      <span
                        onClick={() => setEditingUserId(user.id)}
                        className="cursor-pointer"
                      >
                        {user.role}
                      </span>
                    )}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
          <div className="mt-4">
            {Array.from(
              { length: usersData.totalPages || 1 },
              (_, i) => i + 1
            ).map((pageNumber) => (
              <button
                key={pageNumber}
                onClick={() => handlePageChange(pageNumber)}
                disabled={currentPage === pageNumber}
                className={`px-2 py-1 mx-1 rounded ${
                  currentPage === pageNumber
                    ? "bg-gray-300 dark:bg-gray-600"
                    : "bg-blue-500 dark:bg-blue-700 text-white"
                }`}
              >
                {pageNumber}
              </button>
            ))}
          </div>
        </>
      ) : (
        <p className="mt-4 text-gray-700 dark:text-gray-200">Loading...</p>
      )}
    </div>
  );
};
