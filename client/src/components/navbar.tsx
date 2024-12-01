import { Link, useLocation, useNavigate } from "@tanstack/react-router";
import { useEffect, useRef, useState } from "react";
import { authApi } from "../lib/api";

const Navbar = () => {
  const [isAdminDropdownOpen, setIsAdminDropdownOpen] = useState(false);
  const [isProfileDropdownOpen, setIsProfileDropdownOpen] = useState(false);
  const navigate = useNavigate();
  const location = useLocation();
  const [isAdmin, setIsAdmin] = useState(false);
  const dropdownRef = useRef<HTMLDivElement>(null);

  // If the user is on the login or register page, return an empty fragment
  if (["/login", "/register", "/password-reset"].includes(location.pathname)) {
    return <></>;
  }

  const toggleAdminDropdown = () => {
    setIsAdminDropdownOpen(!isAdminDropdownOpen);
  };

  const toggleProfileDropdown = () => {
    setIsProfileDropdownOpen(!isProfileDropdownOpen);
  };

  const handleLogout = async () => {
    try {
      await authApi.apiV1AuthLogoutPost();
      navigate({ to: "/login" });
    } catch (error) {
      console.error("Error logging out:", error);
    }
  };

  useEffect(() => {
    const fetchUserProfile = async () => {
      try {
        const profile = await authApi.apiV1AuthProfileGet();
        setIsAdmin(profile.role === "admin");
      } catch (error) {
        console.error("Error fetching profile:", error);
        if (!["/login", "/register"].includes(location.pathname)) {
          navigate({ to: "/login" });
        }
      }
    };

    fetchUserProfile();
  }, [location]);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        dropdownRef.current &&
        !dropdownRef.current.contains(event.target as Node)
      ) {
        setIsAdminDropdownOpen(false);
        setIsProfileDropdownOpen(false);
      }
    };

    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, [dropdownRef]);

  return (
    <nav className="bg-blue-800 shadow-lg">
      <div className="container mx-auto px-6 py-3 flex justify-between items-center">
        <div className="flex items-center space-x-4">
          <Link
            to="/"
            activeProps={{ className: "text-white font-bold" }}
            inactiveProps={{ className: "text-gray-200 hover:text-white" }}
            className="text-lg"
          >
            Home
          </Link>
          <Link
            to="/about"
            activeProps={{ className: "text-white font-bold" }}
            inactiveProps={{ className: "text-gray-200 hover:text-white" }}
            className="text-lg"
          >
            About
          </Link>
        </div>
        <div className="flex items-center space-x-4" ref={dropdownRef}>
          {/* User Profile Dropdown */}
          <div className="relative">
            <button
              onClick={toggleProfileDropdown}
              className="text-white text-lg font-medium flex items-center space-x-2"
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
            {isProfileDropdownOpen && (
              <div className="dropdown absolute right-0 mt-2 w-48 bg-white shadow-xl rounded">
                <Link
                  to="/profile/profile"
                  onClick={() => setIsProfileDropdownOpen(false)}
                  className="dropdown-item block px-4 py-2 hover:bg-gray-100"
                >
                  Profile
                </Link>
                <Link
                  to="/profile/2fa"
                  onClick={() => setIsProfileDropdownOpen(false)}
                  className="dropdown-item block px-4 py-2 hover:bg-gray-100"
                >
                  Manage 2FA
                </Link>
              </div>
            )}
          </div>
          {/* Admin Dropdown */}
          {isAdmin && (
            <div className="relative">
              <button
                onClick={toggleAdminDropdown}
                className="text-white text-lg font-medium flex items-center space-x-2"
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
              {isAdminDropdownOpen && (
                <div className="dropdown absolute right-0 mt-2 w-48 bg-white shadow-xl rounded">
                  <Link
                    to="/admin/users"
                    onClick={() => setIsAdminDropdownOpen(false)}
                    className="dropdown-item block px-4 py-2 hover:bg-gray-100"
                  >
                    Users
                  </Link>
                </div>
              )}
            </div>
          )}
          <button
            onClick={handleLogout}
            className="text-white text-lg font-medium hover:text-gray-200"
          >
            Logout
          </button>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
