import { ModelsTask } from "../../../typescript-fetch-client";

interface TaskProps {
  task: ModelsTask;
  handleDragStart: (e: React.DragEvent<HTMLDivElement>, taskId: number) => void;
  isDraggedTask: (taskId: number) => boolean;
  isHoveredTask: (taskId: number) => boolean;
}

export function Task({
  task,
  handleDragStart,
  isDraggedTask,
  isHoveredTask,
}: TaskProps) {
  if (!task.id) return null;
  const dragged = isDraggedTask(task.id);
  const hovered = isHoveredTask(task.id);

  return (
    <div
      draggable
      data-position={task.position}
      data-task-id={task.id}
      onDragStart={(event) => handleDragStart(event, task.id ?? 0)}
      className={`bg-white p-3 rounded-md shadow-sm cursor-move task 
        ${dragged ? "opacity-60" : ""}
        ${hovered ? "ring ring-blue-400" : ""}
        flex flex-col space-y-2 transition-all duration-200
      `}
    >
      <div className="flex items-center justify-between text-sm text-gray-600">
        <span className="truncate">
          {task.assignee?.displayName || "Unassigned"}
        </span>
        <span className="text-gray-400">{`#${task.id}`}</span>
      </div>
      <h3 className="text-base font-medium text-gray-900 truncate">
        {task.title}
      </h3>
      <p className="text-sm text-gray-700 line-clamp-2">{task.description}</p>
    </div>
  );
}
