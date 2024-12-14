import { useMutation } from '@tanstack/react-query';
import { authApi } from '../lib/api';

interface ChangePasswordArgs {
  currentPassword: string;
  newPassword: string;
}

export const useChangePassword = () => {
  return useMutation({
    mutationFn: async ({ currentPassword, newPassword }: ChangePasswordArgs) => {
      return await authApi.apiV1AuthChangePasswordPost({
        passwordChange: {
          currentPassword,
          newPassword,
        },
      });
    },
  });
};
