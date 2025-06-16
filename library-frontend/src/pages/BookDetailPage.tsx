// src/pages/BookDetailPage.tsx
import { useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { useBooks } from "../context/BookContext";
import { Loading } from "../components/common/Loading";
import { Button } from "../components/common/Button";
import {
  ArrowLeft,
  Calendar,
  User,
  BookOpen,
  Edit,
  Trash2,
} from "lucide-react";

export function BookDetailPage() {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const { currentBook, loading, fetchBookById, deleteBook } = useBooks();

  useEffect(() => {
    if (id) {
      fetchBookById(parseInt(id));
    }
  }, [id]);

  const handleDelete = async () => {
    if (
      currentBook &&
      window.confirm("Are you sure you want to delete this book?")
    ) {
      const success = await deleteBook(currentBook.id);
      if (success) {
        navigate("/");
      }
    }
  };

  if (loading) {
    return <Loading text="Loading book details..." />;
  }

  if (!currentBook) {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center">
        <div className="text-center">
          <h2 className="text-2xl font-bold text-gray-900 mb-4">
            Book not found
          </h2>
          <Button onClick={() => navigate("/")}>
            <ArrowLeft size={16} className="mr-2" />
            Back to Dashboard
          </Button>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="mb-6">
          <Button variant="outline" onClick={() => navigate("/")}>
            <ArrowLeft size={16} className="mr-2" />
            Back to Dashboard
          </Button>
        </div>

        <div className="card">
          <div className="flex justify-between items-start mb-6">
            <div className="flex-1">
              <h1 className="text-3xl font-bold text-gray-900 mb-4">
                {currentBook.title}
              </h1>
              <div className="space-y-2 text-lg text-gray-600">
                <div className="flex items-center">
                  <User size={20} className="mr-3" />
                  <span>{currentBook.author}</span>
                </div>
                <div className="flex items-center">
                  <Calendar size={20} className="mr-3" />
                  <span>{currentBook.year}</span>
                </div>
                {currentBook.isbn && (
                  <div className="flex items-center">
                    <BookOpen size={20} className="mr-3" />
                    <span className="font-mono">{currentBook.isbn}</span>
                  </div>
                )}
              </div>
            </div>
            <div className="flex space-x-2 ml-6">
              <Button variant="secondary">
                <Edit size={16} className="mr-2" />
                Edit
              </Button>
              <Button variant="danger" onClick={handleDelete}>
                <Trash2 size={16} className="mr-2" />
                Delete
              </Button>
            </div>
          </div>

          {currentBook.description && (
            <div>
              <h2 className="text-xl font-semibold text-gray-900 mb-3">
                Description
              </h2>
              <p className="text-gray-700 leading-relaxed whitespace-pre-wrap">
                {currentBook.description}
              </p>
            </div>
          )}

          <div className="mt-8 pt-6 border-t border-gray-200">
            <div className="grid grid-cols-2 gap-4 text-sm text-gray-600">
              <div>
                <span className="font-medium">Created:</span>{" "}
                {new Date(currentBook.created_at).toLocaleDateString()}
              </div>
              <div>
                <span className="font-medium">Updated:</span>{" "}
                {new Date(currentBook.updated_at).toLocaleDateString()}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
