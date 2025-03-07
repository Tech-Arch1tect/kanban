import { createLazyFileRoute } from "@tanstack/react-router";
import { useEffect, useState } from "react";
import { authApi, miscApi } from "../../lib/api";
import { QRCodeSVG } from "qrcode.react";
import { ModelsUser } from "../../typescript-fetch-client";

export const Route = createLazyFileRoute("/profile/2fa")({
  component: () => <TwoFAComponent />,
});

const TwoFAComponent = () => {
  const [is2FAEnabled, setIs2FAEnabled] = useState(false);
  const [appName, setAppName] = useState("App Name Not Found");
  const [secret, setSecret] = useState("");
  const [code, setCode] = useState("");
  const [message, setMessage] = useState("");
  const [profile, setProfile] = useState<ModelsUser | undefined>(undefined);

  useEffect(() => {
    fetchAppName();
    fetchProfile();
  }, []);

  async function fetchAppName() {
    try {
      const appName = await miscApi.apiV1MiscAppnameGet();
      setAppName(appName.appName || "App Name Not Found");
    } catch (error) {
      console.error("Error fetching app name:", error);
    }
  }

  async function fetchProfile() {
    try {
      const profile = await authApi.apiV1AuthProfileGet();
      setProfile(profile);
      setIs2FAEnabled(profile.totpEnabled || false);
    } catch (error) {
      console.error("Error fetching profile:", error);
    }
  }

  async function generateQRCode() {
    try {
      const resp = await authApi.apiV1AuthTotpGeneratePost();
      setSecret(resp.secret || "");
    } catch (error) {
      console.error("Error generating QR code:", error);
    }
  }

  async function enable2FA() {
    try {
      const resp = await authApi.apiV1AuthTotpEnablePost({ request: { code } });
      setMessage(resp.message || "2FA enabled successfully");
      setIs2FAEnabled(true);
      setSecret("");
      setCode("");
    } catch (error) {
      console.error("Error enabling 2FA:", error);
      setMessage("Failed to enable 2FA. Please try again.");
    }
  }

  async function disable2FA() {
    try {
      const resp = await authApi.apiV1AuthTotpDisablePost({
        request: { code },
      });
      setMessage(resp.message || "2FA disabled successfully");
      setIs2FAEnabled(false);
      setCode("");
    } catch (error) {
      console.error("Error disabling 2FA:", error);
      setMessage("Failed to disable 2FA. Please try again.");
    }
  }

  return (
    <div className="max-w-md mx-auto p-6 bg-white dark:bg-gray-800 shadow-lg rounded-lg">
      <h2 className="text-2xl font-bold mb-4 text-gray-900 dark:text-gray-200">
        Two-Factor Authentication
      </h2>
      {is2FAEnabled ? (
        <div>
          <p className="mb-4 text-gray-700 dark:text-gray-300">
            2FA is currently enabled.
          </p>
          <input
            type="text"
            value={code}
            onChange={(e) => setCode(e.target.value)}
            placeholder="Enter 2FA code"
            className="w-full p-2 mb-4 border border-gray-300 dark:border-gray-600 rounded bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
          />
          <button
            onClick={disable2FA}
            className="w-full bg-red-500 dark:bg-red-600 text-white p-2 rounded hover:bg-red-600 dark:hover:bg-red-700 transition duration-300"
          >
            Disable 2FA
          </button>
        </div>
      ) : (
        <div>
          <p className="mb-4 text-gray-700 dark:text-gray-300">
            2FA is currently disabled.
          </p>
          {secret ? (
            <div>
              <p className="mb-4 text-gray-700 dark:text-gray-300">
                Scan this QR code with your authenticator app:
              </p>
              <div className="flex justify-center mb-4">
                <QRCodeSVG
                  value={`otpauth://totp/${appName}:${profile?.email}?secret=${secret}&issuer=${appName}`}
                  size={200}
                />
              </div>
              <p className="mb-4 text-gray-700 dark:text-gray-300">
                Or enter this secret manually: {secret}
              </p>
              <input
                type="text"
                value={code}
                onChange={(e) => setCode(e.target.value)}
                placeholder="Enter 2FA code"
                className="w-full p-2 mb-4 border border-gray-300 dark:border-gray-600 rounded bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
              />
              <button
                onClick={enable2FA}
                className="w-full bg-green-500 dark:bg-green-600 text-white p-2 rounded hover:bg-green-600 dark:hover:bg-green-700 transition duration-300"
              >
                Enable 2FA
              </button>
            </div>
          ) : (
            <button
              onClick={generateQRCode}
              className="w-full bg-blue-500 dark:bg-blue-600 text-white p-2 rounded hover:bg-blue-600 dark:hover:bg-blue-700 transition duration-300"
            >
              Generate QR Code
            </button>
          )}
        </div>
      )}
      {message && (
        <p className="mt-4 text-sm text-gray-600 dark:text-gray-400">
          {message}
        </p>
      )}
    </div>
  );
};
