import { useMutation, useQueryClient } from "@tanstack/react-query";
import { commentsApi } from "../../../lib/api";
import { ModelsComment } from "../../../typescript-fetch-client/";
import { toast } from "react-toastify";

export const useDeleteComment = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (comment: ModelsComment) => {
      return await commentsApi
        .apiV1CommentsDeletePost({
          request: {
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
      toast.success("Comment deleted successfully!");
    },
    onError: (err: any) => {
      toast.error(err.message || "Failed to delete comment.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
