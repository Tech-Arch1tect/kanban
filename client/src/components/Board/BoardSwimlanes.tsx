import { useContext } from "react";
import { LocalSettingsContext } from "../../context/LocalSettingsContext";
import ColumnTasks from "./ColumnTasks";
import {
  ModelsColumn,
  ModelsSwimlane,
  ModelsTask,
} from "../../typescript-fetch-client";

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
      <button
        className="flex items-center justify-between w-full px-4 py-3 font-semibold text-lg text-gray-700 dark:text-gray-200 bg-white dark:bg-gray-800 rounded-t-lg shadow-sm dark:shadow-md transition-all hover:bg-gray-100 dark:hover:bg-gray-700"
        onClick={handleToggle}
      >
        <span>{swimlane.name}</span>
      </button>

      {!isCollapsed && (
        <div className="p-2 bg-gray-50 dark:bg-gray-700 rounded-b-lg shadow-sm dark:shadow-md">
          <div
            className="grid gap-4"
            style={{
              gridTemplateColumns: `repeat(auto-fit, minmax(200px, 1fr))`,
            }}
          >
            {columns.map((column) => (
              <div key={column.id} className="flex flex-col">
                <div className="text-center font-semibold text-gray-700 dark:text-gray-200 bg-white dark:bg-gray-800 p-2 rounded-lg shadow-md mb-1">
                  {column.name}
                </div>

                <ColumnTasks
                  column={column}
                  swimlane={swimlane}
                  tasks={tasks}
                />
              </div>
            ))}
          </div>
        </div>
      )}
    </div>
  );
}
