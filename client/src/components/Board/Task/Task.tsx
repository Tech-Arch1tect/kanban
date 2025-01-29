import { Link } from "@tanstack/react-router";
import { ChevronDownIcon } from "@heroicons/react/24/solid";
import { ModelsTask } from "../../../typescript-fetch-client";
import RenderMarkdown from "../../Utility/RenderMarkdown";
import { useState } from "react";

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
  const [isSubtasksOpen, setSubtasksOpen] = useState(false);
  const [showMore, setShowMore] = useState(false);

  if (!task.id) return null;
  const dragged = isDraggedTask(task.id);
  const hovered = isHoveredTask(task.id);
  const subtaskCount = task.subtasks?.length || 0;

  const visibleSubtasks = showMore ? task.subtasks : task.subtasks?.slice(0, 3); // Show up to 3 subtasks initially

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

      {/* Subtasks Section */}
      {subtaskCount > 0 && (
        <div className="mt-4">
          <button
            onClick={() => setSubtasksOpen(!isSubtasksOpen)}
            className="flex items-center justify-between w-full text-sm text-gray-600 dark:text-gray-300 hover:text-blue-500 dark:hover:text-blue-400"
            aria-expanded={isSubtasksOpen ? "true" : "false"}
          >
            <span className="flex items-center space-x-2">
              <span>{subtaskCount} Subtasks</span>
              <ChevronDownIcon
                className={`w-5 h-5 transform transition-transform ${
                  isSubtasksOpen ? "rotate-180" : ""
                }`}
              />
            </span>
          </button>

          {isSubtasksOpen && (
            <div className="mt-2 space-y-2 bg-gray-50 dark:bg-gray-800 p-3 rounded-lg shadow-sm">
              {visibleSubtasks?.map((subtask) => (
                <Link
                  key={subtask.id}
                  //@ts-ignore
                  to={`/task/${subtask.id}`}
                  className={`block text-sm text-gray-700 dark:text-gray-300 hover:text-blue-500 dark:hover:text-blue-400 truncate
                    ${subtask.status === "open" ? "" : "line-through text-gray-500 dark:text-gray-400"}
                  `}
                >
                  {subtask.title} -{" "}
                  {(subtask.assignee?.username &&
                    `@${subtask.assignee.username}`) ||
                    "Unassigned"}
                </Link>
              ))}

              {/* Show more/less button */}
              {(task.subtasks?.length as number) > 3 && !showMore && (
                <button
                  onClick={() => setShowMore(true)}
                  className="w-full text-sm text-gray-600 dark:text-gray-300 hover:text-blue-500 dark:hover:text-blue-400 mt-2"
                >
                  Show More Subtasks
                </button>
              )}

              {showMore && (
                <button
                  onClick={() => setShowMore(false)}
                  className="w-full text-sm text-gray-600 dark:text-gray-300 hover:text-blue-500 dark:hover:text-blue-400 mt-2"
                >
                  Show Less Subtasks
                </button>
              )}
            </div>
          )}
        </div>
      )}
    </div>
  );
}
