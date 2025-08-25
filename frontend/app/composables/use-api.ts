export function useApi<T>(path: string, options: any = {}) {
  const { token } = useAuth();

  return useFetch<T>(path, {
    ...options,
    headers: {
      ...(options.headers || {}),
      Authorization: token.value ? `Bearer ${token.value}` : "",
    },
  });
}
