import { Link, useLocation, useNavigate } from "@tanstack/react-router";
import { useEffect, useState } from "react";
import { ToastContainer } from "react-toastify";
import { Bars3Icon, XMarkIcon } from "@heroicons/react/24/outline";
import { useUserProfile } from "../../hooks/useUserProfile";
import { useAuth } from "../../hooks/auth/useAuth";
import { useServerSettings } from "../../context/ServerSettingsContext";
import DesktopMenu from "./DesktopMenu";
import MobileMenu from "./MobileMenu";

const Navbar = () => {
  const { settings, updateSettings } = useServerSettings();
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

  const navigate = useNavigate();
  const location = useLocation();

  const { profile, error } = useUserProfile();
  const { handleLogout, isAdmin } = useAuth(profile);

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
  }, [error, navigate, location.pathname]);

  if (["/login", "/register", "/password-reset"].includes(location.pathname)) {
    return null;
  }

  const toggleTheme = () =>
    updateSettings({ theme: settings?.theme === "dark" ? "light" : "dark" });

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
          <DesktopMenu
            toggleTheme={toggleTheme}
            settings={settings ?? null}
            isAdmin={isAdmin}
            handleLogout={handleLogout}
          />
        </div>
        {isMobileMenuOpen && (
          <MobileMenu
            toggleTheme={toggleTheme}
            settings={settings ?? null}
            handleLogout={handleLogout}
          />
        )}
      </nav>

      <ToastContainer position="bottom-right" newestOnTop theme="colored" />
    </>
  );
};

export default Navbar;
