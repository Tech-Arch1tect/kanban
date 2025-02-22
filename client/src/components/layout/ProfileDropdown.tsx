import { Link } from "@tanstack/react-router";
import { ChevronDownIcon } from "@heroicons/react/24/outline";
import { useDropdown } from "../../hooks/useDropdown";
import { useEffect, useRef } from "react";

const ProfileDropdown = () => {
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
        Settings <ChevronDownIcon className="h-5 w-5" />
      </button>
      {dropdown.isOpen && (
        <div className="absolute right-0 mt-2 w-48 bg-white dark:bg-gray-800 shadow-xl rounded-lg p-2">
          <Link
            to="/profile/profile"
            onClick={dropdown.closeDropdown}
            className="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-700 dark:text-gray-100"
          >
            Profile
          </Link>
          <Link
            to="/profile/2fa"
            onClick={dropdown.closeDropdown}
            className="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-700 dark:text-gray-100"
          >
            Manage 2FA
          </Link>
          <Link
            to="/profile/notifications"
            onClick={dropdown.closeDropdown}
            className="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-700 dark:text-gray-100"
          >
            Manage Notifications
          </Link>
        </div>
      )}
    </div>
  );
};

export default ProfileDropdown;
