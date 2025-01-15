import { useQuery } from "@tanstack/react-query";
import { tasksApi } from "../../lib/api";
import { toast } from "react-toastify";

export const useGetTaskQueryAllBoards = (query: string) => {
  const { data, isLoading, error } = useQuery({
    queryKey: ["tasksAllBoards", query],
    queryFn: async () => {
      try {
        return await tasksApi.apiV1TasksQueryAllBoardsPost({
          request: {
            query: query,
          },
        });
      } catch (error) {
        toast.error("Failed to fetch tasks");
        throw error;
      }
    },
    enabled: !!query,
    retry: false,
  });

  return { data, isLoading, error };
};
