import { ModelsTask } from "../../../typescript-fetch-client";
import { TaskAssignee } from "./TaskAssignee";
import { TaskColour } from "./TaskColour";
import { TaskDueDate } from "./TaskDueDate";
import { TaskStatus } from "./TaskStatus";
import { TaskTitle } from "./TaskTitle";

import { colourToBorderClass } from "../../Utility/colourToClass";

export function TaskHeading({ task }: { task: ModelsTask }) {
  return (
    <div
      className={`border border-2 p-4 ${colourToBorderClass[task.colour as keyof typeof colourToBorderClass]} rounded-lg`}
    >
      <TaskTitle task={task} />
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-2 text-sm">
        <TaskAssignee task={task} />
        <TaskStatus task={task} />
        <div className="flex items-center">
          <span className="font-medium text-gray-600 dark:text-gray-400 w-24">
            Creator:
          </span>
          <span className="text-gray-800 dark:text-gray-200">
            {task?.creator?.username || "Unknown User"}
          </span>
        </div>
        <div className="flex items-center">
          <span className="font-medium text-gray-600 dark:text-gray-400 w-24">
            Created At:
          </span>
          <span className="text-gray-800 dark:text-gray-200">
            {task?.createdAt
              ? new Date(task?.createdAt).toLocaleString()
              : "N/A"}
          </span>
        </div>
        <div className="flex items-center">
          <span className="font-medium text-gray-600 dark:text-gray-400 w-24">
            Updated At:
          </span>
          <span className="text-gray-800 dark:text-gray-200">
            {task?.updatedAt
              ? new Date(task?.updatedAt).toLocaleString()
              : "N/A"}
          </span>
        </div>
        <div className="flex items-center">
          <span className="font-medium text-gray-600 dark:text-gray-400 w-24">
            Swimlane:
          </span>
          <span className="text-gray-800 dark:text-gray-200">
            {task?.swimlane?.name || "None"}
          </span>
        </div>
        <div className="flex items-center">
          <span className="font-medium text-gray-600 dark:text-gray-400 w-24">
            Column:
          </span>
          <span className="text-gray-800 dark:text-gray-200">
            {task?.column?.name || "None"}
          </span>
        </div>
        <TaskDueDate task={task} />
        <TaskColour task={task} />
      </div>
    </div>
  );
}
