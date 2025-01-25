import { toast } from "react-toastify";
import { boardsApi } from "../../../lib/api";
import {
  ApiV1BoardsRemoveUserPostRequest,
  BoardRemoveUserFromBoardResponse,
} from "../../../typescript-fetch-client";
import { useMutation, useQueryClient } from "@tanstack/react-query";

export const useRemoveUser = () => {
  const queryClient = useQueryClient();
  const { mutate, error, isError, isSuccess, data, isPending } = useMutation<
    BoardRemoveUserFromBoardResponse,
    Error,
    ApiV1BoardsRemoveUserPostRequest
  >({
    mutationFn: async ({
      request: { boardId, userId },
    }: ApiV1BoardsRemoveUserPostRequest) =>
      boardsApi.apiV1BoardsRemoveUserPost({
        request: {
          boardId,
          userId,
        },
      }),
    onSuccess: ({ boardId }) => {
      queryClient.invalidateQueries({
        queryKey: ["board-permissions", boardId],
      });
      toast.success("User removed successfully!");
    },
    onError: (error: any) => {
      toast.error(error.message || "Failed to remove user.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
