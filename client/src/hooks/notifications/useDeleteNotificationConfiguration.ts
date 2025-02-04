import { useMutation } from "@tanstack/react-query";
import { notificationsApi } from "../../lib/api";
import { toast } from "react-toastify";

export const useDeleteNotificationConfiguration = () => {
  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (id: number) => {
      return await notificationsApi.apiV1NotificationsDeletePost({
        request: {
          id: id,
        },
      });
    },
    onSuccess: () => {
      toast.success("Notification configuration deleted successfully!");
    },
    onError: (error: any) => {
      toast.error(
        error.message || "Failed to delete notification configuration."
      );
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
