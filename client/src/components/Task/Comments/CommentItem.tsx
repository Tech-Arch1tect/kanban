import React, { useState, memo } from "react";
import RenderMarkdown from "../../Utility/RenderMarkdown";
import { ModelsComment, ModelsUser } from "../../../typescript-fetch-client";

interface CommentItemProps {
  comment: ModelsComment;
  profile: ModelsUser;
  onEdit: (commentId: number, text: string) => void;
  onDelete: (commentId: number) => void;
}

const CommentItem: React.FC<CommentItemProps> = ({
  comment,
  profile,
  onEdit,
  onDelete,
}) => {
  const [isEditing, setIsEditing] = useState(false);
  const [editText, setEditText] = useState(comment.text || "");

  const handleEditSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (editText.trim()) {
      if (comment.id) {
        onEdit(comment.id, editText);
        setIsEditing(false);
      }
    }
  };

  return (
    <div className="p-4 border border-gray-200 rounded-lg bg-white shadow-sm">
      {isEditing ? (
        <form onSubmit={handleEditSubmit} className="space-y-3">
          <textarea
            className="w-full border border-gray-300 rounded-md p-3 text-gray-800 focus:ring-2 focus:ring-blue-500 focus:outline-none placeholder-gray-400"
            value={editText}
            onChange={(e) => setEditText(e.target.value)}
            rows={3}
          ></textarea>
          <div className="flex space-x-2">
            <button
              type="submit"
              className="py-2 px-4 bg-blue-500 text-white rounded-md hover:bg-blue-600"
            >
              Save
            </button>
            <button
              type="button"
              className="py-2 px-4 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400"
              onClick={() => setIsEditing(false)}
            >
              Cancel
            </button>
          </div>
        </form>
      ) : (
        <div>
          <div className="text-gray-800">
            <RenderMarkdown
              markdown={comment.text || ""}
              className="prose-sm"
            />
          </div>
          <div className="mt-3 text-sm text-gray-600 flex justify-between items-center">
            <span className="font-medium text-gray-700">
              {comment.user?.username || "Unknown User"}
            </span>
            <div className="text-gray-400">
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
                className="text-blue-500 hover:text-blue-600 text-sm"
              >
                Edit
              </button>
              <button
                onClick={() => onDelete(comment.id ?? 0)}
                className="text-red-500 hover:text-red-600 text-sm"
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
