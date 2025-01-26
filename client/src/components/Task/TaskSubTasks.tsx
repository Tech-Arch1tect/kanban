import { useState } from "react";
import { ModelsTask } from "../../typescript-fetch-client";
import { useCreateTask } from "../../hooks/tasks/useCreateTask";
import { ClipboardIcon } from "@heroicons/react/24/outline";

export const TaskSubTasks = ({ task }: { task: ModelsTask }) => {
  const [subtaskTitle, setSubtaskTitle] = useState("");
  const { mutate: createTask } = useCreateTask();

  const handleCreateSubtask = () => {
    if (!subtaskTitle.trim()) return;
    createTask({
      parentTaskId: task.id,
      boardId: task.boardId,
      title: subtaskTitle,
      description: "",
      status: "open",
      swimlaneId: task.swimlaneId,
      columnId: task.columnId,
      assigneeId: task.assigneeId,
    });
    setSubtaskTitle("");
  };

  return (
    <div className="space-y-4">
      <h2 className="text-2xl font-bold text-gray-700 dark:text-gray-200 mb-4">
        Task Subtasks
      </h2>

      {task.subtasks?.length ? (
        <ul className="space-y-4">
          {task.subtasks.map((subtask) => (
            <li
              key={subtask.id}
              className="p-4 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-700 flex items-center space-x-3"
            >
              <ClipboardIcon className="w-5 h-5 text-gray-400 dark:text-gray-500" />
              <div className="flex-1 text-gray-700 dark:text-gray-200">
                <div className="font-semibold">{subtask.title}</div>
                <div className="text-sm text-gray-500 dark:text-gray-400">
                  Assigned to {subtask.assignee?.username || "no one"}
                </div>
              </div>
            </li>
          ))}
        </ul>
      ) : (
        <div className="text-gray-500 dark:text-gray-400 text-center">
          No subtasks available.
        </div>
      )}

      <div className="flex space-x-2">
        <input
          type="text"
          placeholder="New subtask title"
          value={subtaskTitle}
          onChange={(e) => setSubtaskTitle(e.target.value)}
          className="flex-1 p-2 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-700 text-gray-700 dark:text-gray-200"
        />
        <button
          onClick={handleCreateSubtask}
          className="bg-blue-500 hover:bg-blue-700 text-white p-2 rounded-md"
        >
          Add Subtask
        </button>
      </div>
    </div>
  );
};
