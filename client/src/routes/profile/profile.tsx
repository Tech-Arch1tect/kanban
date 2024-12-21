import { createFileRoute } from "@tanstack/react-router";

import ChangePassword from "../../components/profile/changePassword";
import UpdateDisplayName from "../../components/profile/updateDisplayName";
import { useUserProfile } from "../../hooks/profile/useUserProfile";

export const Route = createFileRoute("/profile/profile")({
  component: () => <Profile />,
});

const Profile = () => {
  const user = useUserProfile();

  return (
    <div>
      <ChangePassword />
      {user.profile && <UpdateDisplayName user={user.profile} />}
    </div>
  );
};
