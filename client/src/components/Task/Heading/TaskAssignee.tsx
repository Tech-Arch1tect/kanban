import { useState } from "react";
import { ModelsTask, ModelsUser } from "../../../typescript-fetch-client";
import { useEditTask } from "../../../hooks/tasks/useEditTask";
import { useGetUsersWithAccessToBoard } from "../../../hooks/boards/useGetUsersWithAccessToBoard";

export function TaskAssignee({ task }: { task: ModelsTask }) {
  const [isEditing, setIsEditing] = useState(false);
  const { mutate: editTask } = useEditTask();
  const { data: users, isLoading } = useGetUsersWithAccessToBoard({
    id: task.boardId || 0,
  });

  const handleAssigneeChange = (newAssigneeId: string) => {
    if (!task.id || !task.title || !task.description || !task.status) {
      return;
    }
    editTask({
      id: task.id,
      title: task.title,
      description: task.description,
      status: task.status,
      assigneeId: Number(newAssigneeId),
    });
    setIsEditing(false);
  };

  if (isLoading) return <span>Loading assignees...</span>;

  return (
    <div className="flex items-center">
      <span className="font-medium text-gray-600 w-24">Assignee:</span>
      {isEditing ? (
        <select
          className="border p-1 text-gray-800"
          defaultValue={task?.assignee?.id || ""}
          onChange={(e) => handleAssigneeChange(e.target.value)}
        >
          <option value="">Unassigned</option>
          {users?.users?.map((user: ModelsUser) => (
            <option key={user.id} value={user.id}>
              {user.username}
            </option>
          ))}
        </select>
      ) : (
        <span
          className="text-gray-800 cursor-pointer underline"
          onClick={() => setIsEditing(true)}
        >
          {task?.assignee?.username || "Unassigned"}
        </span>
      )}
    </div>
  );
}
