import { useMutation, useQueryClient } from "@tanstack/react-query";

import { boardsApi } from "../../lib/api";
import { BoardRemovePendingInviteResponse } from "../../typescript-fetch-client";
import { toast } from "react-toastify";

export const useDeletePendingInvite = ({ id }: { id: number }) => {
  const queryClient = useQueryClient();

  return useMutation<BoardRemovePendingInviteResponse>({
    mutationFn: async () =>
      await boardsApi.apiV1BoardsRemovePendingInviteInviteIdPost({
        inviteId: id,
      }),
    onSuccess: (response) => {
      queryClient.invalidateQueries({
        queryKey: ["board-pending-invites", response?.invite?.boardId],
      });
      toast.success("Invite removed successfully!");
    },
    onError: () => {
      toast.error("Failed to remove invite!");
    },
  });
};
