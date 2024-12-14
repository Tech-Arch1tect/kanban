import { useMutation } from '@tanstack/react-query';
import { authApi } from '../../lib/api';

export const useSendResetCode = () => {
  return useMutation({
    mutationFn: async (email: string) => {
      return await authApi.apiV1AuthPasswordResetPost({
        passwordReset: { email },
      });
    },
  });
};
