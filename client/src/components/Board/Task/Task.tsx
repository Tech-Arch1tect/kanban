import { Link } from "@tanstack/react-router";
import { ChevronDownIcon } from "@heroicons/react/24/solid";
import { ModelsTask } from "../../../typescript-fetch-client";
import RenderMarkdown from "../../Utility/RenderMarkdown";
import { useState } from "react";
import {
  colourToClass,
  colourToBorderClass,
} from "../../Utility/colourToClass";

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

  const visibleSubtasks = showMore ? task.subtasks : task.subtasks?.slice(0, 3);

  return (
    <div
      draggable
      data-position={task.position}
      data-task-id={task.id}
      onDragStart={(event) => handleDragStart(event, task.id ?? 0)}
      className={`task rounded-lg shadow-sm dark:shadow-md cursor-move transition-all duration-200 border-2 ${colourToBorderClass[task.colour as keyof typeof colourToBorderClass]}
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
        {/* Task Header & Title */}
        <div
          className={`${colourToClass[task.colour as keyof typeof colourToClass] || "dark:text-white"}  p-4 rounded-t-lg`}
        >
          <div className="flex items-center justify-between text-base">
            <span className="truncate">
              {task.assignee?.username || "Unassigned"}
            </span>
            <span className="">{`#${task.id}`}</span>
          </div>
          <h3 className="text-lg font-semibold truncate">{task.title}</h3>
        </div>

        {/* Task Description */}
        <div className="p-4">
          <div className="text-base leading-relaxed line-clamp-3">
            <RenderMarkdown markdown={task.description || ""} />
          </div>
        </div>
      </Link>

      {/* Subtasks Section */}
      {subtaskCount > 0 && (
        <div className="mt-4">
          <button
            onClick={() => setSubtasksOpen(!isSubtasksOpen)}
            className="flex items-center justify-between w-full text-base hover:text-blue-500 px-2 py-1"
            aria-expanded={isSubtasksOpen}
          >
            <span className="flex items-center space-x-2 dark:text-white">
              <span>{subtaskCount} Subtasks</span>
              <ChevronDownIcon
                className={`w-5 h-5 transition-transform ${
                  isSubtasksOpen ? "rotate-180" : ""
                }`}
              />
            </span>
          </button>

          <div
            className={`transition-all duration-300 overflow-hidden ${
              isSubtasksOpen ? "max-h-[500px] opacity-100" : "max-h-0 opacity-0"
            }`}
          >
            <div className="mt-2 space-y-2 bg-gray-50 dark:bg-gray-800 p-3 rounded-lg shadow-sm dark:text-white">
              {visibleSubtasks?.map((subtask) => (
                <Link
                  key={subtask.id}
                  //@ts-ignore
                  to={`/task/${subtask.id}`}
                  className={`block text-base leading-relaxed hover:text-blue-500 truncate
                    ${subtask.status === "open" ? "" : "line-through"}
                  `}
                >
                  {subtask.title} -{" "}
                  {(subtask.assignee?.username &&
                    `@${subtask.assignee.username}`) ||
                    "Unassigned"}
                </Link>
              ))}

              {(task.subtasks?.length as number) > 3 && (
                <button
                  onClick={() => setShowMore((prev) => !prev)}
                  className="w-full text-base hover:text-blue-500 mt-2 text-black dark:text-white text-left"
                >
                  {showMore ? "Show Less Subtasks" : "Show More Subtasks"}
                </button>
              )}
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
