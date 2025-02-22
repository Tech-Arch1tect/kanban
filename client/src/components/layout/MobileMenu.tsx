import { Link } from "@tanstack/react-router";
import BoardsSelect from "./BoardsSelect";
import { MoonIcon, SunIcon } from "@heroicons/react/24/outline";

interface MobileMenuProps {
  toggleTheme: () => void;
  settings: { theme?: string } | null;
  handleLogout: () => void;
}

const MobileMenu = ({
  toggleTheme,
  settings,
  handleLogout,
}: MobileMenuProps) => {
  return (
    <div className="lg:hidden bg-blue-900 dark:bg-gray-800 p-4 space-y-4">
      <Link
        to="/about"
        className="block text-gray-200 dark:text-gray-400 hover:text-white"
      >
        About
      </Link>
      <BoardsSelect />
      <button
        onClick={toggleTheme}
        className="flex items-center space-x-2 text-white dark:text-gray-100"
      >
        {settings?.theme === "dark" ? (
          <SunIcon className="h-6 w-6 text-yellow-300 dark:text-yellow-400" />
        ) : (
          <MoonIcon className="h-6 w-6 text-gray-200 dark:text-gray-400" />
        )}
        <span>Toggle Theme</span>
      </button>
      <button onClick={handleLogout} className="text-white dark:text-gray-100">
        Logout
      </button>
    </div>
  );
};

export default MobileMenu;
