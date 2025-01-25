import React, { useState } from "react";
import { ModelsTask } from "../../typescript-fetch-client";
import RenderMarkdown from "../Utility/RenderMarkdown";
import { useUpdateTaskDescription } from "../../hooks/tasks/useUpdateTaskDescription";
import MentionableTextarea from "../Utility/MentionableTextarea";
import { useBoardUsernames } from "../../hooks/boards/useBoardUsernames";

export function TaskDescription({ task }: { task: ModelsTask }) {
  const [isEditing, setIsEditing] = useState(false);
  const [description, setDescription] = useState(task?.description || "");
  const { mutate: updateDescription } = useUpdateTaskDescription();

  const { usernames, isLoading } = useBoardUsernames(task.boardId as number);

  const handleEditClick = () => setIsEditing(true);

  const handleSave = () => {
    setIsEditing(false);
    if (description !== task.description) {
      updateDescription({ taskId: task.id as number, description });
    }
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();
      handleSave();
    } else if (e.key === "Escape") {
      e.preventDefault();
      setIsEditing(false);
      setDescription(task.description || "");
    }
  };

  const handleSelectSuggestion = (username: string) => {
    console.log(`Selected username: ${username}`);
  };

  return (
    <div>
      <h2 className="text-lg font-semibold text-gray-900 dark:text-gray-200">
        Description
      </h2>
      {isEditing ? (
        <>
          {!isLoading && (
            <MentionableTextarea
              value={description}
              setValue={setDescription}
              placeholder="Edit the task description..."
              suggestions={usernames}
              onSelectSuggestion={handleSelectSuggestion}
            />
          )}
          <div className="mt-2 flex space-x-2">
            <button
              onClick={handleSave}
              className="px-3 py-1 text-sm text-white bg-blue-500 rounded hover:bg-blue-600"
            >
              Save
            </button>
            <button
              onClick={() => {
                setIsEditing(false);
                setDescription(task.description || "");
              }}
              className="px-3 py-1 text-sm text-gray-700 bg-gray-300 rounded hover:bg-gray-400"
            >
              Cancel
            </button>
          </div>
        </>
      ) : (
        <div className="flex items-center">
          <div className="text-gray-700 dark:text-gray-300 flex-1">
            <RenderMarkdown
              markdown={description || "No description provided."}
            />
          </div>
          <button
            onClick={handleEditClick}
            className="ml-2 px-3 py-1 text-sm text-white bg-blue-500 rounded hover:bg-blue-600"
          >
            Edit
          </button>
        </div>
      )}
    </div>
  );
}
