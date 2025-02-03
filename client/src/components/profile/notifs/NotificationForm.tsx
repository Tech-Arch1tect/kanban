import React, { useState } from "react";
import { useEventData } from "../../../hooks/notifications/useEventData";
import { useCreateNotificationConfiguration } from "../../../hooks/notifications/useCreateNotificationConfiguration";

export const NotificationForm = () => {
  const [name, setName] = useState("");
  const [method, setMethod] = useState<"email" | "webhook">("email");
  const [email, setEmail] = useState("");
  const [webhookUrl, setWebhookUrl] = useState("");
  const [boards, setBoards] = useState("");
  const [selectedEvents, setSelectedEvents] = useState<string[]>([]);

  const {
    data: eventsData,
    isLoading: eventsLoading,
    error: eventsError,
  } = useEventData();
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

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    const boardsArray = boards
      .split(",")
      .map((b) => parseInt(b.trim(), 10))
      .filter((n) => !isNaN(n));

    const payload = {
      name,
      boards: boardsArray,
      email: method === "email" ? email : undefined,
      events: selectedEvents,
      method,
      webhookUrl: method === "webhook" ? webhookUrl : undefined,
    };

    createNotification(payload);
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
          <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
            Boards (comma separated IDs):
          </label>
          <input
            type="text"
            value={boards}
            onChange={(e) => setBoards(e.target.value)}
            required
            className="mt-1 w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
          />
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
              eventsData.events.map((eventType) => (
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
                  <span className="ml-2">{eventType}</span>
                </label>
              ))}
          </fieldset>
        </div>

        <button
          type="submit"
          disabled={isPending}
          className={`w-full px-4 py-2 text-white rounded-md transition-colors ${
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
