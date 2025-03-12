import { useQuery } from "@tanstack/react-query";
import { boardsApi } from "../../lib/api";

export const useCanAdministrate = (boardId: number) => {
  const {
    data: canAdministrate,
    error,
    isLoading,
  } = useQuery({
    queryKey: ["canAdministrate", boardId],
    queryFn: async () => {
      return await boardsApi.apiV1BoardsCanAdministrateBoardIdGet({
        boardId: boardId,
      });
    },
  });

  return { canAdministrate, error, isLoading };
};
