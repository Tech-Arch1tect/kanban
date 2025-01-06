import { useQuery } from "@tanstack/react-query";

import { boardsApi } from "../../lib/api";

export const useGetUsersWithAccessToBoard = ({ id }: { id: number }) => {
  const { data, isLoading, error } = useQuery({
    queryKey: ["board-permissions", id],
    queryFn: () => boardsApi.apiV1BoardsPermissionsBoardIdGet({ boardId: id }),
  });
  return { data, isLoading, error };
};
