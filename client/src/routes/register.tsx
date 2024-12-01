import { createFileRoute } from "@tanstack/react-router";
import AuthForm from "../components/Authform";

export const Route = createFileRoute("/register")({
  component: RegisterComponent,
});

function RegisterComponent() {
  return <AuthForm mode="register" />;
}
