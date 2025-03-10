import React, { useState } from "react";
import { ModelsTask } from "../../../typescript-fetch-client";
import { useUpdateTaskTitle } from "../../../hooks/tasks/useUpdateTaskTitle";
import RenderMarkdown from "../../Utility/RenderMarkdown";

export function TaskTitle({ task }: { task: ModelsTask }) {
  const [isEditing, setIsEditing] = useState(false);
  const [title, setTitle] = useState(task?.title || "");
  const { mutate: updateTitle } = useUpdateTaskTitle();

  const handleTitleClick = () => {
    setIsEditing(true);
  };

  const handleBlur = () => {
    setIsEditing(false);
    if (title !== task.title) {
      if (!task.id) return;
      updateTitle({ taskId: task.id, title });
    }
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter") {
      setIsEditing(false);
      if (title !== task.title) {
        if (!task.id) return;
        updateTitle({ taskId: task.id, title });
      }
    } else if (e.key === "Escape") {
      setIsEditing(false);
      setTitle(task.title || "");
    }
  };

  return isEditing ? (
    <input
      className="text-2xl font-bold text-gray-900 dark:text-gray-200 mb-2 border border-gray-300 dark:border-gray-600 rounded px-2 w-full bg-white dark:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600"
      value={title}
      autoFocus
      onChange={(e) => setTitle(e.target.value)}
      onBlur={handleBlur}
      onKeyDown={handleKeyDown}
    />
  ) : (
    <h1
      className="text-2xl font-bold text-gray-900 dark:text-gray-200 mb-2 cursor-pointer"
      onClick={handleTitleClick}
    >
      <RenderMarkdown markdown={task?.title || ""} />
    </h1>
  );
}
