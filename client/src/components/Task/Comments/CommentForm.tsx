interface CommentFormProps {
  onSubmit: (e: React.FormEvent) => void;
  value: string;
  setValue: React.Dispatch<React.SetStateAction<string>>;
  placeholder: string;
}

const CommentForm: React.FC<CommentFormProps> = ({
  onSubmit,
  value,
  setValue,
  placeholder,
}) => (
  <form onSubmit={onSubmit} className="mt-6 space-y-3">
    <textarea
      className="w-full border border-gray-300 dark:border-gray-600 rounded-md p-3 text-gray-800 dark:text-gray-200 focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 focus:outline-none placeholder-gray-400 dark:placeholder-gray-500 bg-white dark:bg-gray-700"
      placeholder={placeholder}
      value={value}
      onChange={(e) => setValue(e.target.value)}
      rows={4}
    ></textarea>
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

export default CommentForm;
