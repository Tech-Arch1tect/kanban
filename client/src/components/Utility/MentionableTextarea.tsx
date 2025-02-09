import React, { useState, useRef } from "react";

interface MentionableTextareaProps {
  value: string;
  setValue: React.Dispatch<React.SetStateAction<string>>;
  placeholder: string;
  suggestions: string[];
  onSelectSuggestion: (username: string) => void;
  containerClassName?: string;
  textareaClassName?: string;
}

const MentionableTextarea: React.FC<MentionableTextareaProps> = ({
  value,
  setValue,
  placeholder,
  suggestions,
  onSelectSuggestion,
  containerClassName = "",
  textareaClassName = "",
}) => {
  const [filteredSuggestions, setFilteredSuggestions] = useState<string[]>([]);
  const [showSuggestions, setShowSuggestions] = useState(false);
  const [mentionIndex, setMentionIndex] = useState<number | null>(null);
  const [activeIndex, setActiveIndex] = useState<number>(0);

  const textareaRef = useRef<HTMLTextAreaElement | null>(null);

  const handleInputChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    const inputValue = e.target.value;
    setValue(inputValue);

    const cursorPosition = textareaRef.current?.selectionStart || 0;
    const lastAtSign = inputValue.lastIndexOf("@", cursorPosition - 1);

    if (lastAtSign !== -1) {
      const mentionText = inputValue.slice(lastAtSign + 1, cursorPosition);
      if (mentionText.trim().length > 0) {
        const filtered = suggestions.filter((username) =>
          username.toLowerCase().startsWith(mentionText.toLowerCase())
        );
        setFilteredSuggestions(filtered);
        setShowSuggestions(true);
        setMentionIndex(lastAtSign);
        setActiveIndex(0);
      } else {
        setShowSuggestions(false);
      }
    } else {
      setShowSuggestions(false);
    }
  };

  const handleSuggestionClick = (username: string) => {
    if (mentionIndex !== null) {
      const beforeMention = value.slice(0, mentionIndex);
      const afterMention = value.slice(
        textareaRef.current?.selectionStart || 0
      );
      setValue(`${beforeMention}@${username} ${afterMention}`);
      setShowSuggestions(false);
      setFilteredSuggestions([]);
      onSelectSuggestion(username);
    }
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
    if (showSuggestions) {
      if (e.key === "ArrowDown") {
        e.preventDefault();
        setActiveIndex((prev) =>
          prev === filteredSuggestions.length - 1 ? 0 : prev + 1
        );
      } else if (e.key === "ArrowUp") {
        e.preventDefault();
        setActiveIndex((prev) =>
          prev === 0 ? filteredSuggestions.length - 1 : prev - 1
        );
      } else if (e.key === "Enter") {
        if (filteredSuggestions[activeIndex]) {
          e.preventDefault();
          handleSuggestionClick(filteredSuggestions[activeIndex]);
        }
      } else if (e.key === "Tab") {
        e.preventDefault();
        if (filteredSuggestions[activeIndex]) {
          handleSuggestionClick(filteredSuggestions[activeIndex]);
        }
      } else if (e.key === "Escape") {
        setShowSuggestions(false);
      }
    }
  };

  return (
    <div className={`relative ${containerClassName}`}>
      <textarea
        ref={textareaRef}
        className={`w-full border border-gray-300 dark:border-gray-600 rounded-md p-3 text-gray-800 dark:text-gray-200 focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 focus:outline-none placeholder-gray-400 dark:placeholder-gray-500 bg-white dark:bg-gray-700 ${textareaClassName}`}
        placeholder={placeholder}
        value={value}
        onChange={handleInputChange}
        onKeyDown={handleKeyDown}
        rows={4}
      ></textarea>
      {showSuggestions && filteredSuggestions.length > 0 && (
        <ul className="absolute z-10 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-lg w-full max-h-48 overflow-y-auto mt-1">
          {filteredSuggestions.map((username, index) => (
            <li
              key={index}
              className={`px-3 py-2 cursor-pointer ${
                activeIndex === index
                  ? "bg-blue-500 text-white"
                  : "hover:bg-gray-200 dark:hover:bg-gray-600"
              }`}
              onMouseDown={() => handleSuggestionClick(username)}
            >
              {username}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default MentionableTextarea;
