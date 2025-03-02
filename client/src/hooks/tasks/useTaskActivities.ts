import { useQuery } from "@tanstack/react-query";

import { tasksApi } from "../../lib/api";

export const useTaskActivities = ({
  taskId,
  page,
  pageSize,
}: {
  taskId: number;
  page: number;
  pageSize: number;
}) => {
  const { data, isLoading, error } = useQuery({
    queryKey: ["task-activities", taskId, page, pageSize],
    queryFn: () =>
      tasksApi.apiV1TasksGetActivitiesPost({
        request: {
          taskId,
          page,
          pageSize,
        },
      }),
  });
  return { data, isLoading, error };
};
