// src/components/book/BookCard.tsx
import { type Book } from "../../services/types";
import { Calendar, User, BookOpen, Edit, Trash2, Eye } from "lucide-react";
import { Button } from "../common/Button";

interface BookCardProps {
  book: Book;
  onEdit: (book: Book) => void;
  onDelete: (id: number) => void;
  onView: (id: number) => void;
}

export function BookCard({ book, onEdit, onDelete, onView }: BookCardProps) {
  return (
    <div className="card hover:shadow-lg transition-shadow duration-200">
      <div className="flex justify-between items-start mb-4">
        <div className="flex-1">
          <h3 className="text-lg font-semibold text-gray-900 mb-2 line-clamp-2">
            {book.title}
          </h3>
          <div className="space-y-1 text-sm text-gray-600">
            <div className="flex items-center">
              <User size={14} className="mr-2" />
              <span>{book.author}</span>
            </div>
            <div className="flex items-center">
              <Calendar size={14} className="mr-2" />
              <span>{book.year}</span>
            </div>
            {book.isbn && (
              <div className="flex items-center">
                <BookOpen size={14} className="mr-2" />
                <span className="font-mono text-xs">{book.isbn}</span>
              </div>
            )}
          </div>
        </div>
      </div>

      {book.description && (
        <p className="text-sm text-gray-600 mb-4 line-clamp-3">
          {book.description}
        </p>
      )}

      <div className="flex space-x-2">
        <Button
          variant="outline"
          size="sm"
          onClick={() => onView(book.id)}
          className="flex-1"
        >
          <Eye size={14} className="mr-1" />
          View
        </Button>
        <Button variant="secondary" size="sm" onClick={() => onEdit(book)}>
          <Edit size={14} className="mr-1" />
          Edit
        </Button>
        <Button variant="danger" size="sm" onClick={() => onDelete(book.id)}>
          <Trash2 size={14} className="mr-1" />
          Delete
        </Button>
      </div>
    </div>
  );
}
