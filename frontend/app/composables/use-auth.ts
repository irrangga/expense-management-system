import type { LoginResponse } from "~~/shared/types/response";

export const useAuth = () => {
  const token = useCookie<string | null>("auth_token", {
    sameSite: "lax",
    watch: true,
  });

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
    } catch (err) {
      throw err;
    }
  };

  const logout = () => {
    token.value = null;
    user.value = null;
  };

  return { token, user, login, logout };
};
