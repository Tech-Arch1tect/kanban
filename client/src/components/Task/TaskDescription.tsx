import React, { useState } from "react";
import { ModelsTask } from "../../typescript-fetch-client";
import RenderMarkdown from "../Utility/RenderMarkdown";
import { useUpdateTaskDescription } from "../../hooks/tasks/useUpdateTaskDescription";

export function TaskDescription({ task }: { task: ModelsTask }) {
  const [isEditing, setIsEditing] = useState(false);
  const [description, setDescription] = useState(task?.description || "");
  const { mutate: updateDescription } = useUpdateTaskDescription();

  const handleEditClick = () => {
    setIsEditing(true);
  };

  const handleBlur = () => {
    setIsEditing(false);
    if (!task.id) return;
    if (description !== task.description) {
      updateDescription({ taskId: task.id, description });
    }
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();
      setIsEditing(false);
      if (!task.id) return;
      if (description !== task.description) {
        updateDescription({ taskId: task.id, description });
      }
    } else if (e.key === "Escape") {
      setIsEditing(false);
      setDescription(task.description || "");
    }
  };

  return (
    <div>
      <h2 className="text-lg font-semibold text-gray-900">Description</h2>
      {isEditing ? (
        <textarea
          className="w-full border rounded p-2 text-gray-700"
          value={description}
          autoFocus
          onChange={(e) => setDescription(e.target.value)}
          onBlur={handleBlur}
          onKeyDown={handleKeyDown}
          rows={5}
        />
      ) : (
        <div className="flex items-center">
          <div className="text-gray-700 flex-1">
            <RenderMarkdown
              markdown={description || "No description provided."}
            />
          </div>
          <button
            className="ml-2 px-3 py-1 text-sm text-white bg-blue-500 rounded hover:bg-blue-600"
            onClick={handleEditClick}
          >
            Edit
          </button>
        </div>
      )}
    </div>
  );
}
