import { useMutation } from "@tanstack/react-query";
import { authApi } from "../../lib/api";
import { toast } from "react-toastify";

interface ChangePasswordArgs {
  currentPassword: string;
  newPassword: string;
}

export const useChangePassword = () => {
  return useMutation({
    mutationFn: async ({
      currentPassword,
      newPassword,
    }: ChangePasswordArgs) => {
      return await authApi.apiV1AuthChangePasswordPost({
        passwordChange: {
          currentPassword,
          newPassword,
        },
      });
    },
    onSuccess: () => {
      toast.success("Password changed successfully!");
    },
    onError: (error: any) => {
      toast.error(error.message || "Failed to change password.");
    },
  });
};
