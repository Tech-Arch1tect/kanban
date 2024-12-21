import { useMutation } from "@tanstack/react-query";
import { authApi } from "../../lib/api";

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
  });
};
