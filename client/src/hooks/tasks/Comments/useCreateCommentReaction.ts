import { useMutation, useQueryClient } from "@tanstack/react-query";
import { commentsApi } from "../../../lib/api";
import { ModelsReaction } from "../../../typescript-fetch-client";
import { toast } from "react-toastify";

export const useCreateCommentReaction = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (reaction: ModelsReaction) => {
      return await commentsApi
        .apiV1CommentsCreateReactionPost({
          request: {
            commentId: reaction.commentId,
            reaction: reaction.reaction,
          },
        })
        .then(() => {
          queryClient.invalidateQueries({
            queryKey: ["task", reaction.comment?.taskId],
          });
        });
    },
    onSuccess: () => {
      toast.success("Reaction added successfully!");
    },
    onError: (err: any) => {
      toast.error(err.message || "Failed to create reaction.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
