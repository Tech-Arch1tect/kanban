import { useQuery } from "@tanstack/react-query";

import { boardsApi } from "../../lib/api";
import { BoardGetPendingInvitesResponse } from "../../typescript-fetch-client";

export const useGetPendingInvites = ({ id }: { id: number }) => {
  const { data, isLoading, error } = useQuery<BoardGetPendingInvitesResponse>({
    queryKey: ["board-pending-invites", id],
    queryFn: async () =>
      await boardsApi.apiV1BoardsPendingInvitesBoardIdGet({
        boardId: id,
      }),
  });
  return { data, isLoading, error };
};
