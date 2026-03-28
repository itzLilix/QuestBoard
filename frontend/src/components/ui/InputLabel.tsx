export default function InputLabel({
	children,
}: {
	children: React.ReactNode;
}) {
	return (
		<span className="text-base text-[var(--text-primary)]">{children}</span>
	);
}
