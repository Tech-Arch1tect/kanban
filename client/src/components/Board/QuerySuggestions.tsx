interface QuerySuggestionsProps {
  suggestions: string[];
  activeIndex: number;
  onSuggestionClick: (suggestion: string) => void;
}

export default function QuerySuggestions({
  suggestions,
  activeIndex,
  onSuggestionClick,
}: QuerySuggestionsProps) {
  if (suggestions.length === 0) return null;

  return (
    <ul
      className="absolute left-0 right-0 top-full mt-1 bg-white dark:bg-gray-800 border border-gray-300 dark:border-gray-600 rounded-md shadow-lg z-10"
      style={{ maxHeight: "150px", overflowY: "auto" }}
    >
      {suggestions.map((suggestion, index) => (
        <li
          key={suggestion}
          className={`p-2 cursor-pointer 
            ${index === activeIndex ? "bg-blue-500 text-white" : "hover:bg-gray-100 dark:hover:bg-gray-700 text-gray-900 dark:text-gray-100"}`}
          onMouseDown={() => onSuggestionClick(suggestion)}
        >
          {suggestion}
        </li>
      ))}
    </ul>
  );
}
