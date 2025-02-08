import { useQuery } from "@tanstack/react-query";

import { notificationsApi } from "../../lib/api";

export const useEventData = () => {
  const { data, isLoading, error } = useQuery({
    queryKey: ["events"],
    queryFn: () => notificationsApi.apiV1NotificationsEventsGet(),
  });
  return { data, isLoading, error };
};
