import { useQuery } from "@tanstack/react-query";
import { tasksApi } from "../../lib/api";

export const useGetTaskQuery = (query: string, boardId: number) => {
  const { data, isLoading, error } = useQuery({
    queryKey: ["tasks", query, boardId],
    queryFn: async () => {
      try {
        return await tasksApi.apiV1TasksGetQueryBoardIdQueryGet({
          query,
          boardId: boardId,
        });
      } catch (error) {
        throw new Error("Failed to fetch tasks");
      }
    },
    enabled: Boolean(boardId),
    retry: false,
  });

  return { data, isLoading, error };
};
