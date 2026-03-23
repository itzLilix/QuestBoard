import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import App from "./App";
import { BrowserRouter } from "react-router-dom";
import { AuthProvider } from "./context/AuthContext";
import { AuthModalProvider } from "./context/AuthModalContext";

const root = ReactDOM.createRoot(
	document.getElementById("root") as HTMLElement,
);
root.render(
	<React.StrictMode>
		<BrowserRouter>
			<AuthProvider>
				<AuthModalProvider>
					<App />
				</AuthModalProvider>
			</AuthProvider>
		</BrowserRouter>
	</React.StrictMode>,
);
