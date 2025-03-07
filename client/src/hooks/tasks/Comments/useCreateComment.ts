import { useMutation, useQueryClient } from "@tanstack/react-query";
import { commentsApi } from "../../../lib/api";
import { ModelsComment } from "../../../typescript-fetch-client/";
import { toast } from "react-toastify";

export const useCreateComment = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (comment: ModelsComment) => {
      return await commentsApi
        .apiV1CommentsCreatePost({
          request: {
            text: comment.text,
            taskId: comment.taskId,
          },
        })
        .then(() => {
          queryClient.invalidateQueries({
            queryKey: ["task", comment.taskId],
          });
        });
    },
    onSuccess: () => {
      toast.success("Comment added successfully!");
    },
    onError: (err: Error) => {
      toast.error(err.message || "Failed to create comment.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
