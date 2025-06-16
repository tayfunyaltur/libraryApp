// src/components/book/BookForm.tsx
import { useEffect } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { type Book, type CreateBookRequest } from "../../services/types";
import { Input } from "../common/Input";
import { Button } from "../common/Button";

const bookSchema = z.object({
  title: z.string().min(1, "Title is required").max(255, "Title is too long"),
  author: z
    .string()
    .min(1, "Author is required")
    .max(255, "Author is too long"),
  year: z
    .number()
    .min(1000, "Year must be at least 1000")
    .max(new Date().getFullYear(), "Year cannot be in the future"),
  isbn: z
    .string()
    .optional()
    .refine((val) => !val || val.length === 13, "ISBN must be 13 characters"),
  description: z.string().optional(),
});

type BookFormData = z.infer<typeof bookSchema>;

interface BookFormProps {
  book?: Book;
  onSubmit: (data: CreateBookRequest) => Promise<boolean>;
  onCancel: () => void;
  loading?: boolean;
}

export function BookForm({
  book,
  onSubmit,
  onCancel,
  loading = false,
}: BookFormProps) {
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
    reset,
  } = useForm<BookFormData>({
    resolver: zodResolver(bookSchema),
    defaultValues: book
      ? {
          title: book.title,
          author: book.author,
          year: book.year,
          isbn: book.isbn || "",
          description: book.description || "",
        }
      : undefined,
  });

  useEffect(() => {
    if (book) {
      reset({
        title: book.title,
        author: book.author,
        year: book.year,
        isbn: book.isbn || "",
        description: book.description || "",
      });
    }
  }, [book, reset]);

  const handleFormSubmit = async (data: BookFormData) => {
    const success = await onSubmit(data);
    if (success) {
      onCancel();
    }
  };

  return (
    <form onSubmit={handleSubmit(handleFormSubmit)} className="space-y-4">
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        <Input
          label="Title"
          {...register("title")}
          error={errors.title?.message}
          required
        />
        <Input
          label="Author"
          {...register("author")}
          error={errors.author?.message}
          required
        />
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        <Input
          label="Year"
          type="number"
          {...register("year", { valueAsNumber: true })}
          error={errors.year?.message}
          required
        />
        <Input
          label="ISBN"
          {...register("isbn")}
          error={errors.isbn?.message}
          helperText="13 characters (optional)"
        />
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Description
        </label>
        <textarea
          {...register("description")}
          rows={4}
          className="input-field resize-none"
          placeholder="Enter book description..."
        />
        {errors.description && (
          <p className="mt-1 text-sm text-red-600">
            {errors.description.message}
          </p>
        )}
      </div>

      <div className="flex space-x-3 pt-4">
        <Button
          type="submit"
          loading={isSubmitting || loading}
          className="flex-1"
        >
          {book ? "Update Book" : "Create Book"}
        </Button>
        <Button
          type="button"
          variant="outline"
          onClick={onCancel}
          disabled={isSubmitting || loading}
        >
          Cancel
        </Button>
      </div>
    </form>
  );
}
