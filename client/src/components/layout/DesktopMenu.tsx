import { Link } from "@tanstack/react-router";
import BoardsSelect from "./BoardsSelect";
import ProfileDropdown from "./ProfileDropdown";
import AdminDropdown from "./AdminDropdown";
import { MoonIcon, SunIcon } from "@heroicons/react/24/outline";

interface DesktopMenuProps {
  toggleTheme: () => void;
  settings: { theme?: string } | null;
  isAdmin: boolean;
  handleLogout: () => void;
}

const DesktopMenu = ({
  toggleTheme,
  settings,
  isAdmin,
  handleLogout,
}: DesktopMenuProps) => {
  return (
    <div className="hidden lg:flex items-center justify-between w-full">
      {/* Left Section */}
      <div className="flex items-center space-x-4">
        <Link
          to="/about"
          className="ml-2 text-gray-200 dark:text-gray-400 hover:text-white dark:hover:text-gray-100 text-lg"
        >
          About
        </Link>
      </div>

      {/* Centre Section */}
      <div className="flex-1 flex items-center justify-center">
        <BoardsSelect />
      </div>

      {/* Right Section */}
      <div className="flex items-center space-x-4">
        <button
          onClick={toggleTheme}
          className="p-2 rounded-full hover:bg-blue-700 dark:hover:bg-gray-700 transition-all"
          aria-label="Toggle Dark Mode"
        >
          {settings?.theme === "dark" ? (
            <SunIcon className="h-6 w-6 text-yellow-300 dark:text-yellow-400" />
          ) : (
            <MoonIcon className="h-6 w-6 text-gray-200 dark:text-gray-400" />
          )}
        </button>
        <ProfileDropdown />
        {isAdmin && <AdminDropdown />}
        <button
          onClick={handleLogout}
          className="text-white dark:text-gray-100 text-lg"
        >
          Logout
        </button>
      </div>
    </div>
  );
};

export default DesktopMenu;
