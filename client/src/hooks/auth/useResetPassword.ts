import { useMutation } from '@tanstack/react-query';
import { authApi } from '../../lib/api';

export const useResetPassword = () => {
  return useMutation({
    mutationFn: async ({
      email,
      code,
      password,
    }: {
      email: string;
      code: string;
      password: string;
    }) => {
      return await authApi.apiV1AuthResetPasswordPost({
        resetPassword: { email, code, password },
      });
    },
  });
};
