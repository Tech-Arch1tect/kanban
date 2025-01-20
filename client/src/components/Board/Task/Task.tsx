import { Link } from "@tanstack/react-router";
import { ModelsTask } from "../../../typescript-fetch-client";
import RenderMarkdown from "../../Utility/RenderMarkdown";

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
      className={`task bg-white dark:bg-gray-700 p-4 rounded-lg shadow-sm dark:shadow-md cursor-move transition-all duration-200
        ${dragged ? "opacity-50" : ""}
        ${hovered ? "ring-2 ring-blue-500 dark:ring-blue-400" : ""}
        hover:shadow-md dark:hover:shadow-lg hover:border-blue-100 dark:hover:border-blue-800
      `}
    >
      <Link
        //@ts-ignore
        to={`/task/${task.id}`}
        className="block"
      >
        {/* Assignee and Task ID */}
        <div className="flex items-center justify-between text-sm text-gray-600 dark:text-gray-300 mb-2">
          <span className="truncate">
            {task.assignee?.username || "Unassigned"}
          </span>
          <span className="text-gray-400 dark:text-gray-400">{`#${task.id}`}</span>
        </div>

        {/* Task Title */}
        <h3 className="text-base font-semibold text-gray-900 dark:text-gray-100 truncate mb-2">
          {task.title}
        </h3>

        {/* Task Description */}
        <div className="text-sm text-gray-700 dark:text-gray-300 line-clamp-2">
          <RenderMarkdown markdown={task.description || ""} />
        </div>
      </Link>
    </div>
  );
}
