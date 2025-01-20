import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";

export default function RenderMarkdown({ markdown }: { markdown: string }) {
  return (
    <div className="prose dark:prose-invert min-w-full">
      <ReactMarkdown remarkPlugins={[remarkGfm]}>{markdown}</ReactMarkdown>
    </div>
  );
}
