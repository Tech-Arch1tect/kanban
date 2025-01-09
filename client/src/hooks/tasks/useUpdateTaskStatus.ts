import { useMutation, useQueryClient } from "@tanstack/react-query";
import { tasksApi } from "../../lib/api";
import { toast } from "react-toastify";
import { TaskUpdateTaskStatusRequest } from "../../typescript-fetch-client";

export const useUpdateTaskStatus = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (task: TaskUpdateTaskStatusRequest) => {
      return await tasksApi
        .apiV1TasksUpdateStatusPost({
          request: {
            taskId: task.taskId,
            status: task.status,
          },
        })
        .then(() => {
          queryClient.invalidateQueries({ queryKey: ["task", task.taskId] });
        });
    },
    onSuccess: () => {
      toast.success("Task status updated successfully!");
    },
    onError: (error: any) => {
      toast.error(error.message || "Failed to update task status.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
