import { useQuery } from "@tanstack/react-query";

import { settingsApi } from "../../lib/api";

export const useGetSettings = () => {
  const { data, isLoading, error } = useQuery({
    queryKey: ["settings"],
    queryFn: () => settingsApi.apiV1SettingsGetGet(),
  });
  return { data, isLoading, error };
};
