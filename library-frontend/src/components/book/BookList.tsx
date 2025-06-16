// src/components/book/BookList.tsx
import React, { useState } from "react";
import { useBooks } from "../../context/BookContext";
import { BookCard } from "./BookCard";
import { BookModal } from "./BookModal";
import { Loading } from "../common/Loading";
import { Button } from "../common/Button";
import { Plus, Search, Filter, BookOpen } from "lucide-react";
import { Input } from "../common/Input";
import { type Book } from "../../services/types";

export function BookList() {
  const {
    books,
    loading,
    error,
    total,
    page,
    setPage,
    searchBooks,
    setFilters,
    deleteBook,
  } = useBooks();

  const [isModalOpen, setIsModalOpen] = useState(false);
  const [editingBook, setEditingBook] = useState<Book | null>(null);
  const [searchQuery, setSearchQuery] = useState("");
  const [showFilters, setShowFilters] = useState(false);
  const [filterAuthor, setFilterAuthor] = useState("");
  const [filterYear, setFilterYear] = useState("");

  const handleEdit = (book: Book) => {
    setEditingBook(book);
    setIsModalOpen(true);
  };

  const handleDelete = async (id: number) => {
    if (window.confirm("Are you sure you want to delete this book?")) {
      await deleteBook(id);
    }
  };

  const handleView = (id: number) => {
    // Navigate to book detail page
    window.location.href = `/books/${id}`;
  };

  const handleSearch = async (e: React.FormEvent) => {
    e.preventDefault();
    if (searchQuery.trim()) {
      await searchBooks(searchQuery);
    }
  };

  const handleFilter = () => {
    const filters: any = { limit: 10, offset: 0 };
    if (filterAuthor) filters.author = filterAuthor;
    if (filterYear) filters.year = parseInt(filterYear);
    setFilters(filters);
  };

  const clearFilters = () => {
    setFilterAuthor("");
    setFilterYear("");
    setSearchQuery("");
    setFilters({ limit: 10, offset: 0 });
  };

  if (loading && books.length === 0) {
    return <Loading text="Loading books..." />;
  }

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">
            Library Dashboard
          </h1>
          <p className="text-gray-600">Manage your book collection</p>
        </div>
        <Button onClick={() => setIsModalOpen(true)}>
          <Plus size={16} className="mr-2" />
          Add Book
        </Button>
      </div>

      {/* Search and Filters */}
      <div className="card">
        <form onSubmit={handleSearch} className="flex gap-4 mb-4">
          <div className="flex-1">
            <Input
              placeholder="Search books by title, author, or description..."
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
            />
          </div>
          <Button type="submit">
            <Search size={16} className="mr-2" />
            Search
          </Button>
          <Button
            type="button"
            variant="outline"
            onClick={() => setShowFilters(!showFilters)}
          >
            <Filter size={16} className="mr-2" />
            Filters
          </Button>
        </form>

        {showFilters && (
          <div className="border-t pt-4">
            <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
              <Input
                label="Filter by Author"
                placeholder="Author name"
                value={filterAuthor}
                onChange={(e) => setFilterAuthor(e.target.value)}
              />
              <Input
                label="Filter by Year"
                type="number"
                placeholder="Publication year"
                value={filterYear}
                onChange={(e) => setFilterYear(e.target.value)}
              />
              <div className="flex items-end gap-2">
                <Button onClick={handleFilter} className="flex-1">
                  Apply Filters
                </Button>
                <Button variant="outline" onClick={clearFilters}>
                  Clear
                </Button>
              </div>
            </div>
          </div>
        )}
      </div>

      {/* Error Message */}
      {error && (
        <div className="bg-red-50 border border-red-200 rounded-lg p-4">
          <p className="text-red-800">{error}</p>
        </div>
      )}

      {/* Books Grid */}
      {books.length === 0 && !loading ? (
        <div className="text-center py-12">
          <div className="text-gray-400 mb-4">
            <BookOpen size={48} className="mx-auto" />
          </div>
          <h3 className="text-lg font-medium text-gray-900 mb-2">
            No books found
          </h3>
          <p className="text-gray-600 mb-4">
            {searchQuery || filterAuthor || filterYear
              ? "No books match your search criteria."
              : "Get started by adding your first book to the library."}
          </p>
          <Button onClick={() => setIsModalOpen(true)}>
            <Plus size={16} className="mr-2" />
            Add Your First Book
          </Button>
        </div>
      ) : (
        <>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {books.map((book) => (
              <BookCard
                key={book.id}
                book={book}
                onEdit={handleEdit}
                onDelete={handleDelete}
                onView={handleView}
              />
            ))}
          </div>

          {/* Pagination */}
          {total > 10 && (
            <div className="flex justify-center items-center space-x-4">
              <Button
                variant="outline"
                disabled={page === 1}
                onClick={() => setPage(page - 1)}
              >
                Previous
              </Button>
              <span className="text-sm text-gray-600">
                Page {page} of {Math.ceil(total / 10)}
              </span>
              <Button
                variant="outline"
                disabled={page >= Math.ceil(total / 10)}
                onClick={() => setPage(page + 1)}
              >
                Next
              </Button>
            </div>
          )}
        </>
      )}

      {/* Modal */}
      <BookModal
        isOpen={isModalOpen}
        onClose={() => {
          setIsModalOpen(false);
          setEditingBook(null);
        }}
        book={editingBook}
      />
    </div>
  );
}
