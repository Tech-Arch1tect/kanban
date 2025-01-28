import { useMutation, useQueryClient } from "@tanstack/react-query";
import { tasksApi } from "../../lib/api";
import { toast } from "react-toastify";
import { TaskUpdateTaskAssigneeRequest } from "../../typescript-fetch-client";

export const useUpdateTaskAssignee = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (task: TaskUpdateTaskAssigneeRequest) => {
      return await tasksApi.apiV1TasksUpdateAssigneePost({
        request: {
          taskId: task.taskId,
          assigneeId: task.assigneeId,
        },
      });
    },
    onSuccess: (response) => {
      queryClient.invalidateQueries({ queryKey: ["task", response.task?.id] });
      if (response.task?.parentTaskId) {
        queryClient.invalidateQueries({
          queryKey: ["task", response.task?.parentTaskId],
        });
      }
      toast.success("Task assignee updated successfully!");
    },
    onError: (error: any) => {
      toast.error(error.message || "Failed to update task assignee.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
