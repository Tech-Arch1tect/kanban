import { useQuery } from "@tanstack/react-query";
import { tasksApi } from "../../lib/api";

export const useGetTaskQuery = (query: string) => {
  const { data, isLoading, error } = useQuery({
    queryKey: ["tasks", query],
    queryFn: async () => {
      try {
        return await tasksApi.apiV1TasksGetQueryQueryGet({ query });
      } catch (error) {
        throw new Error("Failed to fetch tasks");
      }
    },
  });

  return { data, isLoading, error };
};
