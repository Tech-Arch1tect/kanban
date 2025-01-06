import React, { useState, useCallback } from "react";
import { useTaskData } from "../../hooks/tasks/useTaskData";
import { useUserProfile } from "../../hooks/profile/useUserProfile";
import { useEditComment } from "../../hooks/tasks/Comments/useEditComment";
import { useCreateComment } from "../../hooks/tasks/Comments/useCreateComment";
import { useDeleteComment } from "../../hooks/tasks/Comments/useDeleteComment";
import CommentForm from "./Comments/CommentForm";
import CommentItem from "./Comments/CommentItem";
import { ModelsUser } from "../../typescript-fetch-client";

export default function TaskComments({ taskId }: { taskId: number }) {
  const { data, isLoading, error } = useTaskData({ id: taskId });
  const { profile } = useUserProfile();
  const { mutate: deleteComment } = useDeleteComment();
  const { mutate: editComment } = useEditComment();
  const { mutate: createComment } = useCreateComment();
  const [newComment, setNewComment] = useState("");

  const handleCommentSubmit = useCallback(
    (e: React.FormEvent) => {
      e.preventDefault();
      if (newComment.trim()) {
        createComment({ text: newComment, taskId });
        setNewComment("");
      }
    },
    [newComment, taskId, createComment]
  );

  const handleEdit = useCallback(
    (commentId: number, text: string) => {
      editComment({ id: commentId, text, taskId });
    },
    [taskId, editComment]
  );

  const handleDelete = useCallback(
    (commentId: number) => {
      deleteComment({ id: commentId, taskId });
    },
    [taskId, deleteComment]
  );

  if (isLoading)
    return (
      <div className="flex justify-center items-center py-8">
        <div className="animate-pulse w-3/4 h-6 bg-gray-300 rounded-md"></div>
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
        {data?.task?.comments?.length ? (
          data.task.comments.map((comment) => (
            <CommentItem
              key={comment.id}
              comment={comment}
              profile={profile as ModelsUser}
              onEdit={handleEdit}
              onDelete={handleDelete}
            />
          ))
        ) : (
          <div className="text-gray-500 text-center">
            No comments available.
          </div>
        )}
      </div>
      <CommentForm
        onSubmit={handleCommentSubmit}
        value={newComment}
        setValue={setNewComment}
        placeholder="Write your comment here..."
      />
    </div>
  );
}
