import { useState } from "react";
import {
  ModelsTask,
  TaskCreateTaskLinkRequestLinkTypeEnum,
} from "../../typescript-fetch-client";
import { useGetTaskQueryAllBoards } from "../../hooks/tasks/useTaskQueryAllBoards";
import { useCreateTaskLink } from "../../hooks/tasks/useCreateTaskLink";
import useDebounce from "../../hooks/useDebounce";
import { toast } from "react-toastify";
import { CheckCircleIcon } from "@heroicons/react/24/outline";

export const QueryAndAddTaskLink = ({ task }: { task: ModelsTask }) => {
  const [query, setQuery] = useState("");
  const debouncedQuery = useDebounce(query, 300);
  const [selectedTask, setSelectedTask] = useState<string | null>(null);
  const [linkType, setLinkType] = useState<string>("");

  const {
    data: tasks,
    isLoading,
    error,
  } = useGetTaskQueryAllBoards(debouncedQuery || null);
  const { mutate: createLink, isPending } = useCreateTaskLink();

  const handleCreateLink = () => {
    if (!selectedTask || !linkType) {
      toast.error("Please select a task and specify a link type.");
      return;
    }
    createLink({
      srcTaskId: task.id,
      dstTaskId: Number(selectedTask),
      linkType: linkType as TaskCreateTaskLinkRequestLinkTypeEnum,
    });
  };

  return (
    <div className="mx-auto">
      <form onSubmit={(e) => e.preventDefault()} className="mb-4">
        <input
          type="text"
          placeholder="Search tasks..."
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          className="border border-gray-300 dark:border-gray-600 rounded-md p-2 w-full 
                     focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600
                     bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
        />
      </form>

      {isLoading && (
        <p className="text-gray-500 dark:text-gray-400 mb-2">
          Loading tasks...
        </p>
      )}
      {error && (
        <p className="text-red-500 dark:text-red-400 mb-2">
          Error fetching tasks
        </p>
      )}

      {tasks && tasks.tasks && tasks.tasks.length > 0 && (
        <ul className="space-y-2">
          {tasks.tasks.slice(0, 10).map((foundTask: ModelsTask) => {
            const isSelected = selectedTask === foundTask.id?.toString();
            return (
              <li
                key={foundTask.id}
                onClick={() =>
                  setSelectedTask(foundTask.id?.toString() ?? null)
                }
                className={`cursor-pointer flex items-center justify-between p-2 border rounded-md transition-colors ${
                  isSelected
                    ? "bg-blue-500 text-white border-blue-500"
                    : "bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200 hover:bg-gray-100 dark:hover:bg-gray-600"
                }`}
              >
                <span>{foundTask.title}</span>
                {isSelected && <CheckCircleIcon className="w-5 h-5" />}
              </li>
            );
          })}
        </ul>
      )}

      <div className="mt-4">
        <label className="block text-gray-700 dark:text-gray-200 font-medium mb-1">
          Link Type:
        </label>
        <select
          value={linkType}
          onChange={(e) => setLinkType(e.target.value)}
          className="border border-gray-300 dark:border-gray-600 rounded-md p-2 w-full 
                     focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600
                     bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
        >
          <option value="">Select link type...</option>
          <option value="blocks">Blocks</option>
          <option value="fixes">Fixes</option>
          <option value="depends_on">Depends On</option>
          <option value="blocked_by">Blocked By</option>
          <option value="fixed_by">Fixed By</option>
          <option value="depended_on_by">Depended On By</option>
        </select>
      </div>

      <button
        onClick={handleCreateLink}
        disabled={isPending}
        className="mt-4 px-4 py-2 bg-blue-500 dark:bg-blue-600 text-white rounded-md 
                   hover:bg-blue-600 dark:hover:bg-blue-700 disabled:bg-gray-400 dark:disabled:bg-gray-600
                   disabled:cursor-not-allowed transition-colors"
      >
        {isPending ? "Creating Link..." : "Create Task Link"}
      </button>
    </div>
  );
};
