import { useMutation, useQueryClient } from '@tanstack/react-query';
import { authApi } from '../lib/api';

export const useAuth = (profile?: { role?: string }) => {
  const queryClient = useQueryClient();

  const logoutMutation = useMutation({
    mutationFn: () => authApi.apiV1AuthLogoutPost(),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['userProfile'] });
    },
    onError: (error) => {
      console.error("Error logging out:", error);
    },
  });

  const handleLogout = () => {
    logoutMutation.mutate();
  };

  const isAdmin = profile?.role === "admin";

  return { handleLogout, isAdmin };
};
