import { Routes, Route, Navigate } from "react-router-dom";
import { Dashboard, Auth } from "@/layouts";
import ProtectedRoute from "./pages/auth/ProtectedRoute";
import Lectures from "./pages/dashboard/lectures";

function App() {
  return (
    <Routes>
      {/* <Route element={<ProtectedRoute allowedRoles={['curator']} />} >
        <Route path="/dashboard/home" element={<Dashboard />} />
      </Route>
      <Route element={<ProtectedRoute allowedRoles={['candidat']} />} >
        <Route path="/dashboard/form" element={<Dashboard />} />
      </Route> */}
      <Route element={<ProtectedRoute allowedRoles={['intern', 'candidat', 'curator']} />} >
        <Route path="/dashboard/*" element={<Dashboard />} />
      </Route>
      
      <Route path="/auth/*" element={<Auth />} />      
      <Route path="*" element={<Navigate to="/auth/login" replace />} />
    </Routes>
  );
}

export default App;
