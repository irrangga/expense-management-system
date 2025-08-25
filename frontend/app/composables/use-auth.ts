import type { LoginResponse } from "~~/shared/types/response";

export const useAuth = () => {
  const token = useState<string | null>("auth_token", () => null);
  const user = useState<User | null>("auth_user", () => null);

  const login = async (email: string) => {
    try {
      const { data, error } = await useFetch<LoginResponse>("/api/auth/login", {
        method: "POST",
        body: { email },
      });

      if (error.value) throw error.value;

      token.value = data.value?.token || null;
      user.value = data.value?.user || null;

      if (process.client && token.value) {
        localStorage.setItem("auth_token", token.value);
        localStorage.setItem("auth_user", JSON.stringify(user.value));
      }
    } catch (err) {
      console.error("Login failed:", err);
    }
  };

  return { token, user, login };
};
