import { useState, useEffect } from "react";
import { ModelsTask } from "../../../typescript-fetch-client";
import { useUpdateTaskDueDate } from "../../../hooks/tasks/useUpdateTaskDueDate";

export function TaskDueDate({ task }: { task: ModelsTask }) {
  const [isEditing, setIsEditing] = useState(false);
  const [newDueDate, setNewDueDate] = useState(
    task?.dueDate ? task.dueDate.slice(0, 16) : undefined
  );
  const { mutate, isPending, isSuccess } = useUpdateTaskDueDate();

  useEffect(() => {
    if (isSuccess) {
      setIsEditing(false);
    }
  }, [isSuccess]);

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (!task.id) return;
    const dueDate = newDueDate ? new Date(newDueDate).toISOString() : undefined;
    mutate({ taskId: task.id, dueDate });
  };

  if (isEditing) {
    return (
      <form onSubmit={handleSubmit} className="flex flex-col gap-2">
        <label
          htmlFor="dueDate"
          className="font-medium text-gray-600 dark:text-gray-400"
        >
          New Due Date:
        </label>
        <input
          id="dueDate"
          type="datetime-local"
          value={newDueDate ?? ""}
          onChange={(e) => setNewDueDate(e.target.value)}
          className="border border-gray-300 rounded p-2"
        />
        <div className="flex gap-2 mt-2">
          <button
            type="submit"
            disabled={isPending}
            className="px-4 py-2 bg-blue-500 text-white rounded"
          >
            {isPending ? "Updating..." : "Update Due Date"}
          </button>
          <button
            type="button"
            onClick={() => setIsEditing(false)}
            className="px-4 py-2 bg-gray-500 text-white rounded"
          >
            Cancel
          </button>
        </div>
      </form>
    );
  }

  return (
    <div className="flex items-center">
      <span className="font-medium text-gray-600 dark:text-gray-400 w-24">
        Due Date:
      </span>
      <span className="text-gray-800 dark:text-gray-200 mr-2">
        {task?.dueDate ? new Date(task?.dueDate).toLocaleString() : "None"}
      </span>
      <button
        onClick={() => setIsEditing(true)}
        className="px-2 py-1 bg-green-500 text-white rounded"
      >
        Edit
      </button>
    </div>
  );
}
