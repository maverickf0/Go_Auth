import { Routes, Route, Navigate } from 'react-router-dom';
import Login from './components/Login';
import Register from './components/Register';
import { useUser } from './context/UserContext';
import Dashboard from './components/Dashboard';

function App() {

  const { user } = useUser();

  return (
      <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route path="/" element={<Navigate to="/login" replace />} />
          {
            !!user && <Route path="/dashboard" element={<Dashboard />} />
          }
        </Routes>
      </div>
  );
}

export default App;