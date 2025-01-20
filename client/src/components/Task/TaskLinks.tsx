import { ModelsTask } from "../../typescript-fetch-client";
import { QueryAndAddTaskLink } from "./QueryAndAddTaskLink";
import { TaskLinkLoop } from "./TaskLinkLoop";

export const TaskLinks = ({ task }: { task: ModelsTask }) => {
  return (
    <div className="space-y-4">
      <h2 className="text-2xl font-bold text-gray-700 dark:text-gray-200 mb-4">
        Task Links
      </h2>
      <QueryAndAddTaskLink task={task} />
      <ul>
        <TaskLinkLoop links={task.srcLinks ?? []} />
        <TaskLinkLoop links={task.dstLinks ?? []} />
      </ul>
    </div>
  );
};
