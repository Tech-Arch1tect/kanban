import { createLazyFileRoute } from "@tanstack/react-router";
import ForgotPassword from "../components/Auth/ForgotPassword";

export const Route = createLazyFileRoute("/password-reset")({
  component: () => {
    const { email, code } = Object.fromEntries(
      new URLSearchParams(window.location.search)
    );
    return <ForgotPassword initialEmail={email} initialCode={code} />;
  },
});
