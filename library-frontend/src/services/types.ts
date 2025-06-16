// src/services/types.ts
export interface Book {
  id: number;
  title: string;
  author: string;
  year: number;
  isbn?: string;
  description?: string;
  created_at: string;
  updated_at: string;
}

export interface CreateBookRequest {
  title: string;
  author: string;
  year: number;
  isbn?: string;
  description?: string;
}

export interface UpdateBookRequest {
  title?: string;
  author?: string;
  year?: number;
  isbn?: string;
  description?: string;
}

export interface BookFilter {
  title?: string;
  author?: string;
  year?: number;
  limit?: number;
  offset?: number;
}

export interface ApiResponse<T> {
  success: boolean;
  data: T;
  total?: number;
  message?: string;
}

export interface ApiError {
  success: false;
  error: string;
  code?: string;
  details?: string;
}

export interface BooksResponse {
  success: boolean;
  data: Book[];
  total: number;
  page?: number;
  limit?: number;
}
