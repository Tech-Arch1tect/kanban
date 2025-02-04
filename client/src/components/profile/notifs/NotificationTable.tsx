import React from "react";
import { useNotificationConfigurationData } from "../../../hooks/notifications/useNotificationConfigurationData";
import { useDeleteNotificationConfiguration } from "../../../hooks/notifications/useDeleteNotificationConfiguration";

export const NotificationTable = () => {
  const {
    data: notificationsData,
    isLoading: notificationsLoading,
    error: notificationsError,
  } = useNotificationConfigurationData();

  const { mutate: deleteNotification, isPending: deleting } =
    useDeleteNotificationConfiguration();

  return (
    <div className="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-sm dark:shadow-md">
      <h2 className="text-2xl font-semibold text-gray-800 dark:text-gray-100 mb-4">
        Existing Notification Configurations
      </h2>
      {notificationsLoading && (
        <p className="text-gray-600 dark:text-gray-400">
          Loading notifications…
        </p>
      )}
      {notificationsError && (
        <p className="text-red-600 dark:text-red-400">
          Error loading notifications.
        </p>
      )}
      {notificationsData &&
      notificationsData.notifications &&
      notificationsData.notifications.length > 0 ? (
        <div className="overflow-x-auto">
          <table className="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
            <thead className="bg-gray-100 dark:bg-gray-700">
              <tr>
                <th className="px-4 py-2 text-left text-sm font-medium text-gray-700 dark:text-gray-300">
                  ID
                </th>
                <th className="px-4 py-2 text-left text-sm font-medium text-gray-700 dark:text-gray-300">
                  Name
                </th>
                <th className="px-4 py-2 text-left text-sm font-medium text-gray-700 dark:text-gray-300">
                  Method
                </th>
                <th className="px-4 py-2 text-left text-sm font-medium text-gray-700 dark:text-gray-300">
                  Email/Webhook
                </th>
                <th className="px-4 py-2 text-left text-sm font-medium text-gray-700 dark:text-gray-300">
                  Events
                </th>
                <th className="px-4 py-2 text-left text-sm font-medium text-gray-700 dark:text-gray-300">
                  Boards
                </th>
                <th className="px-4 py-2 text-left text-sm font-medium text-gray-700 dark:text-gray-300">
                  Only Assignee
                </th>
                <th className="px-4 py-2 text-left text-sm font-medium text-gray-700 dark:text-gray-300">
                  Created At
                </th>
                <th className="px-4 py-2 text-left text-sm font-medium text-gray-700 dark:text-gray-300">
                  Actions
                </th>
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-200 dark:divide-gray-700">
              {notificationsData.notifications.map((notification) => (
                <tr key={notification.id}>
                  <td className="px-4 py-2 text-sm text-gray-700 dark:text-gray-200">
                    {notification.id}
                  </td>
                  <td className="px-4 py-2 text-sm text-gray-700 dark:text-gray-200">
                    {notification.name}
                  </td>
                  <td className="px-4 py-2 text-sm text-gray-700 dark:text-gray-200">
                    {notification.method}
                  </td>
                  <td className="px-4 py-2 text-sm text-gray-700 dark:text-gray-200">
                    {notification.method === "email"
                      ? notification.email
                      : notification.webhookUrl}
                  </td>
                  <td className="px-4 py-2 text-sm text-gray-700 dark:text-gray-200">
                    {notification.events &&
                      notification.events.map((e, idx) => (
                        <span key={idx}>
                          {e.name}
                          {idx < (notification.events?.length as number) - 1
                            ? ", "
                            : ""}
                        </span>
                      ))}
                  </td>
                  <td className="px-4 py-2 text-sm text-gray-700 dark:text-gray-200">
                    {notification.boards &&
                      notification.boards.map((board, idx) => (
                        <span key={idx}>
                          {board.name}
                          {idx < (notification.boards?.length as number) - 1
                            ? ", "
                            : ""}
                        </span>
                      ))}
                  </td>
                  <td className="px-4 py-2 text-sm text-gray-700 dark:text-gray-200">
                    {notification.onlyAssignee ? "Yes" : "No"}
                  </td>
                  <td className="px-4 py-2 text-sm text-gray-700 dark:text-gray-200">
                    {notification.createdAt}
                  </td>
                  <td className="px-4 py-2 text-sm">
                    <button
                      className="text-red-600 dark:text-red-400 hover:underline disabled:opacity-50"
                      onClick={() =>
                        deleteNotification(notification.id as number)
                      }
                      disabled={deleting}
                    >
                      Delete
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      ) : (
        <p className="text-gray-600 dark:text-gray-400">
          No notification configurations found.
        </p>
      )}
    </div>
  );
};
