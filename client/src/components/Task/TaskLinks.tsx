import { ModelsTask } from "../../typescript-fetch-client";
import { QueryAndAddTaskLink } from "./QueryAndAddTaskLink";
import { TaskLinkLoop } from "./TaskLinkLoop";

export const TaskLinks = ({ task }: { task: ModelsTask }) => {
  return (
    <div>
      <h2 className="text-2xl font-bold mb-4">Task Links</h2>
      <QueryAndAddTaskLink task={task} />
      <ul>
        <TaskLinkLoop links={task.srcLinks ?? []} />
        <TaskLinkLoop links={task.dstLinks ?? []} />
      </ul>
    </div>
  );
};
