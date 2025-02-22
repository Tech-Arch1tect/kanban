import { Link } from "@tanstack/react-router";
import { ChevronDownIcon } from "@heroicons/react/24/outline";
import { useDropdown } from "../../hooks/useDropdown";
import { useEffect, useRef } from "react";

const AdminDropdown = () => {
  const dropdown = useDropdown();
  const dropdownRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        dropdownRef.current &&
        !dropdownRef.current.contains(event.target as Node)
      ) {
        dropdown.closeDropdown();
      }
    };

    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, [dropdown]);

  return (
    <div className="relative" ref={dropdownRef}>
      <button
        onClick={dropdown.toggleDropdown}
        className="text-white dark:text-gray-100 text-lg flex items-center space-x-2"
      >
        Admin <ChevronDownIcon className="h-5 w-5" />
      </button>
      {dropdown.isOpen && (
        <div className="absolute right-0 mt-2 w-48 bg-white dark:bg-gray-800 shadow-xl rounded-lg p-2">
          <Link
            to="/admin/users"
            className="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-700 dark:text-gray-100"
          >
            Users
          </Link>
          <Link
            to="/admin/boards"
            className="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-700 dark:text-gray-100"
          >
            Boards
          </Link>
        </div>
      )}
    </div>
  );
};

export default AdminDropdown;
