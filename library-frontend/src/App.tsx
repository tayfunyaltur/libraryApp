// src/App.tsx
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import "./index.css";
import { Toaster } from "react-hot-toast";
import { BookProvider } from "./context/BookContext";
import { Dashboard } from "./pages/Dashboard";
import { BookDetailPage } from "./pages/BookDetailPage";

function App() {
  return (
    <BookProvider>
      <Router>
        <div className="App">
          <Toaster
            position="top-right"
            toastOptions={{
              duration: 4000,
              style: {
                background: "#363636",
                color: "#fff",
              },
              success: {
                duration: 3000,
                iconTheme: {
                  primary: "#4ade80",
                  secondary: "#fff",
                },
              },
              error: {
                duration: 5000,
                iconTheme: {
                  primary: "#ef4444",
                  secondary: "#fff",
                },
              },
            }}
          />

          <Routes>
            <Route path="/" element={<Dashboard />} />
            <Route path="/books/:id" element={<BookDetailPage />} />
            <Route path="*" element={<div>Page not found</div>} />
          </Routes>
        </div>
      </Router>
    </BookProvider>
  );
}

export default App;
