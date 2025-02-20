import { useState, memo } from "react";
import RenderMarkdown from "../../Utility/RenderMarkdown";
import { ModelsComment, ModelsUser } from "../../../typescript-fetch-client";
import MentionableTextarea from "../../Utility/MentionableTextarea";
import { useBoardUsernames } from "../../../hooks/boards/useBoardUsernames";
import ReactionList from "./ReactionList";
import ReactionPicker from "./ReactionPicker";

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
  const [isPreview, setIsPreview] = useState(false);
  const [editText, setEditText] = useState(comment.text || "");

  const { usernames, isLoading } = useBoardUsernames(boardId);

  const handleEditSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (editText.trim() && comment.id) {
      onEdit(comment.id, editText);
      setIsEditing(false);
      setIsPreview(false);
    }
  };

  const handleSelectSuggestion = (username: string) => {
    console.log(`Selected username: ${username}`);
  };

  return (
    <div className="p-4 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-700 shadow-sm">
      {isEditing ? (
        <form onSubmit={handleEditSubmit} className="space-y-3">
          {isPreview ? (
            <div className="p-2 border rounded bg-gray-50 dark:bg-gray-800">
              {editText.trim() ? (
                <RenderMarkdown markdown={editText} />
              ) : (
                <p>No comment provided.</p>
              )}
            </div>
          ) : (
            !isLoading && (
              <MentionableTextarea
                value={editText}
                setValue={setEditText}
                placeholder="Edit your comment..."
                suggestions={usernames}
                onSelectSuggestion={handleSelectSuggestion}
                containerClassName="mb-4"
                textareaClassName="shadow-sm"
              />
            )
          )}
          <div className="flex space-x-2">
            <button
              type="button"
              onClick={() => setIsPreview((prev) => !prev)}
              disabled={!isPreview && !editText.trim()}
              className={`px-3 py-1 text-sm text-white bg-blue-500 rounded hover:bg-blue-600 ${
                !isPreview && !editText.trim()
                  ? "opacity-50 cursor-not-allowed"
                  : ""
              }`}
            >
              {isPreview ? "Back to Edit" : "Preview"}
            </button>
            <button
              type="submit"
              disabled={!editText.trim()}
              className="py-2 px-4 bg-blue-500 text-white rounded-md hover:bg-blue-600 dark:bg-blue-600 dark:hover:bg-blue-700"
            >
              Save
            </button>
            <button
              type="button"
              onClick={() => {
                setIsEditing(false);
                setIsPreview(false);
                setEditText(comment.text || "");
              }}
              className="py-2 px-4 bg-gray-300 dark:bg-gray-600 text-gray-700 dark:text-gray-200 rounded-md hover:bg-gray-400 dark:hover:bg-gray-500"
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
          <ReactionList reactions={comment.reactions || []} />
          <div className="mt-3 flex justify-between items-center">
            <span className="font-medium text-gray-700 dark:text-gray-200">
              {comment.user?.username || "Unknown User"}
            </span>
            <div className="flex items-center space-x-3">
              <ReactionPicker
                commentId={comment.id as number}
                taskId={comment.taskId as number}
              />
              <div className="text-gray-600 dark:text-gray-300 text-sm">
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
          </div>

          {profile?.id === comment.user?.id && (
            <div className="flex space-x-2 mt-3">
              <button
                onClick={() => setIsEditing(true)}
                className="text-blue-500 hover:text-blue-600 text-sm"
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
