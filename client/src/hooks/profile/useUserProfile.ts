import { useQuery } from "@tanstack/react-query";
import { authApi } from "../../lib/api";
import { toast } from "react-toastify";

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
        toast.error("Failed to fetch user profile");
        throw new Error("Failed to fetch user profile");
      }
    },
  });

  return { profile, error, isLoading };
};
