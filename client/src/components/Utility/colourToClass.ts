const colourToClass: { [key: string]: string } = {
  slate: "bg-slate-500 dark:bg-slate-700 text-white",
  gray: "bg-gray-500 dark:bg-gray-700 text-white",
  zinc: "bg-zinc-500 dark:bg-zinc-700 text-white",
  neutral: "bg-neutral-500 dark:bg-neutral-700 text-white",
  stone: "bg-stone-500 dark:bg-stone-700 text-white",
  red: "bg-red-500 dark:bg-red-700 text-white",
  orange: "bg-orange-500 dark:bg-orange-700 text-white",
  amber: "bg-amber-500 dark:bg-amber-700 text-white",
  yellow: "bg-yellow-500 dark:bg-yellow-700 text-white",
  lime: "bg-lime-500 dark:bg-lime-700 text-white",
  green: "bg-green-500 dark:bg-green-700 text-white",
  emerald: "bg-emerald-500 dark:bg-emerald-700 text-white",
  teal: "bg-teal-500 dark:bg-teal-700 text-white",
  cyan: "bg-cyan-500 dark:bg-cyan-700 text-white",
  sky: "bg-sky-500 dark:bg-sky-700 text-white",
  blue: "bg-blue-500 dark:bg-blue-700 text-white",
  indigo: "bg-indigo-500 dark:bg-indigo-700 text-white",
  violet: "bg-violet-500 dark:bg-violet-700 text-white",
  purple: "bg-purple-500 dark:bg-purple-700 text-white",
  fuchsia: "bg-fuchsia-500 dark:bg-fuchsia-700 text-white",
  pink: "bg-pink-500 dark:bg-pink-700 text-white",
  rose: "bg-rose-500 dark:bg-rose-700 text-white",
};

const colourToBorderClass: { [key: string]: string } = {
  slate: "border-slate-500 dark:border-slate-700",
  gray: "border-gray-500 dark:border-gray-700",
  zinc: "border-zinc-500 dark:border-zinc-700",
  neutral: "border-neutral-500 dark:border-neutral-700",
  stone: "border-stone-500 dark:border-stone-700",
  red: "border-red-500 dark:border-red-700",
  orange: "border-orange-500 dark:border-orange-700",
  amber: "border-amber-500 dark:border-amber-700",
  yellow: "border-yellow-500 dark:border-yellow-700",
  lime: "border-lime-500 dark:border-lime-700",
  green: "border-green-500 dark:border-green-700",
  emerald: "border-emerald-500 dark:border-emerald-700",
  teal: "border-teal-500 dark:border-teal-700",
  cyan: "border-cyan-500 dark:border-cyan-700",
  sky: "border-sky-500 dark:border-sky-700",
  blue: "border-blue-500 dark:border-blue-700",
  indigo: "border-indigo-500 dark:border-indigo-700",
  violet: "border-violet-500 dark:border-violet-700",
  purple: "border-purple-500 dark:border-purple-700",
  fuchsia: "border-fuchsia-500 dark:border-fuchsia-700",
  pink: "border-pink-500 dark:border-pink-700",
  rose: "border-rose-500 dark:border-rose-700",
};

export { colourToClass, colourToBorderClass };
