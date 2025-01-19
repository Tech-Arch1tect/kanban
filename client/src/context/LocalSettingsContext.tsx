import { createContext, useState, useEffect, ReactNode } from "react";

type LocalSettings = {
  theme: string;
  collapsedSwimlanes: Record<string, boolean>;
};

type LocalSettingsContextValue = {
  localSettings: LocalSettings;
  updateLocalSettings: (updates: Partial<LocalSettings>) => void;
};

const defaultLocalSettings: LocalSettings = {
  theme: "light",
  collapsedSwimlanes: {},
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
      if (stored) {
        const parsed = JSON.parse(stored);
        return {
          ...defaultLocalSettings,
          ...parsed,
          collapsedSwimlanes: {
            ...defaultLocalSettings.collapsedSwimlanes,
            ...parsed.collapsedSwimlanes,
          },
        };
      } else {
        return defaultLocalSettings;
      }
    } catch {
      return defaultLocalSettings;
    }
  });

  useEffect(() => {
    localStorage.setItem("localSettings", JSON.stringify(localSettings));
  }, [localSettings]);

  const updateLocalSettings = (updates: Partial<LocalSettings>) => {
    setLocalSettings((prev) => ({
      ...prev,
      ...updates,
      collapsedSwimlanes: updates.collapsedSwimlanes
        ? {
            ...prev.collapsedSwimlanes,
            ...updates.collapsedSwimlanes,
          }
        : prev.collapsedSwimlanes,
    }));
  };

  return (
    <LocalSettingsContext.Provider
      value={{ localSettings, updateLocalSettings }}
    >
      {children}
    </LocalSettingsContext.Provider>
  );
};
