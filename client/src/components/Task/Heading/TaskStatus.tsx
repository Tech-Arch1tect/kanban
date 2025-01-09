import React, { useState } from "react";
import { useUpdateTaskStatus } from "../../../hooks/tasks/useUpdateTaskStatus";
import {
  ModelsTask,
  TaskUpdateTaskStatusRequestStatusEnum,
} from "../../../typescript-fetch-client";

export function TaskStatus({ task }: { task: ModelsTask }) {
  const [isEditing, setIsEditing] = useState(false);
  const [newStatus, setNewStatus] = useState(task?.status || "open");
  const { mutate, isPending, isError, isSuccess } = useUpdateTaskStatus();

  const handleStatusChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedStatus = e.target.value;
    if (selectedStatus === "open" || selectedStatus === "closed") {
      setNewStatus(selectedStatus);
    }
  };

  const handleUpdateStatus = () => {
    if (!task.id) return;
    mutate({
      taskId: task.id,
      status: newStatus as TaskUpdateTaskStatusRequestStatusEnum,
    });
    setIsEditing(false);
  };

  return (
    <div className="flex items-center">
      <span className="font-medium text-gray-600 w-24">Status:</span>
      {isEditing ? (
        <select
          value={newStatus}
          onChange={handleStatusChange}
          onBlur={handleUpdateStatus}
          className="border px-2 py-1 rounded-md"
          disabled={isPending}
          autoFocus
        >
          <option value="open">Open</option>
          <option value="closed">Closed</option>
        </select>
      ) : (
        <span
          className="text-gray-800 cursor-pointer"
          onClick={() => setIsEditing(true)}
        >
          {task?.status}
        </span>
      )}
    </div>
  );
}
