import { useGetUsersWithAccessToBoard } from "./useGetUsersWithAccessToBoard";

export const useBoardUsernames = (boardId: number) => {
  const { data: usersData, isLoading } = useGetUsersWithAccessToBoard({
    id: boardId,
  });
  const usernames = usersData?.users?.map((user) => user.username) || [];
  return { usernames, isLoading };
};
