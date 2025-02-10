import { useState } from "react";
import { ModelsTask } from "../../typescript-fetch-client";
import RenderMarkdown from "../Utility/RenderMarkdown";
import { useUpdateTaskDescription } from "../../hooks/tasks/useUpdateTaskDescription";
import MentionableTextarea from "../Utility/MentionableTextarea";
import { useBoardUsernames } from "../../hooks/boards/useBoardUsernames";

export function TaskDescription({ task }: { task: ModelsTask }) {
  const [isEditing, setIsEditing] = useState(false);
  const [isPreview, setIsPreview] = useState(false);
  const [description, setDescription] = useState(task?.description || "");
  const { mutate: updateDescription } = useUpdateTaskDescription();

  const { usernames, isLoading } = useBoardUsernames(task.boardId as number);

  const handleEditClick = () => setIsEditing(true);

  const handleSave = () => {
    setIsEditing(false);
    setIsPreview(false);
    if (description !== task.description) {
      updateDescription({ taskId: task.id as number, description });
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
          {isPreview ? (
            <div className="p-2 border rounded bg-gray-50 dark:bg-gray-800">
              <RenderMarkdown
                markdown={description || "No description provided."}
              />
            </div>
          ) : (
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
            </>
          )}
          <div className="mt-2 flex space-x-2">
            {isPreview ? (
              <button
                onClick={() => setIsPreview(false)}
                className="px-3 py-1 text-sm text-white bg-blue-500 rounded hover:bg-blue-600"
              >
                Back to Edit
              </button>
            ) : (
              <button
                onClick={() => setIsPreview(true)}
                className="px-3 py-1 text-sm text-white bg-blue-500 rounded hover:bg-blue-600"
              >
                Preview
              </button>
            )}
            <button
              onClick={handleSave}
              className="px-3 py-1 text-sm text-white bg-green-500 rounded hover:bg-green-600"
            >
              Save
            </button>
            <button
              onClick={() => {
                setIsEditing(false);
                setIsPreview(false);
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
