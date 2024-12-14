import { useQuery } from "@tanstack/react-query";
import { boardApi } from "../../lib/api";

export const useBoards = () => {
  const {
    data: boards,
    error,
    isLoading,
  } = useQuery({
    queryKey: ["boards"],
    queryFn: async () => {
      return await boardApi.apiV1BoardsListGet();
    },
  });

  return { boards, error, isLoading };
};
