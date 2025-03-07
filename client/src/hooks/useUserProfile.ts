import { useQuery } from "@tanstack/react-query";
import { authApi } from "../lib/api";

export const useUserProfile = () => {
  const {
    data: profile,
    error,
    isLoading,
  } = useQuery({
    queryKey: ["userProfile"],
    retry: false,
    queryFn: async () => {
      try {
        return await authApi.apiV1AuthProfileGet();
      } catch (error) {
        throw new Error(
          (error as Error).message || "Failed to fetch user profile"
        );
      }
    },
  });

  return { profile, error, isLoading };
};
