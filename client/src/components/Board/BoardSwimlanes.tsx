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
      <div
        className="font-semibold text-lg text-gray-700 text-center bg-white py-3 rounded-t-lg shadow-sm cursor-pointer"
        onClick={handleToggle}
      >
        {swimlane.name}{" "}
        <span className="ml-2 text-sm">
          ({isCollapsed ? "Expand" : "Collapse"})
        </span>
      </div>

      {!isCollapsed && (
        <div
          className="grid gap-2 bg-gray-50 p-2 rounded-b-lg shadow-sm"
          style={{
            gridTemplateColumns: `repeat(${columns?.length ?? 0}, minmax(200px, 1fr))`,
          }}
        >
          {columns?.map((column) => (
            <ColumnTasks
              key={column.id}
              column={column}
              swimlane={swimlane}
              tasks={tasks}
            />
          ))}
        </div>
      )}
    </div>
  );
}
