import { useMutation, useQueryClient } from "@tanstack/react-query";
import { tasksApi } from "../../lib/api";
import { toast } from "react-toastify";
import {
  TaskUpdateTaskColourRequest,
  TaskUpdateTaskColourResponse,
} from "../../typescript-fetch-client";

export const useUpdateTaskColour = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (task: TaskUpdateTaskColourRequest) => {
      return await tasksApi.apiV1TasksUpdateColourPost({
        request: {
          taskId: task.taskId,
          colour: task.colour,
        },
      });
    },
    onSuccess: (response: TaskUpdateTaskColourResponse) => {
      queryClient.invalidateQueries({ queryKey: ["task", response.task?.id] });
      if (response.task?.parentTaskId) {
        queryClient.invalidateQueries({
          queryKey: ["task", response.task?.parentTaskId],
        });
      }
      toast.success("Task colour updated successfully!");
    },
    onError: (error: Error) => {
      toast.error(error.message || "Failed to update task colour.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
