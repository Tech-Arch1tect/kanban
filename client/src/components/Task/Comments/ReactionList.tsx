import { ModelsReaction } from "../../../typescript-fetch-client";
import { useDeleteCommentReaction } from "../../../hooks/tasks/Comments/useDeleteCommentReaction";
import { useUserProfile } from "../../../hooks/useUserProfile";
import { TrashIcon } from "@heroicons/react/24/outline";

interface ReactionListProps {
  reactions: ModelsReaction[];
}

interface ReactionData {
  id: number;
  username: string;
  userId: number;
}

interface GroupedReaction {
  count: number;
  reactions: ReactionData[];
}

const ReactionList: React.FC<ReactionListProps> = ({ reactions }) => {
  const { profile, isLoading } = useUserProfile();
  const { mutate: deleteReaction } = useDeleteCommentReaction();

  if (!reactions || reactions.length === 0) return null;
  if (isLoading) return null;

  const groupedReactions = reactions.reduce<Record<string, GroupedReaction>>(
    (acc, reaction) => {
      const emoji = reaction.reaction;
      if (!emoji) return acc;
      if (!acc[emoji]) {
        acc[emoji] = { count: 0, reactions: [] };
      }
      acc[emoji].count += 1;
      if (reaction.user) {
        acc[emoji].reactions.push({
          id: reaction.id as number,
          username: reaction.user.username,
          userId: reaction.user.id as number,
        });
      }
      return acc;
    },
    {}
  );

  const handleDelete = (reactionId: number) => {
    deleteReaction(reactionId);
  };

  return (
    <div className="flex space-x-2 mt-2">
      {Object.entries(groupedReactions).map(([emoji, { count, reactions }]) => {
        const userReaction = reactions.find((r) => r.userId === profile?.id);
        return (
          <div
            key={emoji}
            className="flex items-center space-x-1 bg-gray-100 dark:bg-gray-800 px-2 py-1 rounded"
            title={reactions.map((r) => r.username).join(", ")}
          >
            <span className="text-xl">{emoji}</span>
            {count > 1 && (
              <span className="text-xs text-gray-700 dark:text-gray-300">
                {count}
              </span>
            )}
            {userReaction && (
              <button
                onClick={() => handleDelete(userReaction.id)}
                className="ml-1 text-red-500 hover:text-red-700"
                title="Remove your reaction"
              >
                <TrashIcon className="h-4 w-4" />
              </button>
            )}
          </div>
        );
      })}
    </div>
  );
};

export default ReactionList;
