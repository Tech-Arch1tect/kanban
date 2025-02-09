import React, { useState, useCallback } from "react";
import { useTaskData } from "../../hooks/tasks/useTaskData";
import { useUserProfile } from "../../hooks/profile/useUserProfile";
import { useEditComment } from "../../hooks/tasks/Comments/useEditComment";
import { useCreateComment } from "../../hooks/tasks/Comments/useCreateComment";
import { useDeleteComment } from "../../hooks/tasks/Comments/useDeleteComment";
import CommentForm from "./Comments/CommentForm";
import CommentItem from "./Comments/CommentItem";
import { ModelsUser } from "../../typescript-fetch-client";
import { ModelsTask } from "../../typescript-fetch-client";

export default function TaskComments({ task }: { task: ModelsTask }) {
  if (!task.id) return null;
  const { data, isLoading, error } = useTaskData({ id: task.id as number });
  const { profile } = useUserProfile();
  const { mutate: deleteComment } = useDeleteComment();
  const { mutate: editComment } = useEditComment();
  const { mutate: createComment } = useCreateComment();
  const [newComment, setNewComment] = useState("");

  const handleCommentSubmit = useCallback(
    (e: React.FormEvent) => {
      e.preventDefault();
      if (newComment.trim()) {
        createComment({ text: newComment, taskId: task.id as number });
        setNewComment("");
      }
    },
    [newComment, task.id, createComment]
  );

  const handleEdit = useCallback(
    (commentId: number, text: string) => {
      editComment({ id: commentId, text, taskId: task.id as number });
    },
    [task.id, editComment]
  );

  const handleDelete = useCallback(
    (commentId: number) => {
      deleteComment({ id: commentId, taskId: task.id as number });
    },
    [task.id, deleteComment]
  );

  if (isLoading)
    return (
      <div className="flex justify-center items-center py-8">
        <div className="animate-pulse w-3/4 h-6 bg-gray-300 dark:bg-gray-600 rounded-md"></div>
      </div>
    );

  if (error)
    return (
      <div className="text-red-500 dark:text-red-400 font-medium text-center py-4">
        Error: {error.message}
      </div>
    );

  return (
    <div className="mx-auto space-y-6">
      <h2 className="text-2xl font-bold text-gray-700 dark:text-gray-200 flex items-center space-x-2">
        <span>Comments</span>
        <span className="bg-gray-200 text-gray-700 dark:bg-gray-700 dark:text-gray-200 px-2 py-1 rounded text-sm">
          {task.comments?.length || 0}
        </span>
      </h2>
      <div className="space-y-4">
        {data?.task?.comments?.length ? (
          data.task.comments.map((comment) => (
            <CommentItem
              key={comment.id}
              comment={comment}
              profile={profile as ModelsUser}
              onEdit={handleEdit}
              onDelete={handleDelete}
              boardId={task.boardId as number}
            />
          ))
        ) : (
          <div className="text-gray-500 dark:text-gray-400 text-center">
            No comments available.
          </div>
        )}
      </div>
      <CommentForm
        onSubmit={handleCommentSubmit}
        value={newComment}
        setValue={setNewComment}
        placeholder="Write your comment here..."
        boardId={task.boardId as number}
      />
    </div>
  );
}
