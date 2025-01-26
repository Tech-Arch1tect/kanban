import { useMutation, useQueryClient } from "@tanstack/react-query";
import { tasksApi } from "../../lib/api";
import { toast } from "react-toastify";
import { TaskEditTaskRequest } from "../../typescript-fetch-client";

export const useEditTask = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (task: TaskEditTaskRequest) => {
      return await tasksApi
        .apiV1TasksEditPost({
          request: {
            id: task.id,
            title: task.title,
            description: task.description,
            status: task.status,
            assigneeId: task.assigneeId,
          },
        })
        .then((response) => {
          queryClient.invalidateQueries({ queryKey: ["tasks"] });
          queryClient.invalidateQueries({ queryKey: ["task", task.id] });
          if (response.task?.parentTaskId) {
            queryClient.invalidateQueries({
              queryKey: ["task", response.task.parentTaskId],
            });
          }
        });
    },
    onSuccess: () => {
      toast.success("Task updated successfully!");
    },
    onError: (error: any) => {
      toast.error(error.message || "Failed to update task.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
