import { createLazyFileRoute } from "@tanstack/react-router";
import AuthForm from "../components/Auth/Authform";

export const Route = createLazyFileRoute("/login")({
  component: LoginComponent,
});

function LoginComponent() {
  return <AuthForm mode="login" />;
}
