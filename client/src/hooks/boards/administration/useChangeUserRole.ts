import { toast } from "react-toastify";
import { boardsApi } from "../../../lib/api";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import {
  ApiV1BoardsChangeRolePostRequest,
  BoardChangeBoardRoleResponse,
} from "../../../typescript-fetch-client";

export const useChangeUserRole = () => {
  const queryClient = useQueryClient();
  const { mutate, error, isError, isSuccess, data, isPending } = useMutation<
    BoardChangeBoardRoleResponse,
    Error,
    ApiV1BoardsChangeRolePostRequest
  >({
    mutationFn: async ({
      request: { boardId, userId, role },
    }: ApiV1BoardsChangeRolePostRequest) =>
      boardsApi.apiV1BoardsChangeRolePost({
        request: {
          boardId,
          userId,
          role,
        },
      }),
    onSuccess: ({ boardId }) => {
      queryClient.invalidateQueries({
        queryKey: ["board-permissions", boardId],
      });
      toast.success("User role changed successfully!");
    },
    onError: (error: any) => {
      toast.error(error.message || "Failed to change user role.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
