import { useState } from "react";
import MentionableTextarea from "../../Utility/MentionableTextarea";
import RenderMarkdown from "../../Utility/RenderMarkdown";
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
  const [isPreview, setIsPreview] = useState(false);

  const handleSelectSuggestion = (username: string) => {
    console.log(`Selected username: ${username}`);
  };

  const handleSubmit = (e: React.FormEvent) => {
    onSubmit(e);
    setIsPreview(false);
  };

  return (
    <form onSubmit={handleSubmit} className="mt-6 space-y-3">
      {isPreview ? (
        <div className="p-2 border rounded bg-gray-50 dark:bg-gray-800">
          <RenderMarkdown markdown={value} />
        </div>
      ) : (
        !isLoading && (
          <MentionableTextarea
            value={value}
            setValue={setValue}
            placeholder={placeholder}
            suggestions={usernames}
            onSelectSuggestion={handleSelectSuggestion}
            containerClassName="mb-4"
            textareaClassName="shadow-sm"
          />
        )
      )}

      <div className="flex justify-between">
        <button
          type="button"
          onClick={() => setIsPreview(!isPreview)}
          disabled={!isPreview && !value.trim()}
          className={`px-3 py-1 text-sm text-white bg-blue-500 rounded hover:bg-blue-600 ${
            !isPreview && !value.trim() ? "opacity-50 cursor-not-allowed" : ""
          }`}
        >
          {isPreview ? "Back to Edit" : "Preview"}
        </button>
        <button
          type="submit"
          disabled={!value.trim()}
          className={`py-2 px-4 rounded-md text-white ${
            value.trim()
              ? "bg-blue-500 hover:bg-blue-600 dark:bg-blue-600 dark:hover:bg-blue-700"
              : "bg-gray-300 dark:bg-gray-600 cursor-not-allowed"
          }`}
        >
          Submit
        </button>
      </div>
    </form>
  );
};

export default CommentForm;
