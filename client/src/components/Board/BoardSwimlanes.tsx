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
    <div className="mb-6">
      <div className="font-semibold text-center bg-gray-200 py-2 rounded">
        {swimlane.name}
      </div>
      <div
        className="grid gap-4 mt-2"
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
