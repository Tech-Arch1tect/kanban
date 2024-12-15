import { ModelsTask } from "../../typescript-fetch-client";

export function Task({ task }: { task: ModelsTask }) {
  return (
    <div className="bg-white p-4 rounded shadow-md">
      <h3 className="text-lg font-semibold">{task.title}</h3>
      <p className="text-sm text-gray-600 truncate">{task.description}</p>
    </div>
  );
}
