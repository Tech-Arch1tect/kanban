import { createFileRoute } from "@tanstack/react-router";
import AuthForm from "../components/Auth/Authform";

export const Route = createFileRoute("/login")({
  component: LoginComponent,
});

function LoginComponent() {
  return <AuthForm mode="login" />;
}
