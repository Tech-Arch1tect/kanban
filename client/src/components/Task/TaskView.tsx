import { useParams } from "@tanstack/react-router";
import { useTaskData } from "../../hooks/tasks/useTaskData";
import TaskComments from "./TaskComments";
import { TaskHeading } from "./Heading/TaskHeading";
import { ModelsTask } from "../../typescript-fetch-client";
import { TaskDescription } from "./TaskDescription";
import { ViewFiles } from "./Files/ViewFiles";
import { TaskLinks } from "./TaskLinks";
import { TaskExternalLinks } from "./TaskExternalLinks";
import { TaskSubTasks } from "./TaskSubTasks";
import TaskActivities from "./TaskActivities";

export default function TaskView() {
  const { id } = useParams({ from: "/task/$id" });
  const { data, isLoading, error } = useTaskData({ id: Number(id) });

  if (isLoading)
    return (
      <div className="flex justify-center items-center text-gray-500 dark:text-gray-400">
        Loading...
      </div>
    );
  if (error)
    return (
      <div className="text-red-500 dark:text-red-400 font-medium">
        Error: {error.message}
      </div>
    );

  return (
    <div className="p-4 bg-gray-50 dark:bg-gray-900 min-h-screen">
      <div className="max-w-screen-2xl mx-auto bg-white dark:bg-gray-800 rounded-lg shadow-md p-6 space-y-6">
        <TaskHeading task={data?.task as ModelsTask} />

        <div className="space-y-2">
          <TaskDescription task={data?.task as ModelsTask} />
        </div>

        <TaskSubTasks task={data?.task as ModelsTask} />

        <ViewFiles task={data?.task as ModelsTask} />

        <TaskLinks task={data?.task as ModelsTask} />

        <TaskExternalLinks task={data?.task as ModelsTask} />

        <TaskActivities task={data?.task as ModelsTask} />

        <TaskComments task={data?.task as ModelsTask} />
      </div>
    </div>
  );
}
