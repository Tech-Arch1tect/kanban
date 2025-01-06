import { createFileRoute } from "@tanstack/react-router";

import ChangePassword from "../../components/profile/changePassword";
import { useUserProfile } from "../../hooks/profile/useUserProfile";
import UpdateUsername from "../../components/profile/updateUsername";

export const Route = createFileRoute("/profile/profile")({
  component: () => <Profile />,
});

const Profile = () => {
  const user = useUserProfile();

  return (
    <div>
      <ChangePassword />
      {user.profile && <UpdateUsername user={user.profile} />}
    </div>
  );
};
