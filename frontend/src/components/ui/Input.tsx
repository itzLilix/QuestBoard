import { FC } from "react";

type InputProps = React.InputHTMLAttributes<HTMLInputElement> & {
	csize?: "sm" | "md";
};

const Input: FC<InputProps> = ({ csize = "md", className, ...props }) => {
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
};

export default Input;
