import { ModelsTask, ModelsTaskLinks } from "../../typescript-fetch-client";
import { QueryAndAddTaskLink } from "./QueryAndAddTaskLink";
import { TaskLinkLoop } from "./TaskLinkLoop";

export const TaskLinks = ({ task }: { task: ModelsTask }) => {
  const hasSrcLinks = task.srcLinks && task.srcLinks.length > 0;
  const hasDstLinks = task.dstLinks && task.dstLinks.length > 0;

  return (
    <div className="space-y-4">
      <h2 className="text-2xl font-bold text-gray-700 dark:text-gray-200 mb-4">
        Task Links
      </h2>
      <QueryAndAddTaskLink task={task} />
      {hasSrcLinks || hasDstLinks ? (
        <ul>
          {hasSrcLinks && (
            <TaskLinkLoop links={task.srcLinks as ModelsTaskLinks[]} />
          )}
          {hasDstLinks && (
            <TaskLinkLoop links={task.dstLinks as ModelsTaskLinks[]} />
          )}
        </ul>
      ) : (
        <div className="text-gray-500 dark:text-gray-400 text-center">
          No task links available.
        </div>
      )}
    </div>
  );
};
