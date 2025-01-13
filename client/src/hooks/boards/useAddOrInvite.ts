import { toast } from "react-toastify";
import { boardsApi } from "../../lib/api";
import {
  ApiV1BoardsAddOrInvitePostRequest,
  BoardAddOrInviteUserToBoardRequestRoleEnum,
  BoardAddOrInviteUserToBoardResponse,
} from "../../typescript-fetch-client";
import { useMutation } from "@tanstack/react-query";

export const useAddOrInvite = () => {
  const { mutate, error, isError, isSuccess, data, isPending } = useMutation<
    BoardAddOrInviteUserToBoardResponse,
    Error,
    ApiV1BoardsAddOrInvitePostRequest
  >({
    mutationFn: async ({
      request: { boardId, email, role },
    }: ApiV1BoardsAddOrInvitePostRequest) =>
      boardsApi.apiV1BoardsAddOrInvitePost({
        request: {
          boardId,
          email,
          role: role as BoardAddOrInviteUserToBoardRequestRoleEnum,
        },
      }),
    onSuccess: () => {
      toast.success("Invite sent successfully!");
    },
    onError: (error: any) => {
      toast.error(error.message || "Failed to send invite.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
