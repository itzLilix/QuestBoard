import { Link } from "react-router-dom";
import { useEffect, useState } from "react";
import useAuthModal from "../../hooks/useAuthModal";
import Input from "../ui/Input";
import Tab from "../ui/Tab";
import Button from "../ui/Button";
import { api } from "../../api/axios";
import { AxiosError } from "axios";
import useAuth from "../../hooks/useAuth";

const LOGIN_URL = "/auth/login";
const REGISTER_URL = "/auth/register";

const USER_REGEX = /^[A-Za-z][A-Za-z0-9-_]{3,32}$/;
const PWD_REGEX = /^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%]).{8,24}$/;

export default function AuthModal() {
	const { isOpen, modalType, openModal, closeModal } = useAuthModal();
	const { login } = useAuth();

	const [user, setUser] = useState("");
	const [pwd, setPwd] = useState("");
	const [email, setEmail] = useState("");

	const [errMsg, setErrMsg] = useState("");
	useEffect(() => {
		setErrMsg("");
	}, [user, pwd]);
	useEffect(() => {
		setUser("");
		setPwd("");
		setEmail("");
		setErrMsg("");
	}, [modalType]);

	if (!isOpen) return null;
	return (
		<>
			<div
				className="fixed inset-0 bg-black/20 z-40"
				onClick={closeModal}
			></div>
			<div className="fixed h-[720px] w-[600px] m-auto inset-0 bg-[var(--bg-base-tp)] backdrop-blur-lg flex items-stretch justify-start flex-col rounded-2xl border border-[var(--border)] p-12 gap-6 z-50">
				<div className="text-3xl font-display text-[var(--text-primary)] select-none mx-auto mt-6">
					<span className="text-[var(--accent)]">Quest</span>
					<span>Board</span>
				</div>
				<nav className="w-full flex justify-center gap-6 text-3xl">
					<Tab
						isActive={modalType === "login"}
						onClick={() => {
							openModal("login");
						}}
					>
						Вход
					</Tab>
					<Tab
						isActive={modalType === "register"}
						onClick={() => {
							openModal("register");
						}}
					>
						Регистрация
					</Tab>
				</nav>
				<span className="absolute top-[220px] left-1/2 transform -translate-x-1/2 text-sm text-[var(--error)] text-center">
					{errMsg}
				</span>
				{modalType === "login" ? (
					<div key="login">
						<form
							onSubmit={async (e) => {
								e.preventDefault();

								try {
									const response = await api.post(LOGIN_URL, {
										username: user,
										password: pwd,
									});
									console.log(response);
									login(response.data.user);
									closeModal();
								} catch (err) {
									const axiosErr = err as AxiosError<{
										message: string;
									}>;
									console.error(err);
									setErrMsg(
										axiosErr.response?.data?.message ||
											"Login failed",
									);
								}
							}}
							className="flex flex-col gap-4 mt-6"
						>
							<div className="flex flex-col gap-2">
								<span className="text-base text-[var(--text-primary)]">
									Имя пользователя
								</span>
								<Input
									csize="md"
									className="w-full"
									name="username"
									onChange={(e) => setUser(e.target.value)}
									required
								></Input>
							</div>
							<div className="flex flex-col gap-2">
								<span className="text-base text-[var(--text-primary)]">
									Пароль
								</span>
								<Input
									type="password"
									csize="md"
									className="w-full"
									name="password"
									onChange={(e) => setPwd(e.target.value)}
									required
								></Input>
								<Link
									to="/auth/reset-password"
									onClick={closeModal}
									className="text-sm text-[var(--text-secondary)] hover:text-[var(--accent)] self-center mt-2"
								>
									Сброс пароля
								</Link>
							</div>
							<Button
								variant={"primary"}
								csize={"md"}
								type="submit"
								className="mx-6 mt-6"
							>
								Войти
							</Button>
						</form>
						<div className="flex items-center gap-4 my-12 mx-24">
							<div className="flex-1 border-t border-[var(--accent)]"></div>
							<span className="text-base font-body text-[var(--accent)]">
								или
							</span>
							<div className="flex-1 border-t border-[var(--accent)]"></div>
						</div>
					</div>
				) : (
					<div key="register">
						<form
							onSubmit={async (e) => {
								e.preventDefault();

								const v1 = USER_REGEX.test(user);
								const v2 = PWD_REGEX.test(pwd);
								if (!v1) {
									setErrMsg(
										"Имя пользователя: 4-32 символа, буквы и цифры",
									);
									return;
								} else if (!v2) {
									setErrMsg(
										"Пароль: 8-24 символа, буквы, цифры и !@#$%",
									);
									return;
								}

								try {
									const response = await api.post(
										REGISTER_URL,
										{
											email: email,
											username: user,
											password: pwd,
										},
									);
									console.log(response);
									login(response.data.user);
									closeModal();
								} catch (err) {
									console.error(err);
									const axiosErr = err as AxiosError<{
										message: string;
									}>;
									setErrMsg(
										axiosErr.response?.data?.message ||
											"Registration failed",
									);
								}
							}}
							className="flex flex-col gap-4 mt-6"
						>
							<div className="flex flex-col gap-2">
								<span className="text-base text-[var(--text-primary)]">
									Электронная почта
								</span>
								<Input
									csize="md"
									className="w-full"
									type="email"
									name="email"
									onChange={(e) => setEmail(e.target.value)}
									required
								></Input>
							</div>
							<div className="flex flex-col gap-2">
								<span className="text-base text-[var(--text-primary)]">
									Имя пользователя
								</span>
								<Input
									csize="md"
									className="w-full"
									name="username"
									onChange={(e) => setUser(e.target.value)}
									required
								></Input>
							</div>
							<div className="flex flex-col gap-2">
								<span className="text-base text-[var(--text-primary)]">
									Пароль
								</span>
								<Input
									type="password"
									csize="md"
									className="w-full"
									name="password"
									onChange={(e) => setPwd(e.target.value)}
									required
								></Input>
							</div>
							<Button
								variant={"primary"}
								csize={"md"}
								type="submit"
								className="mx-6 mt-6"
							>
								Зарегистрироваться
							</Button>
						</form>
					</div>
				)}
			</div>
		</>
	);
}
