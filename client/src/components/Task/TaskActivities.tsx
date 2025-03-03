import { useState } from "react";
import { ModelsTask, ModelsTaskActivity } from "../../typescript-fetch-client";
import { useTaskActivities } from "../../hooks/tasks/useTaskActivities";
import { friendlyEventNames } from "../profile/notifs/friendlyEventNames";
import { ChevronDownIcon, ChevronUpIcon } from "@heroicons/react/24/outline";

export default function TaskActivities({ task }: { task: ModelsTask }) {
  const [showActivities, setShowActivities] = useState(false);
  const [page, setPage] = useState(1);
  const [pageSize, setPageSize] = useState(10);

  const { data, isLoading, error } = useTaskActivities({
    taskId: task.id as number,
    page,
    pageSize,
  });

  const totalPages = data?.totalPages || 1;

  return (
    <div className="bg-white dark:bg-gray-800 rounded-lg shadow-sm">
      <div className="flex items-center justify-between mb-4">
        <h2 className="text-2xl font-bold text-gray-700 dark:text-gray-200 flex items-center space-x-2">
          <span>Activity Log</span>
          <span className="bg-gray-200 text-gray-700 dark:bg-gray-700 dark:text-gray-200 px-2 py-1 rounded text-sm">
            {data?.totalRecords || 0}
          </span>
        </h2>
        <button
          onClick={() => setShowActivities((prev) => !prev)}
          className="text-blue-500 hover:text-blue-700"
          title={showActivities ? "Hide activities" : "Show activities"}
        >
          {showActivities ? (
            <ChevronUpIcon className="w-6 h-6" />
          ) : (
            <ChevronDownIcon className="w-6 h-6" />
          )}
        </button>
      </div>

      {showActivities && (
        <>
          {isLoading ? (
            <div className="space-y-2">
              {Array.from({ length: pageSize }).map((_, index) => (
                <div
                  key={index}
                  className="animate-pulse p-2 border-b border-gray-300 dark:border-gray-600 last:border-b-0 bg-gray-200 dark:bg-gray-700"
                >
                  <div className="h-3 bg-gray-300 dark:bg-gray-500 rounded w-1/2 mb-1"></div>
                  <div className="h-2 bg-gray-300 dark:bg-gray-500 rounded w-1/3"></div>
                </div>
              ))}
            </div>
          ) : error ? (
            <div className="text-red-500 dark:text-red-400 font-medium">
              Error loading activities.
            </div>
          ) : (
            <div className="space-y-3">
              {data?.taskActivities && data.taskActivities.length > 0 ? (
                <>
                  <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-2 mb-3">
                    <div className="flex items-center gap-2">
                      <button
                        disabled={page === 1}
                        onClick={() => setPage(page - 1)}
                        className="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded transition-colors disabled:opacity-50 focus:outline-none focus:ring-2 focus:ring-blue-400"
                      >
                        Previous
                      </button>
                      <button
                        disabled={page >= totalPages}
                        onClick={() => setPage(page + 1)}
                        className="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded transition-colors disabled:opacity-50 focus:outline-none focus:ring-2 focus:ring-blue-400"
                      >
                        Next
                      </button>
                      <div className="flex items-center gap-2">
                        <label
                          htmlFor="pageSizeSelect"
                          className="text-gray-700 dark:text-gray-300 text-xs"
                        >
                          Entries per page:
                        </label>
                        <select
                          id="pageSizeSelect"
                          value={pageSize}
                          onChange={(e) => {
                            setPageSize(Number(e.target.value));
                            setPage(1);
                          }}
                          className="p-1 border border-gray-300 dark:border-gray-600 rounded bg-white dark:bg-gray-700 text-gray-700 dark:text-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-400 text-xs"
                        >
                          <option value={5}>5</option>
                          <option value={10}>10</option>
                          <option value={20}>20</option>
                          <option value={50}>50</option>
                        </select>
                      </div>
                    </div>
                    <div className="text-gray-700 dark:text-gray-300 text-xs">
                      Page {page} of {totalPages}
                    </div>
                  </div>

                  <div className="flex flex-col">
                    {data.taskActivities.map((activity: ModelsTaskActivity) => (
                      <div
                        key={activity.id}
                        className="pb-2 border-b border-gray-200 dark:border-gray-700 last:border-b-0"
                      >
                        <div className="text-gray-800 dark:text-gray-200 text-sm">
                          {activity.user && (
                            <span className="font-semibold">
                              @{activity.user.username}:{" "}
                            </span>
                          )}
                          <span className="font-semibold">
                            {friendlyEventNames[
                              activity.event as keyof typeof friendlyEventNames
                            ] || activity.event}
                          </span>{" "}
                          <span className="text-xs text-gray-500 dark:text-gray-400">
                            on{" "}
                            {activity.createdAt
                              ? new Date(activity.createdAt).toLocaleString()
                              : "Unknown"}
                          </span>
                        </div>

                        <div className="mt-1 text-xs text-gray-600 dark:text-gray-400 flex flex-col gap-1">
                          {activity.oldData && (
                            <div className="p-1 bg-red-100 dark:bg-red-900 text-red-700 dark:text-red-200 rounded">
                              {activity.oldData}
                            </div>
                          )}
                          {activity.newData && (
                            <div className="p-1 bg-green-100 dark:bg-green-900 text-green-700 dark:text-green-200 rounded">
                              {activity.newData}
                            </div>
                          )}
                        </div>
                      </div>
                    ))}
                  </div>
                </>
              ) : (
                <div className="text-center text-gray-500 dark:text-gray-400 text-xs">
                  No activities found for this task.
                </div>
              )}
            </div>
          )}
        </>
      )}
    </div>
  );
}
