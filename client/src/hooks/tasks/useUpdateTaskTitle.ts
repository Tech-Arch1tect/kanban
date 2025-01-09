import { useMutation, useQueryClient } from "@tanstack/react-query";
import { tasksApi } from "../../lib/api";
import { toast } from "react-toastify";
import { TaskUpdateTaskTitleRequest } from "../../typescript-fetch-client";

export const useUpdateTaskTitle = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (task: TaskUpdateTaskTitleRequest) => {
      return await tasksApi
        .apiV1TasksUpdateTitlePost({
          request: {
            taskId: task.taskId,
            title: task.title,
          },
        })
        .then(() => {
          queryClient.invalidateQueries({ queryKey: ["task", task.taskId] });
        });
    },
    onSuccess: () => {
      toast.success("Task title updated successfully!");
    },
    onError: (error: any) => {
      toast.error(error.message || "Failed to update task title.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
