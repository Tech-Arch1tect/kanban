import { createFileRoute } from "@tanstack/react-router";

import ChangePassword from "../../components/profile/changePassword";

export const Route = createFileRoute("/profile/profile")({
  component: () => <Profile />,
});

const Profile = () => {
  return (
    <div>
      <ChangePassword />
    </div>
  );
};
