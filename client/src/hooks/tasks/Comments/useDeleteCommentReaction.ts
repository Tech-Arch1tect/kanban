import { useMutation, useQueryClient } from "@tanstack/react-query";
import { commentsApi } from "../../../lib/api";
import { toast } from "react-toastify";

export const useDeleteCommentReaction = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (reactionId: number) => {
      return await commentsApi.apiV1CommentsDeleteReactionPost({
        request: {
          reactionId: reactionId,
        },
      });
    },
    onSuccess: (resp) => {
      queryClient.invalidateQueries({
        queryKey: ["task", resp.reaction?.comment?.taskId],
      });
      toast.success("Reaction deleted successfully!");
    },
    onError: (err: Error) => {
      toast.error(err.message || "Failed to delete reaction.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
