import { createContext, useState, useEffect, ReactNode } from "react";

type LocalSettings = {
  theme: string;
};

type LocalSettingsContextValue = {
  localSettings: LocalSettings;
  updateLocalSettings: (updates: Partial<LocalSettings>) => void;
};

const defaultLocalSettings: LocalSettings = {
  theme: "light",
};

export const LocalSettingsContext = createContext<LocalSettingsContextValue>({
  localSettings: defaultLocalSettings,
  updateLocalSettings: () => {},
});

type Props = {
  children: ReactNode;
};

export const LocalSettingsProvider = ({ children }: Props) => {
  const [localSettings, setLocalSettings] = useState<LocalSettings>(() => {
    try {
      const stored = localStorage.getItem("localSettings");
      return stored ? JSON.parse(stored) : defaultLocalSettings;
    } catch {
      return defaultLocalSettings;
    }
  });

  useEffect(() => {
    localStorage.setItem("localSettings", JSON.stringify(localSettings));
  }, [localSettings]);

  const updateLocalSettings = (updates: Partial<LocalSettings>) => {
    setLocalSettings((prev) => ({ ...prev, ...updates }));
  };

  return (
    <LocalSettingsContext.Provider
      value={{ localSettings, updateLocalSettings }}
    >
      {children}
    </LocalSettingsContext.Provider>
  );
};
