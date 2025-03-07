import { useMutation, useQueryClient } from "@tanstack/react-query";
import { tasksApi } from "../../lib/api";
import { ServerApiControllersTaskMoveTaskRequest } from "../../typescript-fetch-client";
import { toast } from "react-toastify";
export const useMoveTask = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (task: ServerApiControllersTaskMoveTaskRequest) => {
      return await tasksApi
        .apiV1TasksMovePost({
          request: {
            taskId: task.taskId,
            columnId: task.columnId,
            swimlaneId: task.swimlaneId,
            position: task.position,
          },
        })
        .then(() => {
          queryClient.invalidateQueries({ queryKey: ["tasks"] });
        });
    },
    onSuccess: () => {
      toast.success("Task moved successfully!");
    },
    onError: (error: Error) => {
      toast.error(error.message || "Failed to move task.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
