import { useState } from "react";
import { ModelsTask } from "../../typescript-fetch-client";
import { useCreateTask } from "../../hooks/tasks/useCreateTask";
import { useEditTask } from "../../hooks/tasks/useEditTask";
import {
  ClipboardIcon,
  PencilIcon,
  CheckIcon,
  XMarkIcon,
  PlusIcon,
} from "@heroicons/react/24/outline";
import { useGetUsersWithAccessToBoard } from "../../hooks/boards/useGetUsersWithAccessToBoard";
import { Link } from "@tanstack/react-router";

export const TaskSubTasks = ({ task }: { task: ModelsTask }) => {
  const [subtaskTitle, setSubtaskTitle] = useState("");
  const [assigneeId, setAssigneeId] = useState<number | null>(null);
  const [editingSubtask, setEditingSubtask] = useState<number | null>(null);
  const [newSubtaskTitle, setNewSubtaskTitle] = useState("");
  const { mutate: createTask } = useCreateTask();
  const { mutate: editTask } = useEditTask();
  const { data: users, isLoading: usersLoading } = useGetUsersWithAccessToBoard(
    { id: task.boardId as number }
  );

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
      assigneeId: assigneeId || task.assigneeId,
    });
    setSubtaskTitle("");
    setAssigneeId(null);
  };

  const handleAssigneeChange = (
    subtask: ModelsTask,
    newAssigneeId: number | null
  ) => {
    editTask({
      id: subtask.id as number,
      title: subtask.title as string,
      description: subtask.description || " ",
      status: subtask.status as string,
      assigneeId: newAssigneeId as number,
    });
  };

  const handleEditSubtaskTitle = (subtask: ModelsTask) => {
    setEditingSubtask(subtask.id as number);
    setNewSubtaskTitle(subtask.title as string);
  };

  const handleSaveSubtaskTitle = (subtask: ModelsTask) => {
    if (!newSubtaskTitle.trim()) return;
    editTask({
      id: subtask.id as number,
      title: newSubtaskTitle,
      description: subtask.description || " ",
      status: subtask.status as string,
      assigneeId: subtask.assigneeId as number,
    });
    setEditingSubtask(null);
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
                {editingSubtask === subtask.id ? (
                  <div className="flex items-center space-x-2">
                    <input
                      type="text"
                      value={newSubtaskTitle}
                      onChange={(e) => setNewSubtaskTitle(e.target.value)}
                      className="p-2 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-700 text-gray-700 dark:text-gray-200"
                    />
                    <button
                      onClick={() => handleSaveSubtaskTitle(subtask)}
                      className="bg-green-500 hover:bg-green-700 text-white p-2 rounded-md flex items-center"
                    >
                      <CheckIcon className="w-5 h-5" />
                    </button>
                    <button
                      onClick={() => setEditingSubtask(null)}
                      className="bg-red-500 hover:bg-red-700 text-white p-2 rounded-md flex items-center"
                    >
                      <XMarkIcon className="w-5 h-5" />
                    </button>
                  </div>
                ) : (
                  <>
                    {/* @ts-ignore */}
                    <Link to={`/task/${subtask.id as number}`}>
                      <div className="font-semibold">{subtask.title}</div>
                    </Link>
                    <div className="text-sm text-gray-500 dark:text-gray-400">
                      Assigned to {subtask.assignee?.username || "no one"}
                    </div>
                    <button
                      onClick={() => handleEditSubtaskTitle(subtask)}
                      className="text-blue-500 hover:text-blue-700 text-sm flex items-center"
                    >
                      <PencilIcon className="w-5 h-5 mr-1" />
                    </button>
                  </>
                )}
              </div>
              <select
                value={subtask.assignee?.id || ""}
                onChange={(e) =>
                  handleAssigneeChange(subtask, Number(e.target.value) || null)
                }
                className="p-2 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-700 text-gray-700 dark:text-gray-200"
              >
                <option value="">Unassigned</option>
                {users?.users?.map((user) => (
                  <option key={user.id} value={user.id}>
                    {user.username}
                  </option>
                ))}
              </select>
            </li>
          ))}
        </ul>
      ) : (
        <div className="text-gray-500 dark:text-gray-400 text-center">
          No subtasks available.
        </div>
      )}

      <div className="space-y-2">
        <div className="flex space-x-2">
          <input
            type="text"
            placeholder="New subtask title"
            value={subtaskTitle}
            onChange={(e) => setSubtaskTitle(e.target.value)}
            className="flex-1 p-2 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-700 text-gray-700 dark:text-gray-200"
          />
          <select
            value={assigneeId || ""}
            onChange={(e) => setAssigneeId(Number(e.target.value) || null)}
            className="p-2 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-700 text-gray-700 dark:text-gray-200"
            disabled={usersLoading}
          >
            <option value="">Select Assignee</option>
            {users?.users?.map((user) => (
              <option key={user.id} value={user.id}>
                {user.username}
              </option>
            ))}
          </select>
          <button
            onClick={handleCreateSubtask}
            className="bg-blue-500 hover:bg-blue-700 text-white p-2 rounded-md flex items-center"
          >
            <PlusIcon className="w-5 h-5 mr-1" />
            Add Subtask
          </button>
        </div>
        {usersLoading && (
          <div className="text-gray-500 dark:text-gray-400 text-sm">
            Loading users...
          </div>
        )}
      </div>
    </div>
  );
};
