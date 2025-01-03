import { createFileRoute } from "@tanstack/react-router";
import ForgotPassword from "../components/Auth/ForgotPassword";

export const Route = createFileRoute("/password-reset")({
  component: () => {
    const { email, code } = Object.fromEntries(
      new URLSearchParams(window.location.search)
    );
    return <ForgotPassword initialEmail={email} initialCode={code} />;
  },
});
