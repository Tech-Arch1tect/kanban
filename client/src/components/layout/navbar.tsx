import { Link, useLocation, useNavigate } from "@tanstack/react-router";
import { useEffect, useState } from "react";
import { useUserProfile } from "../../hooks/useUserProfile";
import { useAuth } from "../../hooks/auth/useAuth";
import { useDropdown } from "../../hooks/useDropdown";
import BoardsSelect from "./BoardsSelect";
import { ToastContainer } from "react-toastify";
import {
  MoonIcon,
  SunIcon,
  ChevronDownIcon,
  Bars3Icon,
  XMarkIcon,
} from "@heroicons/react/24/outline";
import { useServerSettings } from "../../context/ServerSettingsContext";

const Navbar = () => {
  const { settings, updateSettings } = useServerSettings();
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

  const navigate = useNavigate();
  const location = useLocation();

  const { profile, error } = useUserProfile();
  const { handleLogout, isAdmin } = useAuth(profile);

  const profileDropdown = useDropdown();
  const adminDropdown = useDropdown();

  useEffect(() => {
    document.body.classList.toggle("dark", settings?.theme === "dark");
  }, [settings?.theme]);

  useEffect(() => {
    if (
      error &&
      !["/login", "/register"].includes(location.pathname) &&
      !location.pathname.startsWith("/password-reset")
    ) {
      navigate({ to: "/login" });
    }
  }, [error, navigate]);

  if (["/login", "/register", "/password-reset"].includes(location.pathname)) {
    return null;
  }

  const toggleTheme = () => {
    updateSettings({ theme: settings?.theme === "dark" ? "light" : "dark" });
  };

  return (
    <>
      <nav className="bg-blue-800 dark:bg-gray-900 shadow-lg">
        <div className="container mx-auto px-6 py-3 flex justify-between items-center">
          <div className="flex items-center space-x-4">
            <button
              onClick={() => setIsMobileMenuOpen(!isMobileMenuOpen)}
              className="lg:hidden text-white dark:text-gray-100"
            >
              {isMobileMenuOpen ? (
                <XMarkIcon className="h-7 w-7" />
              ) : (
                <Bars3Icon className="h-7 w-7" />
              )}
            </button>
            <Link to="/" className="text-white text-lg font-bold">
              Home
            </Link>
          </div>

          {/* Desktop Menu */}
          <div className="hidden lg:flex items-center space-x-4">
            <Link
              to="/about"
              className="text-gray-200 dark:text-gray-400 hover:text-white dark:hover:text-gray-100 text-lg"
            >
              About
            </Link>
            <BoardsSelect />
          </div>

          {/* Right Actions */}
          <div className="hidden lg:flex items-center space-x-4">
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

            {/* Profile Dropdown */}
            <div className="relative">
              <button
                onClick={profileDropdown.toggleDropdown}
                className="text-white dark:text-gray-100 text-lg flex items-center space-x-2"
              >
                Settings <ChevronDownIcon className="h-5 w-5" />
              </button>
              {profileDropdown.isOpen && (
                <div className="absolute right-0 mt-2 w-48 bg-white dark:bg-gray-800 shadow-xl rounded-lg p-2">
                  <Link
                    to="/profile/profile"
                    onClick={profileDropdown.closeDropdown}
                    className="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-700 dark:text-gray-100"
                  >
                    Profile
                  </Link>
                  <Link
                    to="/profile/2fa"
                    onClick={profileDropdown.closeDropdown}
                    className="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-700 dark:text-gray-100"
                  >
                    Manage 2FA
                  </Link>
                  <Link
                    to="/profile/notifications"
                    onClick={profileDropdown.closeDropdown}
                    className="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-700 dark:text-gray-100"
                  >
                    Manage Notifications
                  </Link>
                </div>
              )}
            </div>

            {/* Admin Dropdown */}
            {isAdmin && (
              <div className="relative">
                <button
                  onClick={adminDropdown.toggleDropdown}
                  className="text-white dark:text-gray-100 text-lg flex items-center space-x-2"
                >
                  Admin <ChevronDownIcon className="h-5 w-5" />
                </button>
                {adminDropdown.isOpen && (
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
            )}

            <button
              onClick={handleLogout}
              className="text-white dark:text-gray-100 text-lg"
            >
              Logout
            </button>
          </div>
        </div>

        {/* Mobile Menu */}
        {isMobileMenuOpen && (
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
            <button
              onClick={handleLogout}
              className="text-white dark:text-gray-100"
            >
              Logout
            </button>
          </div>
        )}
      </nav>

      <ToastContainer position="bottom-right" newestOnTop theme="colored" />
    </>
  );
};

export default Navbar;
