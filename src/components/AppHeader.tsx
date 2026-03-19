import Button from "./ui/Button";
import Input from "./ui/Input";

export default function Header() {
	return (
		<header className="bg-[var(--bg-base-tp)] h-18 z-10 sticky top-0 w-full border-b border-[var(--border)] backdrop-blur-sm">
			<div className="p-4 flex items-center justify-between h-full max-w-1600 mx-auto">
				<a
					className="text-[var(--text-primary)] text-2xl font-display"
					href="#"
				>
					<span className="text-[var(--accent)]">Quest</span>
					<span>Board</span>
				</a>
				<div className="flex items-center space-x-4">
					<a
						href="#"
						className="text-[var(--text-secondary)] hover:text-[var(--accent)]"
					>
						Сессии
					</a>
					<a
						href="#"
						className="text-[var(--text-secondary)] hover:text-[var(--accent)]"
					>
						Мастера
					</a>
				</div>
				<Input placeholder={"Поиск..."} csize={"sm"} />
				<Button onClick={() => {}} variant="secondary" csize="sm">
					+ Создать сессию
				</Button>
				<div></div>
			</div>
		</header>
	);
}
