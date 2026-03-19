type InputProps = React.InputHTMLAttributes<HTMLInputElement> & {
	csize?: "sm" | "md";
};

export default function Input({
	csize = "md",
	className,
	...props
}: InputProps) {
	return (
		<input
			className={`
				bg-[var(--bg-surface)]
				text-[var(--text-primary)]
				border border-[var(--border)]
				rounded-xl 
                w-full ${csize === "sm" ? "max-w-80" : ""}
                px-3 ${csize === "sm" ? "py-2" : "py-3"}
				focus:outline-none focus:ring-2 focus:ring-[var(--accent)]
				focus:ring-offset-2 focus:ring-offset-[var(--bg-base)]
				transition-colors duration-200
				placeholder:text-[var(--text-muted)]
				${className || ""}
			`}
			{...props}
		/>
	);
}
