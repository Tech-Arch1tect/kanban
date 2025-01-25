import MentionableTextarea from "../../Utility/MentionableTextarea";
import { useBoardUsernames } from "../../../hooks/boards/useBoardUsernames";

interface CommentFormProps {
  onSubmit: (e: React.FormEvent) => void;
  value: string;
  setValue: React.Dispatch<React.SetStateAction<string>>;
  placeholder: string;
  boardId: number;
}

const CommentForm: React.FC<CommentFormProps> = ({
  onSubmit,
  value,
  setValue,
  placeholder,
  boardId,
}) => {
  const { usernames, isLoading } = useBoardUsernames(boardId);

  const handleSelectSuggestion = (username: string) => {
    console.log(`Selected username: ${username}`);
  };

  return (
    <form onSubmit={onSubmit} className="mt-6 space-y-3">
      {!isLoading && (
        <MentionableTextarea
          value={value}
          setValue={setValue}
          placeholder={placeholder}
          suggestions={usernames}
          onSelectSuggestion={handleSelectSuggestion}
        />
      )}
      <button
        type="submit"
        disabled={!value.trim()}
        className={`w-full py-2 px-4 rounded-md text-white ${
          value.trim()
            ? "bg-blue-500 hover:bg-blue-600 dark:bg-blue-600 dark:hover:bg-blue-700"
            : "bg-gray-300 dark:bg-gray-600 cursor-not-allowed"
        }`}
      >
        Submit
      </button>
    </form>
  );
};

export default CommentForm;
