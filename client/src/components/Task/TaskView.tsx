import { useParams } from "@tanstack/react-router";
import { useTaskData } from "../../hooks/tasks/useTaskData";
import RenderMarkdown from "../Utility/RenderMarkdown";
import TaskComments from "./TaskComments";
import { TaskHeading } from "./Heading/TaskHeading";
import { ModelsTask } from "../../typescript-fetch-client";

export default function TaskView() {
  const { id } = useParams({ from: "/task/$id" });
  const { data, isLoading, error } = useTaskData({ id: Number(id) });

  if (isLoading)
    return (
      <div className="flex justify-center items-center text-gray-500">
        Loading...
      </div>
    );
  if (error)
    return (
      <div className="text-red-500 font-medium">Error: {error.message}</div>
    );

  return (
    <div className="p-4 bg-gray-50 min-h-screen">
      <div className="max-w-screen-2xl mx-auto bg-white rounded-lg shadow-md p-6 space-y-6">
        <TaskHeading task={data?.task as ModelsTask} />

        <div className="space-y-2">
          <h2 className="text-lg font-semibold text-gray-900">Description</h2>
          <div className="text-gray-700">
            <RenderMarkdown
              markdown={data?.task?.description || "No description provided."}
            />
          </div>
        </div>

        <TaskComments taskId={Number(id)} />
      </div>
    </div>
  );
}
