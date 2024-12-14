import { useQuery } from "@tanstack/react-query";
import { boardsApi } from "../../lib/api";

export const useBoards = () => {
  const {
    data: boards,
    error,
    isLoading,
  } = useQuery({
    queryKey: ["boards"],
    queryFn: async () => {
      return await boardsApi.apiV1BoardsListGet();
    },
  });

  return { boards, error, isLoading };
};
