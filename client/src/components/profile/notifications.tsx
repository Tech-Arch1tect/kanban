import { NotificationForm } from "./notifs/NotificationForm";
import { NotificationTable } from "./notifs/NotificationTable";

export default function Notifications() {
  return (
    <div className="p-6 bg-gray-50 dark:bg-gray-900 min-h-screen">
      <h1 className="text-3xl font-bold text-gray-800 dark:text-gray-100 mb-6">
        Notification Configurations
      </h1>
      <NotificationForm />
      <NotificationTable />
    </div>
  );
}
