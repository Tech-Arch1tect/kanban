import { useState } from "react";
import { useUpdateNotificationConfiguration } from "../../../hooks/notifications/useUpdateNotificationConfiguration";
import { useBoards } from "../../../hooks/boards/useBoards";
import { useEventData } from "../../../hooks/notifications/useEventData";
import { friendlyEventNames } from "./friendlyEventNames";
import {
  ModelsBoard,
  ModelsNotificationConfiguration,
  ModelsNotificationEvent,
  NotificationUpdateNotificationRequest,
} from "../../../typescript-fetch-client";

interface NotificationEditFormProps {
  notification: ModelsNotificationConfiguration;
  onCancel: () => void;
  onUpdated: () => void;
}

export const NotificationEditForm = ({
  notification,
  onCancel,
  onUpdated,
}: NotificationEditFormProps) => {
  const [name, setName] = useState(notification.name);
  const [method, setMethod] = useState<"email" | "webhook">(
    notification.method as "email" | "webhook"
  );
  const [email, setEmail] = useState(notification.email || "");
  const [webhookUrl, setWebhookUrl] = useState(notification.webhookUrl || "");
  const [selectedBoards, setSelectedBoards] = useState<number[]>(
    notification.boards
      ? notification.boards.map((b: ModelsBoard) => b.id as number)
      : []
  );
  const [selectedEvents, setSelectedEvents] = useState<string[]>(
    notification.events
      ? notification.events.map((e: ModelsNotificationEvent) => {
          if (e.name) {
            return e.name;
          }
          return "";
        })
      : []
  );
  const [onlyAssignee, setOnlyAssignee] = useState(notification.onlyAssignee);

  const { boards, isLoading: boardsLoading, error: boardsError } = useBoards();
  const {
    data: eventsData,
    isLoading: eventsLoading,
    error: eventsError,
  } = useEventData();

  const { mutate: updateNotification, isPending } =
    useUpdateNotificationConfiguration();

  const handleBoardChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const boardId = parseInt(e.target.value, 10);
    if (e.target.checked) {
      setSelectedBoards((prev) => [...prev, boardId]);
    } else {
      setSelectedBoards((prev) => prev.filter((id) => id !== boardId));
    }
  };

  const handleCheckboxChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const eventValue = e.target.value;
    if (e.target.checked) {
      setSelectedEvents((prev) => [...prev, eventValue]);
    } else {
      setSelectedEvents((prev) => prev.filter((item) => item !== eventValue));
    }
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    const payload = {
      id: notification.id,
      name,
      boards: selectedBoards,
      email: method === "email" ? email : undefined,
      events: selectedEvents,
      method,
      webhookUrl: method === "webhook" ? webhookUrl : undefined,
      onlyAssignee,
    };

    updateNotification(payload as NotificationUpdateNotificationRequest, {
      onSuccess: () => {
        onUpdated();
      },
    });
  };

  return (
    <div className="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-sm dark:shadow-md mb-8">
      <form onSubmit={handleSubmit} className="space-y-4">
        {/* Name */}
        <div>
          <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
            Name:
          </label>
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
            className="mt-1 w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
          />
        </div>

        {/* Boards */}
        <div>
          <span className="block text-sm font-medium text-gray-700 dark:text-gray-300">
            Boards:
          </span>
          {boardsLoading && (
            <p className="text-gray-600 dark:text-gray-400">Loading boards…</p>
          )}
          {boardsError && (
            <p className="text-red-600 dark:text-red-400">
              Error loading boards.
            </p>
          )}
          {boards &&
            boards.boards?.map((board: ModelsBoard) => (
              <label
                key={board.id}
                className="inline-flex items-center mr-4 text-gray-700 dark:text-gray-300"
              >
                <input
                  type="checkbox"
                  value={board.id}
                  checked={selectedBoards.includes(board.id as number)}
                  onChange={handleBoardChange}
                  className="form-checkbox text-blue-600"
                />
                <span className="ml-2">{board.name}</span>
              </label>
            ))}
        </div>

        {/* Method */}
        <div>
          <span className="block text-sm font-medium text-gray-700 dark:text-gray-300">
            Method:
          </span>
          <div className="mt-1 flex items-center space-x-4">
            <label className="inline-flex items-center text-gray-700 dark:text-gray-300">
              <input
                type="radio"
                name="method"
                value="email"
                checked={method === "email"}
                onChange={() => setMethod("email")}
                className="form-radio text-blue-600"
              />
              <span className="ml-2">Email</span>
            </label>
            <label className="inline-flex items-center text-gray-700 dark:text-gray-300">
              <input
                type="radio"
                name="method"
                value="webhook"
                checked={method === "webhook"}
                onChange={() => setMethod("webhook")}
                className="form-radio text-blue-600"
              />
              <span className="ml-2">Webhook</span>
            </label>
          </div>
        </div>

        {/* Email or Webhook URL */}
        {method === "email" && (
          <div>
            <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
              Email:
            </label>
            <input
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
              className="mt-1 w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
            />
          </div>
        )}
        {method === "webhook" && (
          <div>
            <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
              Webhook URL:
            </label>
            <input
              type="url"
              value={webhookUrl}
              onChange={(e) => setWebhookUrl(e.target.value)}
              required
              className="mt-1 w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
            />
          </div>
        )}

        {/* Events */}
        <div>
          <fieldset>
            <legend className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
              Select Events:
            </legend>
            {eventsLoading && (
              <p className="text-gray-600 dark:text-gray-400">
                Loading events…
              </p>
            )}
            {eventsError && (
              <p className="text-red-600 dark:text-red-400">
                Error loading events.
              </p>
            )}
            {eventsData &&
              eventsData.events &&
              eventsData.events.map((eventType: string) => (
                <label
                  key={eventType}
                  className="inline-flex items-center mr-4 text-gray-700 dark:text-gray-300"
                >
                  <input
                    type="checkbox"
                    value={eventType}
                    checked={selectedEvents.includes(eventType)}
                    onChange={handleCheckboxChange}
                    className="form-checkbox text-blue-600"
                  />
                  <span className="ml-2">
                    {friendlyEventNames[eventType] || eventType}
                  </span>
                </label>
              ))}
          </fieldset>
        </div>

        {/* Only Assignee */}
        <div>
          <label className="inline-flex items-center text-gray-700 dark:text-gray-300">
            <input
              type="checkbox"
              checked={onlyAssignee}
              onChange={(e) => setOnlyAssignee(e.target.checked)}
              className="form-checkbox text-blue-600"
            />
            <span className="ml-2">
              Only notify if I am the assignee of the task
            </span>
          </label>
        </div>

        {/* Buttons */}
        <div className="flex space-x-4">
          <button
            type="submit"
            disabled={isPending}
            className={`w-full px-4 py-2 text-white rounded-md transition-colours ${
              isPending
                ? "bg-gray-400 dark:bg-gray-600 cursor-not-allowed"
                : "bg-blue-600 dark:bg-blue-700 hover:bg-blue-700 dark:hover:bg-blue-800"
            }`}
          >
            {isPending ? "Saving…" : "Save Changes"}
          </button>
          <button
            type="button"
            onClick={onCancel}
            className="px-4 py-2 bg-gray-200 dark:bg-gray-700 rounded-md text-gray-700 dark:text-gray-300"
          >
            Cancel
          </button>
        </div>
      </form>
    </div>
  );
};
