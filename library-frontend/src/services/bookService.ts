// src/services/bookService.ts
import api from "./api";
import {
  type Book,
  type CreateBookRequest,
  type UpdateBookRequest,
  type BookFilter,
  type BooksResponse,
  type ApiResponse,
} from "./types";

export const bookService = {
  // Get all books with filters
  async getBooks(filters: BookFilter = {}): Promise<BooksResponse> {
    const params = new URLSearchParams();

    Object.entries(filters).forEach(([key, value]) => {
      if (value !== undefined && value !== null && value !== "") {
        params.append(key, value.toString());
      }
    });

    const response = await api.get<BooksResponse>(`/books?${params}`);
    return response.data;
  },

  // Get book by ID
  async getBookById(id: number): Promise<ApiResponse<Book>> {
    const response = await api.get<ApiResponse<Book>>(`/books/${id}`);
    return response.data;
  },

  // Create new book
  async createBook(book: CreateBookRequest): Promise<ApiResponse<Book>> {
    const response = await api.post<ApiResponse<Book>>("/books", book);
    return response.data;
  },

  // Update book
  async updateBook(
    id: number,
    book: UpdateBookRequest
  ): Promise<ApiResponse<Book>> {
    const response = await api.put<ApiResponse<Book>>(`/books/${id}`, book);
    return response.data;
  },

  // Delete book
  async deleteBook(id: number): Promise<void> {
    await api.delete(`/books/${id}`);
  },

  // Search books
  async searchBooks(query: string): Promise<BooksResponse> {
    const response = await api.get<BooksResponse>(
      `/books/search?q=${encodeURIComponent(query)}`
    );
    return response.data;
  },
};
