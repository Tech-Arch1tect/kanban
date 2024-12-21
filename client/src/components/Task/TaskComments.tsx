import { useState } from "react";
import { useCreateComment } from "../../hooks/tasks/Comments/useCreateComment";
import { useTaskData } from "../../hooks/tasks/useTaskData";
import RenderMarkdown from "../Utility/RenderMarkdown";

export default function TaskComments({ taskId }: { taskId: number }) {
  const { data, isLoading, error } = useTaskData({ id: taskId });
  const {
    mutate,
    isError,
    isSuccess,
    error: mutationError,
  } = useCreateComment();
  const [newComment, setNewComment] = useState("");

  const handleCommentSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (newComment.trim()) {
      mutate({ text: newComment, taskId });
      setNewComment("");
    }
  };

  if (isLoading)
    return (
      <div className="flex justify-center items-center text-gray-500 py-8">
        Loading...
      </div>
    );
  if (error)
    return (
      <div className="text-red-500 font-medium text-center py-4">
        Error: {error.message}
      </div>
    );

  return (
    <div className="mx-auto">
      <h2 className="text-lg font-semibold text-gray-900 pb-1">Comments</h2>
      <div className="space-y-4 mt-4">
        {data?.task?.comments?.length && data?.task?.comments?.length > 0 ? (
          data.task.comments.map((comment) => (
            <div
              key={comment.id}
              className="p-4 border border-gray-200 rounded-lg bg-white shadow-sm"
            >
              <p className="text-gray-800">
                <RenderMarkdown
                  markdown={comment.text || ""}
                  className="prose-sm"
                />
              </p>
              <div className="mt-3 text-sm text-gray-600 flex justify-between items-center">
                <span className="font-medium text-gray-700">
                  {comment.user?.displayName || "Unknown User"}
                </span>
                <div className="text-gray-400">
                  <span className="block">
                    Created:{" "}
                    {comment.createdAt?.toLocaleString() || "Unknown Date"}
                  </span>
                  <span className="block">
                    Updated:{" "}
                    {comment.updatedAt?.toLocaleString() || "Unknown Date"}
                  </span>
                </div>
              </div>
            </div>
          ))
        ) : (
          <div className="text-gray-500 text-center">
            No comments available.
          </div>
        )}
      </div>

      <form onSubmit={handleCommentSubmit} className="mt-6 space-y-3">
        <textarea
          className="w-full border border-gray-300 rounded-md p-3 text-gray-800 focus:ring-2 focus:ring-blue-500 focus:outline-none placeholder-gray-400"
          placeholder="Write your comment here..."
          value={newComment}
          onChange={(e) => setNewComment(e.target.value)}
          rows={4}
        ></textarea>
        <button
          type="submit"
          disabled={!newComment.trim()}
          className={`w-full py-2 px-4 rounded-md text-white ${
            newComment.trim()
              ? "bg-blue-500 hover:bg-blue-600"
              : "bg-gray-300 cursor-not-allowed"
          }`}
        >
          Submit
        </button>
      </form>

      {isError && (
        <div className="text-red-500 font-medium mt-2 text-center">
          Error: {mutationError?.message || "Failed to create comment."}
        </div>
      )}

      {isSuccess && (
        <div className="text-green-500 font-medium mt-2 text-center">
          Comment added successfully!
        </div>
      )}
    </div>
  );
}
