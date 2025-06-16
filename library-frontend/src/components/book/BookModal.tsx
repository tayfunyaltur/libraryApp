// src/components/book/BookModal.tsx
import { Modal } from "../common/Modal";
import { BookForm } from "./BookForm";
import { type Book, type CreateBookRequest } from "../../services/types";
import { useBooks } from "../../context/BookContext";

interface BookModalProps {
  isOpen: boolean;
  onClose: () => void;
  book?: Book | null;
}

export function BookModal({ isOpen, onClose, book }: BookModalProps) {
  const { createBook, updateBook, loading } = useBooks();

  const handleSubmit = async (data: CreateBookRequest): Promise<boolean> => {
    if (book) {
      return await updateBook(book.id, data);
    } else {
      return await createBook(data);
    }
  };

  return (
    <Modal
      isOpen={isOpen}
      onClose={onClose}
      title={book ? "Edit Book" : "Add New Book"}
      size="lg"
    >
      <BookForm
        book={book || undefined}
        onSubmit={handleSubmit}
        onCancel={onClose}
        loading={loading}
      />
    </Modal>
  );
}
