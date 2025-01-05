import ReactMarkdown from "react-markdown";

export default function RenderMarkdown({
  markdown,
  className,
}: {
  markdown: string;
  className?: string;
}) {
  return <ReactMarkdown className={className}>{markdown}</ReactMarkdown>;
}
