import { Link } from "react-router-dom";
import Button from "../ui/Button";
import Input from "../ui/Input";

export default function Header() {
	return (
		<header className="bg-[var(--bg-base-tp)] h-18 z-10 sticky top-0 w-full border-b border-[var(--border)] backdrop-blur-sm">
			<div className="p-4 flex items-center justify-between h-full max-w-1600 mx-auto">
				<Link
					to="/"
					className="text-[var(--text-primary)] text-2xl font-display"
				>
					<span className="text-[var(--accent)]">Quest</span>
					<span>Board</span>
				</Link>
				<div className="flex items-center space-x-4">
					<Link
						to="#"
						className="text-[var(--text-secondary)] hover:text-[var(--accent)]"
					>
						Сессии
					</Link>
					<Link
						to="#"
						className="text-[var(--text-secondary)] hover:text-[var(--accent)]"
					>
						Мастера
					</Link>
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
