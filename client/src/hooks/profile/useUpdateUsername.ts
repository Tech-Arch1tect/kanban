import { useMutation } from "@tanstack/react-query";
import { authApi } from "../../lib/api";
import { toast } from "react-toastify";

interface UpdateUsernameArgs {
  username: string;
}

export const useUpdateUsername = () => {
  return useMutation({
    mutationFn: async ({ username }: UpdateUsernameArgs) => {
      return await authApi.apiV1AuthChangeUsernamePost({
        usernameChange: {
          username: username,
        },
      });
    },
    onSuccess: () => {
      toast.success("Username updated successfully!");
    },
    onError: (error: Error) => {
      toast.error(error.message || "Failed to update username.");
    },
  });
};
