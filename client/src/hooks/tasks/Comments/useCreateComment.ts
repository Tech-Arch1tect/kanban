import { useMutation, useQueryClient } from "@tanstack/react-query";
import { commentsApi } from "../../../lib/api";
import { ModelsComment } from "../../../typescript-fetch-client/models";

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
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
