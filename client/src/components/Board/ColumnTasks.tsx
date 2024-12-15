import { ModelsColumn, ModelsTask } from "../../typescript-fetch-client";

export default function ColumnTasks({ column }: { column: ModelsColumn }) {
  return (
    <div className="bg-gray-100 p-4 rounded shadow min-h-[100px]">
      <div className="space-y-2">{/* tasks should go here */}</div>
    </div>
  );
}
