import { useMutation, useQueryClient } from "@tanstack/react-query";
import { tasksApi } from "../../lib/api";
import { toast } from "react-toastify";
import { TaskUpdateTaskDescriptionRequest } from "../../typescript-fetch-client";

export const useUpdateTaskDescription = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (task: TaskUpdateTaskDescriptionRequest) => {
      return await tasksApi.apiV1TasksUpdateDescriptionPost({
        request: {
          taskId: task.taskId,
          description: task.description,
        },
      });
    },
    onSuccess: (response) => {
      queryClient.invalidateQueries({
        queryKey: ["task", response.task?.id],
      });
      toast.success("Task description updated successfully!");
    },
    onError: (error: any) => {
      toast.error(error.message || "Failed to update task description.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
