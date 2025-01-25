import React, { useState, memo } from "react";
import RenderMarkdown from "../../Utility/RenderMarkdown";
import { ModelsComment, ModelsUser } from "../../../typescript-fetch-client";
import MentionableTextarea from "../../Utility/MentionableTextarea";
import { useBoardUsernames } from "../../../hooks/boards/useBoardUsernames";

interface CommentItemProps {
  comment: ModelsComment;
  profile: ModelsUser;
  onEdit: (commentId: number, text: string) => void;
  onDelete: (commentId: number) => void;
  boardId: number;
}

const CommentItem: React.FC<CommentItemProps> = ({
  comment,
  profile,
  onEdit,
  onDelete,
  boardId,
}) => {
  const [isEditing, setIsEditing] = useState(false);
  const [editText, setEditText] = useState(comment.text || "");

  const { usernames, isLoading } = useBoardUsernames(boardId);

  const handleEditSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (editText.trim()) {
      if (comment.id) {
        onEdit(comment.id, editText);
        setIsEditing(false);
      }
    }
  };

  const handleSelectSuggestion = (username: string) => {
    console.log(`Selected username: ${username}`);
  };

  return (
    <div className="p-4 border border-gray-200 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 shadow-sm">
      {isEditing ? (
        <form onSubmit={handleEditSubmit} className="space-y-3">
          {!isLoading && (
            <MentionableTextarea
              value={editText}
              setValue={setEditText}
              placeholder="Edit your comment..."
              suggestions={usernames}
              onSelectSuggestion={handleSelectSuggestion}
            />
          )}
          <div className="flex space-x-2">
            <button
              type="submit"
              className="py-2 px-4 bg-blue-500 text-white rounded-md hover:bg-blue-600 dark:bg-blue-600 dark:hover:bg-blue-700"
            >
              Save
            </button>
            <button
              type="button"
              className="py-2 px-4 bg-gray-300 dark:bg-gray-600 text-gray-700 dark:text-gray-200 rounded-md hover:bg-gray-400 dark:hover:bg-gray-500"
              onClick={() => setIsEditing(false)}
            >
              Cancel
            </button>
          </div>
        </form>
      ) : (
        <div>
          <div className="text-gray-800 dark:text-gray-200">
            <RenderMarkdown markdown={comment.text || ""} />
          </div>
          <div className="mt-3 text-sm flex justify-between items-center">
            <span className="font-medium text-gray-700 dark:text-gray-200">
              {comment.user?.username || "Unknown User"}
            </span>
            <div className="text-gray-600 dark:text-gray-300">
              <span className="block">
                Created:{" "}
                {comment.createdAt
                  ? new Date(comment.createdAt).toLocaleString()
                  : "Unknown Date"}
              </span>
              {comment.updatedAt !== comment.createdAt && (
                <span className="block">
                  Updated:{" "}
                  {comment.updatedAt
                    ? new Date(comment.updatedAt).toLocaleString()
                    : "Unknown Date"}
                </span>
              )}
            </div>
          </div>
          {profile?.id === comment.user?.id && (
            <div className="flex space-x-2 mt-3">
              <button
                onClick={() => setIsEditing(true)}
                className="text-blue-500 hover:text-blue-600 dark:text-blue-400 dark:hover:text-blue-500 text-sm"
              >
                Edit
              </button>
              <button
                onClick={() => onDelete(comment.id ?? 0)}
                className="text-red-500 hover:text-red-600 dark:text-red-400 dark:hover:text-red-500 text-sm"
              >
                Delete
              </button>
            </div>
          )}
        </div>
      )}
    </div>
  );
};

export default memo(CommentItem);
