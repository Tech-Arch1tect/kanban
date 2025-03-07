import { useMutation, useQueryClient } from "@tanstack/react-query";
import { notificationsApi } from "../../lib/api";
import { toast } from "react-toastify";
import { NotificationUpdateNotificationRequest } from "../../typescript-fetch-client";

export const useUpdateNotificationConfiguration = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (notification: NotificationUpdateNotificationRequest) => {
      return await notificationsApi.apiV1NotificationsUpdatePost({
        request: notification,
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["notifications"] });
      toast.success("Notification configuration updated successfully!");
    },
    onError: (error: Error) => {
      toast.error(
        error.message || "Failed to update notification configuration."
      );
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
