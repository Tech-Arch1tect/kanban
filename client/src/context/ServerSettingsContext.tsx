import React, { createContext, useContext, ReactNode, useMemo } from "react";
import { useGetSettings } from "../hooks/settings/useGetSettings";
import { useUpdateSettings } from "../hooks/settings/useUpdateSettings";
import { ModelsSettings } from "../typescript-fetch-client/models/ModelsSettings";

interface ServerSettingsContextValue {
  settings?: ModelsSettings;
  isLoading: boolean;
  error: unknown;
  updateSettings: (updatedSettings: ModelsSettings) => void;
  isUpdating: boolean;
}

const ServerSettingsContext = createContext<ServerSettingsContextValue>(
  {} as ServerSettingsContextValue
);

export const ServerSettingsProvider = ({
  children,
}: {
  children: ReactNode;
}) => {
  const { data: response, isLoading, error } = useGetSettings();
  const { mutate: updateSettingsMutation, isPending: isUpdating } =
    useUpdateSettings();

  const updateSettings = (updatedSettings: ModelsSettings) => {
    updateSettingsMutation(updatedSettings);
  };

  const value = useMemo(
    () => ({
      settings: response?.settings,
      isLoading,
      error,
      updateSettings,
      isUpdating,
    }),
    [response?.settings, isLoading, error, isUpdating]
  );

  return (
    <ServerSettingsContext.Provider value={value}>
      {children}
    </ServerSettingsContext.Provider>
  );
};

export const useServerSettings = () => useContext(ServerSettingsContext);
