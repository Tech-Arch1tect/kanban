import { useMutation, useQueryClient } from "@tanstack/react-query";
import { commentsApi } from "../../../lib/api";
import { ModelsComment } from "../../../typescript-fetch-client/";
import { toast } from "react-toastify";

export const useEditComment = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (comment: ModelsComment) => {
      return await commentsApi
        .apiV1CommentsEditPost({
          request: {
            text: comment.text,
            id: comment.id,
          },
        })
        .then(() => {
          queryClient.invalidateQueries({
            queryKey: ["task", comment.taskId],
          });
        });
    },
    onSuccess: () => {
      toast.success("Comment updated successfully!");
    },
    onError: (err: any) => {
      toast.error(err.message || "Failed to update comment.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
