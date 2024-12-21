import { useParams } from "@tanstack/react-router";
import { useTaskData } from "../../hooks/tasks/useTaskData";
import RenderMarkdown from "../Utility/RenderMarkdown";

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
        <div className="border-b pb-4">
          <h1 className="text-2xl font-bold text-gray-900">
            {data?.task?.title}
          </h1>
          <div className="grid grid-cols-2 sm:grid-cols-3 gap-4 text-sm mt-4">
            <div className="flex flex-col">
              <span className="font-medium text-gray-600">Assignee</span>
              <span className="text-gray-800">
                {data?.task?.assignee?.displayName || "Unassigned"}
              </span>
            </div>
            <div className="flex flex-col">
              <span className="font-medium text-gray-600">Status</span>
              <span className="text-gray-800">{data?.task?.status}</span>
            </div>
            <div className="flex flex-col">
              <span className="font-medium text-gray-600">Creator</span>
              <span className="text-gray-800">
                {data?.task?.creator?.displayName}
              </span>
            </div>
            <div className="flex flex-col">
              <span className="font-medium text-gray-600">Created At</span>
              <span className="text-gray-800">{data?.task?.createdAt}</span>
            </div>
            <div className="flex flex-col">
              <span className="font-medium text-gray-600">Updated At</span>
              <span className="text-gray-800">{data?.task?.updatedAt}</span>
            </div>
            <div className="flex flex-col">
              <span className="font-medium text-gray-600">Swimlane</span>
              <span className="text-gray-800">
                {data?.task?.swimlane?.name || "None"}
              </span>
            </div>
            <div className="flex flex-col">
              <span className="font-medium text-gray-600">Column</span>
              <span className="text-gray-800">
                {data?.task?.column?.name || "None"}
              </span>
            </div>
          </div>
        </div>

        <div className="space-y-2">
          <h2 className="text-lg font-semibold text-gray-900">Description</h2>
          <p className="text-gray-700">
            <RenderMarkdown
              markdown={data?.task?.description || "No description provided."}
            />
          </p>
        </div>

        <div>
          <h2 className="text-lg font-semibold text-gray-900">Comments</h2>
          <div className="space-y-4 mt-4">
            {data?.task?.comments?.length &&
            data?.task?.comments?.length > 0 ? (
              data.task.comments.map((comment) => (
                <div
                  key={comment.id}
                  className="p-4 border border-gray-200 rounded-md bg-gray-50"
                >
                  <p className="text-gray-800">
                    <RenderMarkdown markdown={comment.text || ""} />
                  </p>
                  <div className="mt-2 text-sm text-gray-600 flex justify-between">
                    <span>{comment.user?.displayName || "Unknown User"}</span>
                    <span>{comment.createdAt}</span>
                  </div>
                </div>
              ))
            ) : (
              <div className="text-gray-500">No comments available.</div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
