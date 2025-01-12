import { ModelsTask } from "../../../typescript-fetch-client";
import { TaskAssignee } from "./TaskAssignee";
import { TaskStatus } from "./TaskStatus";
import { TaskTitle } from "./TaskTitle";

export function TaskHeading({ task }: { task: ModelsTask }) {
  return (
    <div className="border-b pb-4">
      <TaskTitle task={task} />
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-2 text-sm">
        <TaskAssignee task={task} />
        <TaskStatus task={task} />
        <div className="flex items-center">
          <span className="font-medium text-gray-600 w-24">Creator:</span>
          <span className="text-gray-800">
            {task?.creator?.username || "Unknown User"}
          </span>
        </div>
        <div className="flex items-center">
          <span className="font-medium text-gray-600 w-24">Created At:</span>
          <span className="text-gray-800">
            {task?.createdAt
              ? new Date(task?.createdAt).toLocaleString()
              : "N/A"}
          </span>
        </div>
        <div className="flex items-center">
          <span className="font-medium text-gray-600 w-24">Updated At:</span>
          <span className="text-gray-800">
            {task?.updatedAt
              ? new Date(task?.updatedAt).toLocaleString()
              : "N/A"}
          </span>
        </div>
        <div className="flex items-center">
          <span className="font-medium text-gray-600 w-24">Swimlane:</span>
          <span className="text-gray-800">
            {task?.swimlane?.name || "None"}
          </span>
        </div>
        <div className="flex items-center">
          <span className="font-medium text-gray-600 w-24">Column:</span>
          <span className="text-gray-800">{task?.column?.name || "None"}</span>
        </div>
      </div>
    </div>
  );
}
