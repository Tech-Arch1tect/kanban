import { useQuery } from "@tanstack/react-query";
import { boardsApi } from "../../lib/api";

export const useBoardData = (boardId: string) => {
  const { data, error, isLoading } = useQuery({
    queryKey: ["boardData", boardId],
    queryFn: async () => await boardsApi.apiV1BoardsGetIdGet({ id: boardId }),
  });

  return { data, error, isLoading };
};
