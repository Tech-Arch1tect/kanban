import { useQuery } from "@tanstack/react-query";

import { notificationsApi } from "../../lib/api";

export const useNotificationConfigurationData = () => {
  const { data, isLoading, error } = useQuery({
    queryKey: ["notifications"],
    queryFn: () => notificationsApi.apiV1NotificationsListGet(),
  });
  return { data, isLoading, error };
};
