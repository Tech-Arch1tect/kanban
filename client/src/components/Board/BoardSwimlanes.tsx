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
  return (
    <div className="mb-2">
      <div className="font-semibold text-lg text-gray-700 text-center bg-white py-3 rounded-t-lg shadow-sm">
        {swimlane.name}
      </div>
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
    </div>
  );
}
