import { createFileRoute } from "@tanstack/react-router";
import ForgotPassword from "../components/ForgotPassword";

export const Route = createFileRoute("/password-reset")({
  component: () => <ForgotPassword />,
});
