import { useContext } from "react";
import { LocalSettingsContext } from "../../context/LocalSettingsContext";
import ColumnTasks from "./ColumnTasks";
import {
  ModelsColumn,
  ModelsSwimlane,
  ModelsTask,
} from "../../typescript-fetch-client";
import { ChevronDownIcon, ChevronUpIcon } from "@heroicons/react/24/solid";

export default function BoardSwimlanes({
  swimlane,
  columns,
  tasks,
}: {
  swimlane: ModelsSwimlane;
  columns: ModelsColumn[];
  tasks: ModelsTask[];
}) {
  const { localSettings, updateLocalSettings } =
    useContext(LocalSettingsContext);

  const isCollapsed =
    localSettings.collapsedSwimlanes[swimlane.id ?? ""] ?? false;

  const handleToggle = () => {
    updateLocalSettings({
      collapsedSwimlanes: {
        ...localSettings.collapsedSwimlanes,
        [swimlane.id ?? ""]: !isCollapsed,
      },
    });
  };

  return (
    <div className="mb-2">
      {/* Swimlane Header */}
      <button
        className="flex items-center justify-between w-full px-4 py-3 font-semibold text-lg text-gray-700 dark:text-gray-200 bg-white dark:bg-gray-800 rounded-t-lg shadow-sm dark:shadow-md transition-all hover:bg-gray-100 dark:hover:bg-gray-700"
        onClick={handleToggle}
      >
        <span>{swimlane.name}</span>
        {isCollapsed ? (
          <ChevronDownIcon className="w-5 h-5 text-gray-500 dark:text-gray-400" />
        ) : (
          <ChevronUpIcon className="w-5 h-5 text-gray-500 dark:text-gray-400" />
        )}
      </button>

      {/* Swimlane Content */}
      <div
        className={`grid gap-2 bg-gray-50 dark:bg-gray-700 p-2 rounded-b-lg shadow-sm dark:shadow-md transition-all duration-300 ${
          isCollapsed ? "max-h-0 overflow-hidden opacity-0" : "opacity-100"
        }`}
        style={{
          gridTemplateColumns: `repeat(${columns?.length ?? 0}, minmax(200px, 1fr))`,
        }}
      >
        {!isCollapsed &&
          columns?.map((column) => (
            <ColumnTasks
              key={column.id}
              column={column}
              swimlane={swimlane}
              tasks={tasks}
            />
          ))}
      </div>
    </div>
  );
}
