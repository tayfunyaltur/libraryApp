import {
  createContext,
  useContext,
  useReducer,
  useEffect,
  type ReactNode,
} from "react";
import {
  type Book,
  type BookFilter,
  type CreateBookRequest,
  type UpdateBookRequest,
} from "../services/types";
import { bookService } from "../services/bookService";
import toast from "react-hot-toast";

interface BookState {
  books: Book[];
  currentBook: Book | null;
  loading: boolean;
  error: string | null;
  filters: BookFilter;
  total: number;
  page: number;
}

type BookAction =
  | { type: "SET_LOADING"; payload: boolean }
  | { type: "SET_BOOKS"; payload: { books: Book[]; total: number } }
  | { type: "SET_CURRENT_BOOK"; payload: Book | null }
  | { type: "ADD_BOOK"; payload: Book }
  | { type: "UPDATE_BOOK"; payload: Book }
  | { type: "DELETE_BOOK"; payload: number }
  | { type: "SET_ERROR"; payload: string | null }
  | { type: "SET_FILTERS"; payload: BookFilter }
  | { type: "SET_PAGE"; payload: number };

const initialState: BookState = {
  books: [],
  currentBook: null,
  loading: false,
  error: null,
  filters: { limit: 10, offset: 0 },
  total: 0,
  page: 1,
};

function bookReducer(state: BookState, action: BookAction): BookState {
  switch (action.type) {
    case "SET_LOADING":
      return { ...state, loading: action.payload };
    case "SET_BOOKS":
      return {
        ...state,
        books: action.payload.books,
        total: action.payload.total,
        loading: false,
        error: null,
      };
    case "SET_CURRENT_BOOK":
      return { ...state, currentBook: action.payload };
    case "ADD_BOOK":
      return {
        ...state,
        books: [action.payload, ...state.books],
        total: state.total + 1,
      };
    case "UPDATE_BOOK":
      return {
        ...state,
        books: state.books.map((book) =>
          book.id === action.payload.id ? action.payload : book
        ),
        currentBook:
          state.currentBook?.id === action.payload.id
            ? action.payload
            : state.currentBook,
      };
    case "DELETE_BOOK":
      return {
        ...state,
        books: state.books.filter((book) => book.id !== action.payload),
        total: state.total - 1,
        currentBook:
          state.currentBook?.id === action.payload ? null : state.currentBook,
      };
    case "SET_ERROR":
      return { ...state, error: action.payload, loading: false };
    case "SET_FILTERS":
      return { ...state, filters: action.payload };
    case "SET_PAGE":
      return { ...state, page: action.payload };
    default:
      return state;
  }
}

interface BookContextType extends BookState {
  fetchBooks: () => Promise<void>;
  fetchBookById: (id: number) => Promise<void>;
  createBook: (book: CreateBookRequest) => Promise<boolean>;
  updateBook: (id: number, book: UpdateBookRequest) => Promise<boolean>;
  deleteBook: (id: number) => Promise<boolean>;
  searchBooks: (query: string) => Promise<void>;
  setFilters: (filters: BookFilter) => void;
  setPage: (page: number) => void;
  clearCurrentBook: () => void;
}

const BookContext = createContext<BookContextType | undefined>(undefined);

export function useBooks() {
  const context = useContext(BookContext);
  if (context === undefined) {
    throw new Error("useBooks must be used within a BookProvider");
  }
  return context;
}

interface BookProviderProps {
  children: ReactNode;
}

export function BookProvider({ children }: BookProviderProps) {
  const [state, dispatch] = useReducer(bookReducer, initialState);

  const fetchBooks = async () => {
    try {
      dispatch({ type: "SET_LOADING", payload: true });
      const response = await bookService.getBooks(state.filters);
      dispatch({
        type: "SET_BOOKS",
        payload: { books: response.data, total: response.total },
      });
    } catch (error) {
      dispatch({ type: "SET_ERROR", payload: "Failed to fetch books" });
    }
  };

  const fetchBookById = async (id: number) => {
    try {
      dispatch({ type: "SET_LOADING", payload: true });
      const response = await bookService.getBookById(id);
      dispatch({ type: "SET_CURRENT_BOOK", payload: response.data });
      dispatch({ type: "SET_LOADING", payload: false });
    } catch (error) {
      dispatch({ type: "SET_ERROR", payload: "Failed to fetch book" });
    }
  };

  const createBook = async (book: CreateBookRequest): Promise<boolean> => {
    try {
      const response = await bookService.createBook(book);
      dispatch({ type: "ADD_BOOK", payload: response.data });
      toast.success("Book created successfully");
      return true;
    } catch (error) {
      return false;
    }
  };

  const updateBook = async (
    id: number,
    book: UpdateBookRequest
  ): Promise<boolean> => {
    try {
      const response = await bookService.updateBook(id, book);
      dispatch({ type: "UPDATE_BOOK", payload: response.data });
      toast.success("Book updated successfully");
      return true;
    } catch (error) {
      return false;
    }
  };

  const deleteBook = async (id: number): Promise<boolean> => {
    try {
      await bookService.deleteBook(id);
      dispatch({ type: "DELETE_BOOK", payload: id });
      toast.success("Book deleted successfully");
      return true;
    } catch (error) {
      return false;
    }
  };

  const searchBooks = async (query: string) => {
    try {
      dispatch({ type: "SET_LOADING", payload: true });
      const response = await bookService.searchBooks(query);
      dispatch({
        type: "SET_BOOKS",
        payload: { books: response.data, total: response.total },
      });
    } catch (error) {
      dispatch({ type: "SET_ERROR", payload: "Search failed" });
    }
  };

  const setFilters = (filters: BookFilter) => {
    dispatch({ type: "SET_FILTERS", payload: filters });
  };

  const setPage = (page: number) => {
    dispatch({ type: "SET_PAGE", payload: page });
    const newOffset = (page - 1) * (state.filters.limit || 10);
    setFilters({ ...state.filters, offset: newOffset });
  };

  const clearCurrentBook = () => {
    dispatch({ type: "SET_CURRENT_BOOK", payload: null });
  };

  // Auto-fetch books when filters change
  useEffect(() => {
    fetchBooks();
  }, [state.filters]);

  const contextValue: BookContextType = {
    ...state,
    fetchBooks,
    fetchBookById,
    createBook,
    updateBook,
    deleteBook,
    searchBooks,
    setFilters,
    setPage,
    clearCurrentBook,
  };

  return (
    <BookContext.Provider value={contextValue}>{children}</BookContext.Provider>
  );
}
