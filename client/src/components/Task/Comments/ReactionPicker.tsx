import { useState } from "react";
import Picker, { EmojiStyle, SuggestionMode } from "emoji-picker-react";
import { useCreateCommentReaction } from "../../../hooks/tasks/Comments/useCreateCommentReaction";
import { RocketLaunchIcon } from "@heroicons/react/24/outline";

interface ReactionPickerProps {
  commentId: number;
  taskId: number;
}

const ReactionPicker: React.FC<ReactionPickerProps> = ({
  commentId,
  taskId,
}) => {
  const { mutate } = useCreateCommentReaction();
  const [showPicker, setShowPicker] = useState(false);

  const handleReaction = (emojiData: any, event: MouseEvent) => {
    mutate({ commentId, reaction: emojiData.emoji, comment: { taskId } });
    setShowPicker(false);
  };

  return (
    <div className="relative inline-block">
      <button
        onClick={() => setShowPicker(!showPicker)}
        className="p-1 text-gray-500 hover:text-gray-700"
        aria-label="Add reaction"
      >
        <RocketLaunchIcon className="h-6 w-6" />
      </button>
      {showPicker && (
        <div className="absolute right-0 mt-2 z-10">
          <Picker
            onEmojiClick={handleReaction}
            suggestedEmojisMode={SuggestionMode.RECENT}
            reactionsDefaultOpen={true}
            lazyLoadEmojis={true}
            emojiStyle={EmojiStyle.NATIVE}
          />
        </div>
      )}
    </div>
  );
};

export default ReactionPicker;
