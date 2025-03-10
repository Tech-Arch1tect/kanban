import { useState, useRef, useCallback, useEffect } from "react";
import QuerySuggestions from "./QuerySuggestions";

// todo: get from backend
const FIELD_SUGGESTIONS = ["status", "title", "search", "assignee", "creator"];
const OPERATOR_SUGGESTIONS = [
  "==",
  "!=",
  ">",
  "<",
  ">=",
  "<=",
  "like",
  "AND",
  "OR",
];

interface QueryInputProps {
  value: string;
  onChange: (value: string) => void;
  placeholder?: string;
}

export default function QueryInput({
  value,
  onChange,
  placeholder,
}: QueryInputProps) {
  const inputRef = useRef<HTMLInputElement>(null);
  const [filteredSuggestions, setFilteredSuggestions] = useState<string[]>([]);
  const [activeSuggestionIndex, setActiveSuggestionIndex] = useState(0);

  useEffect(() => {
    const tokens = value.split(/\s+/);
    const lastToken = tokens[tokens.length - 1].toLowerCase();
    const suggestions = [...FIELD_SUGGESTIONS, ...OPERATOR_SUGGESTIONS].filter(
      (item) =>
        item.toLowerCase().startsWith(lastToken) &&
        item.toLowerCase() !== lastToken
    );
    setFilteredSuggestions(suggestions);
    setActiveSuggestionIndex(0);
  }, [value]);

  const handleSuggestionClick = useCallback(
    (suggestion: string) => {
      const tokens = value.split(/\s+/);
      tokens[tokens.length - 1] = suggestion;
      onChange(tokens.join(" "));
      inputRef.current?.focus();
    },
    [value, onChange]
  );

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (filteredSuggestions.length > 0) {
      if (e.key === "ArrowDown") {
        e.preventDefault();
        setActiveSuggestionIndex((prev) =>
          prev === filteredSuggestions.length - 1 ? 0 : prev + 1
        );
      } else if (e.key === "ArrowUp") {
        e.preventDefault();
        setActiveSuggestionIndex((prev) =>
          prev === 0 ? filteredSuggestions.length - 1 : prev - 1
        );
      } else if (e.key === "Enter") {
        e.preventDefault();
        handleSuggestionClick(filteredSuggestions[activeSuggestionIndex]);
      }
    }
  };

  return (
    <div className="relative">
      <input
        ref={inputRef}
        type="text"
        value={value}
        onChange={(e) => onChange(e.target.value)}
        onKeyDown={handleKeyDown}
        placeholder={placeholder}
        className="w-full p-2 border border-gray-300 dark:border-gray-600 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 focus:border-transparent bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
      />
      <QuerySuggestions
        suggestions={filteredSuggestions}
        activeIndex={activeSuggestionIndex}
        onSuggestionClick={handleSuggestionClick}
      />
    </div>
  );
}
