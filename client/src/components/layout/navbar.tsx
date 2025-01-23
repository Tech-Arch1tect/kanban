import { Link, useLocation, useNavigate } from "@tanstack/react-router";
import { useEffect } from "react";
import { useUserProfile } from "../../hooks/useUserProfile";
import { useAuth } from "../../hooks/auth/useAuth";
import { useDropdown } from "../../hooks/useDropdown";
import BoardsSelect from "./BoardsSelect";
import { ToastContainer } from "react-toastify";
import { MoonIcon, SunIcon } from "@heroicons/react/24/outline";
import { useServerSettings } from "../../context/ServerSettingsContext";

const Navbar = () => {
  const { settings, updateSettings } = useServerSettings();

  const navigate = useNavigate();
  const location = useLocation();

  const { profile, error } = useUserProfile();
  const { handleLogout, isAdmin } = useAuth(profile);

  const profileDropdown = useDropdown();
  const adminDropdown = useDropdown();

  useEffect(() => {
    if (settings?.theme === "dark") {
      document.body.classList.add("dark");
    } else {
      document.body.classList.remove("dark");
    }
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
    updateSettings({
      theme: settings?.theme === "dark" ? "light" : "dark",
    });
  };

  return (
    <>
      <nav className="bg-blue-800 dark:bg-gray-900 shadow-lg">
        <div className="container mx-auto px-6 py-3 flex justify-between items-center">
          <div className="flex items-center space-x-4">
            <Link
              to="/"
              activeProps={{
                className: "text-white dark:text-gray-100 font-bold",
              }}
              inactiveProps={{
                className:
                  "text-gray-200 dark:text-gray-400 hover:text-white dark:hover:text-gray-100",
              }}
              className="text-lg"
            >
              Home
            </Link>
            <Link
              to="/about"
              activeProps={{
                className: "text-white dark:text-gray-100 font-bold",
              }}
              inactiveProps={{
                className:
                  "text-gray-200 dark:text-gray-400 hover:text-white dark:hover:text-gray-100",
              }}
              className="text-lg"
            >
              About
            </Link>
          </div>

          <div className="flex items-center space-x-4">
            <BoardsSelect />
          </div>

          <div className="flex items-center space-x-4">
            <button
              onClick={toggleTheme}
              className="p-2 rounded-full hover:bg-blue-700 dark:hover:bg-gray-700 focus:outline-none"
              aria-label="Toggle Dark Mode"
            >
              {settings?.theme === "dark" ? (
                <SunIcon className="h-6 w-6 text-yellow-300 dark:text-yellow-400" />
              ) : (
                <MoonIcon className="h-6 w-6 text-gray-200 dark:text-gray-400" />
              )}
            </button>

            {/* User Profile Dropdown */}
            <div className="relative" ref={profileDropdown.ref}>
              <button
                onClick={profileDropdown.toggleDropdown}
                className="text-white dark:text-gray-100 text-lg font-medium flex items-center space-x-2"
              >
                Settings
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-4 w-4"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M19 9l-7 7-7-7"
                  />
                </svg>
              </button>
              {profileDropdown.isOpen && (
                <div className="dropdown absolute right-0 mt-2 w-48 bg-white dark:bg-gray-800 shadow-xl rounded">
                  <Link
                    to="/profile/profile"
                    onClick={profileDropdown.closeDropdown}
                    className="dropdown-item block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-700 dark:text-gray-100"
                  >
                    Profile
                  </Link>
                  <Link
                    to="/profile/2fa"
                    onClick={profileDropdown.closeDropdown}
                    className="dropdown-item block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-700 dark:text-gray-100"
                  >
                    Manage 2FA
                  </Link>
                </div>
              )}
            </div>

            {/* Admin Dropdown */}
            {isAdmin && (
              <div className="relative" ref={adminDropdown.ref}>
                <button
                  onClick={adminDropdown.toggleDropdown}
                  className="text-white dark:text-gray-100 text-lg font-medium flex items-center space-x-2"
                >
                  Admin
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="h-4 w-4"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M19 9l-7 7-7-7"
                    />
                  </svg>
                </button>
                {adminDropdown.isOpen && (
                  <div className="dropdown absolute right-0 mt-2 w-48 bg-white dark:bg-gray-800 shadow-xl rounded">
                    <Link
                      to="/admin/users"
                      onClick={adminDropdown.closeDropdown}
                      className="dropdown-item block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-700 dark:text-gray-100"
                    >
                      Users
                    </Link>
                    <Link
                      to="/admin/boards"
                      onClick={adminDropdown.closeDropdown}
                      className="dropdown-item block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-700 dark:text-gray-100"
                    >
                      Boards
                    </Link>
                  </div>
                )}
              </div>
            )}

            <button
              onClick={handleLogout}
              className="text-white dark:text-gray-100 text-lg font-medium hover:text-gray-200 dark:hover:text-gray-300"
            >
              Logout
            </button>
          </div>
        </div>
      </nav>
      <ToastContainer position="bottom-right" newestOnTop theme="colored" />
    </>
  );
};

export default Navbar;
