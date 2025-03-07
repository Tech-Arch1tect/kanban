import { useMutation, useQueryClient } from "@tanstack/react-query";
import { tasksApi } from "../../lib/api";
import { toast } from "react-toastify";
import { TaskUpdateTaskDueDateRequest } from "../../typescript-fetch-client";

export const useUpdateTaskDueDate = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (task: TaskUpdateTaskDueDateRequest) => {
      return await tasksApi.apiV1TasksUpdateDueDatePost({
        request: {
          taskId: task.taskId,
          dueDate: task.dueDate,
        },
      });
    },
    onSuccess: (response) => {
      queryClient.invalidateQueries({
        queryKey: ["task", response.task?.id],
      });
      toast.success("Task due date updated successfully!");
    },
    onError: (error: Error) => {
      toast.error(error.message || "Failed to update task due date.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
