import type { User } from "./user";

export interface LoginResponse {
  token: string;
  user: User;
}

export interface PaginatedResponse<T> {
  page: number;
  page_size: number;
  total: number;
  total_pages: number;
  data: T;
}
