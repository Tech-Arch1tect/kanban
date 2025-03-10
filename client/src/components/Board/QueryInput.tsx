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
  initialValue?: string;
  onSearch: (value: string) => void;
  placeholder?: string;
}

export default function QueryInput({
  initialValue = "",
  onSearch,
  placeholder,
}: QueryInputProps) {
  const [inputValue, setInputValue] = useState(initialValue);
  const inputRef = useRef<HTMLInputElement>(null);
  const containerRef = useRef<HTMLDivElement>(null);
  const [filteredSuggestions, setFilteredSuggestions] = useState<string[]>([]);
  const [activeSuggestionIndex, setActiveSuggestionIndex] = useState(0);

  useEffect(() => {
    const tokens = inputValue.trim().split(/\s+/);
    const lastToken = tokens[tokens.length - 1].toLowerCase();
    if (!lastToken) {
      setFilteredSuggestions([]);
      return;
    }
    const suggestions = [...FIELD_SUGGESTIONS, ...OPERATOR_SUGGESTIONS].filter(
      (item) =>
        item.toLowerCase().startsWith(lastToken) &&
        item.toLowerCase() !== lastToken
    );
    setFilteredSuggestions(suggestions);
    setActiveSuggestionIndex(0);
  }, [inputValue]);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        containerRef.current &&
        !containerRef.current.contains(event.target as Node)
      ) {
        setFilteredSuggestions([]);
      }
    };

    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, []);

  const handleSuggestionClick = useCallback(
    (suggestion: string) => {
      const tokens = inputValue.split(/\s+/);
      tokens[tokens.length - 1] = suggestion;
      setInputValue(tokens.join(" "));
      setFilteredSuggestions([]);
      inputRef.current?.focus();
    },
    [inputValue]
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
        if (filteredSuggestions.length > 0) {
          handleSuggestionClick(filteredSuggestions[activeSuggestionIndex]);
        } else {
          onSearch(inputValue);
        }
      }
    } else if (e.key === "Enter") {
      e.preventDefault();
      onSearch(inputValue);
    }
  };

  return (
    <div ref={containerRef} className="relative flex">
      <input
        ref={inputRef}
        type="text"
        value={inputValue}
        onChange={(e) => setInputValue(e.target.value)}
        onKeyDown={handleKeyDown}
        placeholder={placeholder}
        className="w-full p-2 border border-gray-300 dark:border-gray-600 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 focus:border-transparent bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
      />
      <button
        type="button"
        onClick={() => {
          setFilteredSuggestions([]);
          onSearch(inputValue);
        }}
        className="ml-2 px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600"
      >
        Search
      </button>
      <QuerySuggestions
        suggestions={filteredSuggestions}
        activeIndex={activeSuggestionIndex}
        onSuggestionClick={handleSuggestionClick}
      />
    </div>
  );
}
