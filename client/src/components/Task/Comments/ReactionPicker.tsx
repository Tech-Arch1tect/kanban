import Picker, {
  EmojiClickData,
  EmojiStyle,
  SuggestionMode,
} from "emoji-picker-react";
import { useCreateCommentReaction } from "../../../hooks/tasks/Comments/useCreateCommentReaction";

interface ReactionPickerProps {
  commentId: number;
  taskId: number;
}

const ReactionPicker: React.FC<ReactionPickerProps> = ({
  commentId,
  taskId,
}) => {
  const { mutate } = useCreateCommentReaction();

  const handleReaction = (emojiData: EmojiClickData) => {
    mutate({ commentId, reaction: emojiData.emoji, comment: { taskId } });
  };

  return (
    <div className="absolute right-0 z-10">
      <Picker
        onEmojiClick={handleReaction}
        suggestedEmojisMode={SuggestionMode.RECENT}
        reactionsDefaultOpen={true}
        lazyLoadEmojis={true}
        emojiStyle={EmojiStyle.NATIVE}
      />
    </div>
  );
};

export default ReactionPicker;
