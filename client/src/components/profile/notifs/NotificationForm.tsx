import { useState } from "react";
import { useEventData } from "../../../hooks/notifications/useEventData";
import { useCreateNotificationConfiguration } from "../../../hooks/notifications/useCreateNotificationConfiguration";
import { useBoards } from "../../../hooks/boards/useBoards";

export const NotificationForm = () => {
  const [name, setName] = useState("");
  const [method, setMethod] = useState<"email" | "webhook">("email");
  const [email, setEmail] = useState("");
  const [webhookUrl, setWebhookUrl] = useState("");
  const [selectedBoards, setSelectedBoards] = useState<number[]>([]);
  const [selectedEvents, setSelectedEvents] = useState<string[]>([]);
  const [onlyAssignee, setOnlyAssignee] = useState(false);

  const {
    data: eventsData,
    isLoading: eventsLoading,
    error: eventsError,
  } = useEventData();

  const { boards, isLoading: boardsLoading, error: boardsError } = useBoards();

  const { mutate: createNotification, isPending } =
    useCreateNotificationConfiguration();

  const handleCheckboxChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const eventValue = e.target.value;
    if (e.target.checked) {
      setSelectedEvents((prev) => [...prev, eventValue]);
    } else {
      setSelectedEvents((prev) => prev.filter((item) => item !== eventValue));
    }
  };

  const handleBoardChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const boardId = parseInt(e.target.value, 10);
    if (e.target.checked) {
      setSelectedBoards((prev) => [...prev, boardId]);
    } else {
      setSelectedBoards((prev) => prev.filter((id) => id !== boardId));
    }
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    const payload = {
      name,
      boards: selectedBoards,
      email: method === "email" ? email : undefined,
      events: selectedEvents,
      method,
      webhookUrl: method === "webhook" ? webhookUrl : undefined,
      onlyAssignee,
    };

    createNotification(payload);
  };

  const friendlyEventNames: Record<string, string> = {
    "task.created": "Task created",
    "task.updated.title": "Task title updated",
    "task.updated.description": "Task description updated",
    "task.updated.status": "Task status updated",
    "task.updated.assignee": "Task assignee updated",
    "task.deleted": "Task deleted",
    "task.moved": "Task moved",
    "comment.created": "Comment created",
    "comment.updated": "Comment updated",
    "comment.deleted": "Comment deleted",
    "file.created": "File created",
    "file.updated": "File updated",
    "file.deleted": "File deleted",
    "link.created": "Link created",
    "link.deleted": "Link deleted",
    "externallink.created": "External link created",
    "externallink.updated": "External link updated",
    "externallink.deleted": "External link deleted",
  };

  return (
    <div className="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-sm dark:shadow-md mb-8">
      <form onSubmit={handleSubmit} className="space-y-4">
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
            boards.boards?.map((board) => (
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

        <button
          type="submit"
          disabled={isPending}
          className={`w-full px-4 py-2 text-white rounded-md transition-colours ${
            isPending
              ? "bg-gray-400 dark:bg-gray-600 cursor-not-allowed"
              : "bg-blue-600 dark:bg-blue-700 hover:bg-blue-700 dark:hover:bg-blue-800"
          }`}
        >
          {isPending ? "Creating…" : "Create Notification"}
        </button>
      </form>
    </div>
  );
};
