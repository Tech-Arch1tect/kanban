import { useState } from "react";
import { ModelsTask, ModelsUser } from "../../../typescript-fetch-client";
import { useGetUsersWithAccessToBoard } from "../../../hooks/boards/useGetUsersWithAccessToBoard";
import { useUpdateTaskAssignee } from "../../../hooks/tasks/useUpdateTaskAssignee";

export function TaskAssignee({ task }: { task: ModelsTask }) {
  const [isEditing, setIsEditing] = useState(false);
  const { mutate: updateTaskAssignee } = useUpdateTaskAssignee();
  const { data: users, isLoading } = useGetUsersWithAccessToBoard({
    id: task.boardId || 0,
  });

  const handleAssigneeChange = (newAssigneeId: string) => {
    if (!task.id) return;
    updateTaskAssignee({
      taskId: task.id,
      assigneeId: Number(newAssigneeId),
    });
    setIsEditing(false);
  };

  if (isLoading) return <span>Loading assignees...</span>;

  return (
    <div className="flex items-center">
      <span className="font-medium text-gray-600 dark:text-gray-400 w-24">
        Assignee:
      </span>
      {isEditing ? (
        <select
          className="border border-gray-300 dark:border-gray-600 p-1 text-gray-800 dark:text-gray-200 bg-white dark:bg-gray-700 rounded-md"
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
          className="text-gray-800 dark:text-gray-200 cursor-pointer underline"
          onClick={() => setIsEditing(true)}
        >
          {task?.assignee?.username || "Unassigned"}
        </span>
      )}
    </div>
  );
}
