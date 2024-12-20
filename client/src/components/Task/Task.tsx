import { ModelsTask } from "../../typescript-fetch-client";

export function Task({ task }: { task: ModelsTask }) {
  return (
    <div
      draggable
      onDragStart={(event: React.DragEvent<HTMLDivElement>) => {
        const data = JSON.stringify({
          taskId: task.id,
          position: task.position,
        });
        event.dataTransfer.setData("text/plain", data);
      }}
      className="bg-white p-4 rounded shadow-md cursor-move task"
    >
      <h3 className="text-lg font-semibold">{task.title}</h3>
      <p className="text-sm text-gray-600 truncate">{task.description}</p>
    </div>
  );
}
