import { useState } from "react";
import { ModelsTask, ModelsTaskLinks } from "../../typescript-fetch-client";
import { QueryAndAddTaskLink } from "./QueryAndAddTaskLink";
import { TaskLinkLoop } from "./TaskLinkLoop";
import { PlusIcon, XMarkIcon } from "@heroicons/react/24/outline";

export const TaskLinks = ({ task }: { task: ModelsTask }) => {
  const [showAddLink, setShowAddLink] = useState(false);
  const srcCount = task.srcLinks ? task.srcLinks.length : 0;
  const dstCount = task.dstLinks ? task.dstLinks.length : 0;
  const totalLinks = srcCount + dstCount;

  return (
    <div className="space-y-4">
      <div className="flex items-center justify-between">
        <h2 className="text-2xl font-bold text-gray-700 dark:text-gray-200 flex items-center space-x-2">
          <span>Task Links</span>
          <span className="bg-gray-200 text-gray-700 dark:bg-gray-700 dark:text-gray-200 px-2 py-1 rounded text-sm">
            {totalLinks}
          </span>
        </h2>
        <button
          onClick={() => setShowAddLink(!showAddLink)}
          className="flex-shrink-0 text-blue-500 hover:text-blue-700"
          title={showAddLink ? "Close" : "Add link"}
        >
          {showAddLink ? (
            <XMarkIcon className="w-6 h-6" />
          ) : (
            <PlusIcon className="w-6 h-6" />
          )}
        </button>
      </div>

      {showAddLink && <QueryAndAddTaskLink task={task} />}

      {totalLinks > 0 && (
        <ul className="space-y-2">
          {task.srcLinks && task.srcLinks.length > 0 && (
            <TaskLinkLoop links={task.srcLinks as ModelsTaskLinks[]} />
          )}
          {task.dstLinks && task.dstLinks.length > 0 && (
            <TaskLinkLoop links={task.dstLinks as ModelsTaskLinks[]} />
          )}
        </ul>
      )}
    </div>
  );
};
