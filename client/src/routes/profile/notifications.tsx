import { createFileRoute } from "@tanstack/react-router";
import Notifications from "../../components/profile/notifications";

export const Route = createFileRoute("/profile/notifications")({
  component: Notifications,
});
