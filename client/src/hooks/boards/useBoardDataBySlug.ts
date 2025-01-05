import { useQuery } from "@tanstack/react-query";
import { boardsApi } from "../../lib/api";

export const useBoardDataBySlug = (slug: string) => {
  const { data, error, isLoading } = useQuery({
    queryKey: ["boardDataBySlug", slug],
    queryFn: async () =>
      await boardsApi.apiV1BoardsGetBySlugSlugGet({ slug: slug }),
    retry: false,
  });

  return { data, error, isLoading };
};
