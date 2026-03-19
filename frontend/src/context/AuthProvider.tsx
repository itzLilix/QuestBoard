import { createContext, useState, ReactNode } from "react";

const AuthContext = createContext({});

export const AuthProvider = ({ children }: { children: ReactNode }) => {
	const [user, setUser] = useState(null);

	return (
		<AuthContext.Provider value={{ user, setUser }}>
			{children}
		</AuthContext.Provider>
	);
};

export default AuthContext;
