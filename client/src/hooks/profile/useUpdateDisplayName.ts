import { useMutation } from "@tanstack/react-query";
import { authApi } from "../../lib/api";
import { toast } from "react-toastify";

interface UpdateDisplayNameArgs {
  displayName: string;
}

export const useUpdateDisplayName = () => {
  return useMutation({
    mutationFn: async ({ displayName }: UpdateDisplayNameArgs) => {
      return await authApi.apiV1AuthUpdateDisplayNamePost({
        displayName: {
          displayName,
        },
      });
    },
    onSuccess: () => {
      toast.success("Display name updated successfully!");
    },
    onError: (error: any) => {
      toast.error(error.message || "Failed to update display name.");
    },
  });
};
